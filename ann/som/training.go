package som

import (
	"context"
	"github.com/publiczny81/ml/ann/initializers"
	"github.com/publiczny81/ml/calculus/vector"
	"github.com/publiczny81/ml/calculus/vector/operations"
	"github.com/publiczny81/ml/sampling"
	"github.com/publiczny81/ml/utils"
	"runtime"
	"sync"
)

var (
	defaultInitializer = initializers.NewNormal(utils.Rand)
)

type sampler interface {
	Samples(ctx context.Context) <-chan sampling.Sample[[]float64]
}

type learningRateSchedule interface {
	LearningRate(epoch int) float64
}

type neighborhood interface {
	NeighborRate([]float64, []float64, int) float64
}

type initializer interface {
	Initialize(s []float64)
}

type Trainer struct {
	initializer
	sampler
	learningRateSchedule
	neighborhood
}

type TrainerOption func(*Trainer)

func WithInitializer(i initializer) TrainerOption {
	return func(t *Trainer) {
		t.initializer = i
	}
}

func NewTrainer(sampler sampler, schedule learningRateSchedule, neighborhood neighborhood, opts ...TrainerOption) (t *Trainer) {
	t = &Trainer{
		initializer:          defaultInitializer,
		sampler:              sampler,
		learningRateSchedule: schedule,
		neighborhood:         neighborhood,
	}
	for _, opt := range opts {
		opt(t)
	}
	return
}

func (t *Trainer) Train(ctx context.Context, network *Network, epochs int) (err error) {
	t.Initialize(network.Weights)

	return t.train(ctx, network, epochs, 1)
}

func (t *Trainer) train(ctx context.Context, network *Network, epochs, epoch int) (err error) {
	if epochs < epoch {
		return
	}
	if err = t.trainSample(ctx, network, epochs, epoch, t.sampler.Samples(ctx)); err != nil {
		return
	}
	return t.train(ctx, network, epochs, epoch+1)
}

func (t *Trainer) trainSample(ctx context.Context, network *Network, epochs, epoch int, samples <-chan sampling.Sample[[]float64]) (err error) {
	select {
	case <-ctx.Done():
		err = ctx.Err()
		return
	case sample, ok := <-samples:
		if !ok {
			return
		}
		if sample.Error != nil {
			err = sample.Error
			return
		}

		bmu := network.BestMatchingUnit(sample.Value)

		if err = t.update(ctx, network, epoch, sample.Value, bmu); err != nil {
			return
		}
		return t.trainSample(ctx, network, epochs, epoch, samples)
	}
}

func (t *Trainer) update(ctx context.Context, network *Network, epoch int, features []float64, bmu Point) (err error) {
	var (
		wg      sync.WaitGroup
		threads = min(runtime.NumCPU()*2-1, len(network.Neurons))
		ch      = make(chan *Neuron, min(threads, len(network.Neurons)))
	)

	for range threads {
		wg.Add(1)
		go func() {
			var count int
			defer func() {
				wg.Done()
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case n, ok := <-ch:
					if !ok {
						return
					}
					count++
					var factor = t.learningRateSchedule.LearningRate(epoch) * t.neighborhood.NeighborRate(bmu, n.Point, epoch)
					if factor <= 0 {
						continue
					}
					var weights = vector.Subtract(features, n.Weights)
					vector.Wrap(weights).Apply(operations.Multiply(factor))
					vector.Wrap(n.Weights).Apply(operations.Add(weights))
				}
			}
		}()
	}

	go func() {
		defer close(ch)
		for _, n := range network.Neurons {
			select {
			case <-ctx.Done():
				err = ctx.Err()
				return
			case ch <- n:

			}
		}
	}()
	wg.Wait()
	return
}

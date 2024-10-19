package som

import (
	"context"
	"github.com/publiczny81/ml/calculus/vector"
	"github.com/publiczny81/ml/calculus/vector/operations"
	"github.com/publiczny81/ml/sampling"
	"runtime"
	"sync"
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

type Trainer struct {
	sampler
	learningRateSchedule
	neighborhood
}

func NewTrainer(sampler sampler, schedule learningRateSchedule, neighborhood neighborhood) *Trainer {
	return &Trainer{
		sampler:              sampler,
		learningRateSchedule: schedule,
		neighborhood:         neighborhood,
	}
}

func (t *Trainer) Train(ctx context.Context, network *Network, epochs int) (err error) {
	for epoch := range epochs {
		if err = t.train(ctx, network, epoch); err != nil {
			return
		}
	}
	return
}

func (t *Trainer) train(ctx context.Context, network *Network, epoch int) (err error) {
	for sample := range t.sampler.Samples(ctx) {
		if sample.Error != nil {
			err = sample.Error
			return
		}

		bmu := network.BestMatchingUnit(sample.Value)

		if err = t.update(ctx, network, epoch, sample.Value, bmu); err != nil {
			return
		}
	}
	return
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

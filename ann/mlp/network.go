package mlp

import (
	"context"
	"github.com/publiczny81/ml/activate"
	"github.com/publiczny81/ml/calculus/vector"
	"github.com/publiczny81/ml/errors"
	"runtime"
	"sync"
)

type Rand interface {
	Float64() float64
}

type LayerSpec struct {
	Activation string
	Neurons    int
}

type Options struct {
	Input   int
	Layers  []LayerSpec
	Weights []float64
}

func (o *Options) CountWeights() (count int) {
	var previous = o.Input + 1
	for i := 0; i < len(o.Layers); i++ {
		count += previous * o.Layers[i].Neurons
		previous = o.Layers[i].Neurons + 1
	}
	return
}

type Option func(options *Options) error

func AddLayer(neurons int, activation string) Option {
	return func(options *Options) error {
		options.Layers = append(options.Layers, LayerSpec{
			Activation: activation,
			Neurons:    neurons,
		})
		return nil
	}
}

func WithWeights(weights []float64) Option {
	return func(options *Options) error {
		options.Weights = weights
		return nil
	}
}

type layer struct {
	Activation activate.Activate
	Input      []float64
	Output     []float64
	Weights    []float64
}

func (l *layer) getThreads() int {
	return min(runtime.NumCPU()*2-1, len(l.Weights)/len(l.Input))
}

func (l *layer) Activate(ctx context.Context) (err error) {
	type Neuron struct {
		Start int
		End   int
		Index int
	}
	var (
		threads = l.getThreads()
		wg      sync.WaitGroup
		ch      = make(chan *Neuron, threads)
	)
	for range threads {
		wg.Add(1)
		go func() {
			for n := range ch {
				l.Output[n.Index] = l.Activation.Function(vector.DotProduct(l.Input, l.Weights[n.Start:n.End]))
			}
			wg.Done()
		}()
	}

	go func() {
		defer close(ch)
		var (
			start = 0
			end   = len(l.Input)
		)
		for idx := range len(l.Output) {
			select {
			case <-ctx.Done():
				err = ctx.Err()
				return
			case ch <- &Neuron{
				Start: start,
				End:   end,
				Index: idx,
			}:
			}
		}
	}()
	wg.Wait()
	return
}

type Network struct {
	Options
	Layers []layer
}

func New(input int, opts ...Option) (net *Network, err error) {
	var options Options

	options.Input = input

	for _, o := range opts {
		if err = o(&options); err != nil {
			return
		}
	}

	net = &Network{
		Options: options,
	}
	return
}

func (net *Network) Init(opts ...Option) (err error) {
	for _, o := range opts {
		if err = o(&net.Options); err != nil {
			return
		}
	}
	if err = net.validate(); err != nil {
		return
	}

	return net.init()

}

func (net *Network) Activate(ctx context.Context, input []float64) (output []float64, err error) {
	if len(input) != net.Options.Input {
		err = errors.WithMessagef(errors.InvalidParameterValueError, "len(input)=%d)", len(input))
		return
	}
	copy(net.Layers[0].Input, input)
	for _, l := range net.Layers {
		if err = l.Activate(ctx); err != nil {
			return
		}
	}
	copy(output, net.Layers[len(net.Layers)-1].Output)
	return
}

func (net *Network) validate() (err error) {
	var options = &net.Options
	if options.Input < 1 {
		return errors.WithMessagef(errors.InvalidParameterValueError, "network.Input=%d", options.Input)
	}

	if len(options.Layers) < 1 {
		return errors.WithMessage(errors.InvalidParameterValueError, "network.Layers={}")
	}

	if len(options.Weights) == 0 {
		net.Weights = make([]float64, net.CountWeights())
		return
	}

	if len(options.Weights) != options.CountWeights() {
		return errors.WithMessagef(errors.InvalidParameterValueError, "incompatible parameters values len(network.Weights)=%d and network.Layers=%v", len(options.Weights), options.Layers)
	}

	return
}

func (net *Network) init() (err error) {
	var options = &net.Options
	var (
		previous   = options.Input + 1
		start      = 0
		end        = 0
		input      = make([]float64, options.Input, options.Input+1)
		output     []float64
		activation activate.Activate
		found      bool
	)

	for i := 0; i < len(options.Layers); i++ {
		input = append(input, 1.0)
		output = make([]float64, options.Layers[i].Neurons, options.Layers[i].Neurons+1)
		end += previous * options.Layers[i].Neurons
		previous = options.Layers[i].Neurons + 1
		if activation, found = activate.Get(options.Layers[i].Activation); !found {
			return errors.WithMessagef(errors.InvalidParameterValueError, "network.Layer[%d].Activation=")
		}

		net.Layers = append(net.Layers, layer{
			Activation: activation,
			Input:      input,
			Output:     output,
			Weights:    options.Weights[start:end],
		})
		start = end
		input = output
	}
	return
}

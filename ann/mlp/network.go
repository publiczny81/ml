package mlp

//
//import (
//	"github.com/pkg/errors"
//	"github.com/publiczny81/ml/ann/neuron"
//	"github.com/publiczny81/ml/calculus/types"
//	errors2 "github.com/publiczny81/ml/errors"
//	"github.com/publiczny81/ml/utils/slices"
//)
//
//type NetworkOptions[T types.Float] struct {
//	Layers     []int
//	Weights    []T
//	Activation neuron.ActivateFunc[T]
//	Rand       neuron.Rand[T]
//}
//
//type NetworkOption[T types.Float] func(options *NetworkOptions[T]) error
//
//func WithRand[T types.Float](rand neuron.Rand[T]) NetworkOption[T] {
//	return func(options *NetworkOptions[T]) error {
//		options.Rand = rand
//		return nil
//	}
//}
//
//func WithWeights[T types.Float](weights []T) NetworkOption[T] {
//	return func(options *NetworkOptions[T]) error {
//		options.Weights = weights
//		return nil
//	}
//}
//
//func WithActivateFunc[T types.Float](activation neuron.ActivateFunc[T]) NetworkOption[T] {
//	return func(options *NetworkOptions[T]) error {
//		options.Activation = activation
//		return nil
//	}
//}
//
//type Layer[T types.Float] []*neuron.Neuron[T]
//
//type Network[T types.Float] struct {
//	Weights []T
//	Neurons []Layer[T]
//	//Result  [][]T
//}
//
//func New[T types.Float](layers []int, opts ...NetworkOption[T]) (net *Network[T], err error) {
//	var options = NetworkOptions[T]{
//		Layers: layers,
//	}
//
//	for _, o := range opts {
//		if err = o(&options); err != nil {
//			return
//		}
//	}
//
//	if err = options.ValidateAndInitNetworkOptions(); err != nil {
//		return
//	}
//
//	net = &Network[T]{
//		Weights: options.Weights,
//		Neurons: make([]Layer[T], len(options.Layers)-1),
//		//Result:  make([][]T, len(options.Layers)),
//	}
//	var (
//		start = 0
//	)
//	//slice.ApplyWithIndex(net.Result, func(i int, t []T) []T {
//	//	return make([]T, options.Layers[i])
//	//})
//
//	slices.ApplyWithIndex(net.Neurons, func(i int, layer Layer[T]) Layer[T] {
//		var delta = options.Layers[i] + 1
//		layer = make(Layer[T], options.Layers[i+1])
//		slices.Apply(layer, func(n *neuron.Neuron[T]) *neuron.Neuron[T] {
//			var end = start + delta
//			n = neuron.New(options.Activation, options.Weights[start:end])
//			start = end
//			return n
//		})
//		return layer
//	})
//	return
//}
//
//func (net *Network[T]) MakeResult(features int) (result [][]T) {
//	result = make([][]T, len(net.Neurons)+1)
//	result[0] = make([]T, features)
//	slices.ApplyWithIndex(result[1:], func(i int, ts []T) []T {
//		return make([]T, len(net.Neurons[i]))
//	})
//	return
//}
//
//func (net *Network[T]) Process(input []T) []T {
//	var result = net.MakeResult(len(input))
//	net.processWithResult(input, result)
//
//	return result[len(net.Neurons)]
//}
//
//type Result[T types.Float] [][]T
//
//func (r Result[T]) Adjust(layer, required int) Result[T] {
//	if layer > len(r) {
//
//	}
//	if len(r[layer]) == required {
//		return r
//	}
//	if len(r[layer]) > required {
//		r[layer] = r[layer][:required]
//		return r
//	}
//	if cap(r[layer]) >= required {
//		r[layer] = r[layer][:required]
//		return r
//	}
//	r[layer] = make([]T, required)
//	return r
//}
//
//func (net *Network[T]) validate(input []T, result [][]T) (err error) {
//	if len(result) != len(net.Neurons)+1 {
//		err = errors.Wrap(errors2.InvalidParameterError, "size of result")
//		return
//	}
//	if len(input) != len(result[0]) {
//		err = errors.Wrapf(errors2.InvalidParameterError, "incompatible input and result")
//		return
//	}
//	slices.IterateWithIndex(net.Neurons, func(i int, l Layer[T]) bool {
//		if len(l) != len(result[i+1]) {
//			err = errors.Wrap(errors2.InvalidParameterError, "size of result")
//			return false
//		}
//		return true
//	})
//	return
//}
//
//func (net *Network[T]) ProcessWithResult(input []T, result [][]T) {
//	if err := net.validate(input, result); err != nil {
//		panic(err)
//	}
//	net.processWithResult(input, result)
//}
//
//func (net *Network[T]) processWithResult(input []T, result [][]T) {
//	copy(result[0], input)
//	slices.IterateWithIndex(net.Neurons, func(i int, layer Layer[T]) bool {
//		slices.IterateWithIndex(layer, func(j int, n *neuron.Neuron[T]) bool {
//			result[i+1][j] = n.Activate(result[i])
//			return true
//		})
//		return true
//	})
//}
//
//func (options *NetworkOptions[T]) ValidateAndInitNetworkOptions() (err error) {
//	if len(options.Layers) < 1 {
//		err = errors.WithMessage(errors2.InvalidParameterError, "layers")
//		return
//	}
//	var (
//		countWeights int
//	)
//
//	slices.IterateWithIndex(options.Layers[1:], func(i int, i2 int) bool {
//		countWeights += (options.Layers[i] + 1) * i2
//		return true
//	})
//
//	if len(options.Weights) > 0 && len(options.Weights) != countWeights {
//		err = errors.WithMessage(errors2.InvalidParameterError, "incompatible parameters layers and weights")
//		return
//	}
//	if len(options.Weights) == 0 && options.Rand == nil {
//		err = errors.WithMessage(errors2.InvalidParameterError, "missing parameter rand or weights")
//		return
//	}
//	if len(options.Weights) == 0 {
//		options.Weights = make([]T, countWeights)
//		slices.Apply(options.Weights, func(_ T) T {
//			return options.Rand.Float()
//		})
//	}
//	if options.Activation == nil {
//		err = errors.WithMessage(errors2.InvalidParameterError, "missing parameter activation")
//		return
//	}
//	return
//}

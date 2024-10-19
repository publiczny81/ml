package neuron

//
//import (
//	"github.com/publiczny81/ml/functions"
//)
//
//var DefaultActivationFunc = functions.Sigmoid
//
//type ActivationFunc func(float64) float64
//
//type Rand interface {
//	Float64() float64
//}
//
//type Option func(*Neuron)
//
//func WithWeights(weights []float64) Option {
//	return func(neuron *Neuron) {
//		neuron.Weights = weights
//	}
//}
//
//func WithActivationFunc(activation ActivationFunc) Option {
//	return func(neuron *Neuron) {
//		neuron.Activation = activation
//	}
//}
//
//func WithInput(input int, rand Rand) Option {
//	return func(neuron *Neuron) {
//		neuron.Weights = make([]float64, input+1)
//		slices.Apply(neuron.Weights, func(f float64) float64 {
//			return rand.Float64()
//		})
//	}
//}
//
//type Neuron struct {
//	Activation ActivationFunc
//	Weights    []float64
//}
//
//func New(options ...Option) (n *Neuron) {
//	n = &Neuron{
//		Activation: DefaultActivationFunc,
//	}
//	for _, o := range options {
//		o(n)
//	}
//	return
//}
//
//func (n *Neuron) Randomize(rand Rand) {
//	for i := range n.Weights {
//		n.Weights[i] = rand.Float64()
//	}
//}
//
//func (n *Neuron) Activate(input []float64) (value float64) {
//	for i, v := range input {
//		value += v * n.Weights[i]
//	}
//	value = n.Activation(value + n.Weights[len(n.Weights)-1])
//	return
//}

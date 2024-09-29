package neuron

import (
	"math"
)

func Sigmoid(value float64) float64 {
	return 1 / (1 + math.Exp(value))
}

var DefaultActivationFunc = Sigmoid

type ActivationFunc func(float64) float64

type Rand interface {
	Float64() float64
}

type Neuron struct {
	Activation ActivationFunc
	Weights    []float64
	Bias       float64
}

func New(input int) *Neuron {
	return &Neuron{
		Activation: DefaultActivationFunc,
		Weights:    make([]float64, input),
	}
}

func (n *Neuron) Randomize(rand Rand) {
	for i := range n.Weights {
		n.Weights[i] = rand.Float64()
	}
	n.Bias = rand.Float64()
}

func (n *Neuron) Process(input []float64) (value float64) {
	for i, v := range input {
		value += v * n.Weights[i]
	}
	value = n.Activation(value + n.Bias)
	return
}

type Builder func() *Neuron

func NewBuilder(input int) Builder {
	return func() *Neuron {
		return New(input)
	}
}

func (b Builder) WithRandomizedWeights(rand Rand) Builder {
	var n = b()
	return func() *Neuron {
		n.Randomize(rand)
		return n
	}
}

func (b Builder) WithWeights(weights []float64) Builder {
	var neuron = b()
	return func() *Neuron {
		var limit = min(len(neuron.Weights), len(weights))
		copy(neuron.Weights, weights[:limit])
		return neuron
	}
}

func (b Builder) WithBias(bias float64) Builder {
	var n = b()
	return func() *Neuron {
		n.Bias = bias
		return n
	}
}

func (b Builder) WithActivationFunc(activationFunc ActivationFunc) Builder {
	var neuron = b()
	return func() *Neuron {
		neuron.Activation = activationFunc
		return neuron
	}
}

func (b Builder) Build() *Neuron {
	return b()
}

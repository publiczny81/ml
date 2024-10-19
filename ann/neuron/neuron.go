package neuron

import (
	"github.com/pkg/errors"
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/vector"
	errors2 "github.com/publiczny81/ml/errors"
	"github.com/publiczny81/ml/utils/slices"
)

type ActivateFunc[S ~[]T, T types.Float] func(S, S) T

type Neuron[T types.Float] struct {
	Weights []T
	ActivateFunc[[]T, T]
}

func New[T types.Float](activateFunc ActivateFunc[[]T, T], weights []T) *Neuron[T] {
	if len(weights) == 0 {
		panic(errors.WithMessage(errors2.InvalidParameterError, "neuron.New: weights is nil"))
	}
	if activateFunc == nil {
		panic(errors.WithMessage(errors2.InvalidParameterError, "neuron.New: activationFunc is nil"))
	}
	return &Neuron[T]{
		Weights:      weights,
		ActivateFunc: activateFunc,
	}
}

type Builder[T types.Float] func() *Neuron[T]

func NewBuilder[T types.Float]() Builder[T] {
	return func() *Neuron[T] {
		return &Neuron[T]{}
	}
}

func (b Builder[T]) WithActivateFunc(activateFunc ActivateFunc[[]T, T]) Builder[T] {
	return func() (n *Neuron[T]) {
		n = b()
		n.ActivateFunc = activateFunc
		return
	}
}

func (b Builder[T]) WithWeights(weights []T) Builder[T] {
	return func() (n *Neuron[T]) {
		n = b()
		n.Weights = weights
		return
	}
}

func (b Builder[T]) WithFeatures(features int, bias bool, rand Rand[T]) Builder[T] {
	if rand == nil {
		panic(errors.WithMessage(errors2.InvalidParameterError, "neuron.WithFeatures: missing rand"))
	}
	return func() (n *Neuron[T]) {
		n = b()
		n.Weights = make([]T,
			func() int {
				if bias {
					return features + 1
				}
				return features
			}(),
		)
		slices.Apply(n.Weights, func(v T) T {
			return rand.Float()
		})
		return n
	}
}

func (b Builder[T]) Build() (n *Neuron[T]) {
	n = b()
	if len(n.Weights) == 0 {
		panic(errors.WithMessage(errors2.InvalidParameterError, "neuron.Build: missing weights"))
	}
	if n.ActivateFunc == nil {
		panic(errors.WithMessage(errors2.InvalidParameterError, "neuron.Build: missing activation function"))
	}
	return
}

func (n *Neuron[T]) Activate(features []T) T {
	return n.ActivateFunc(features, n.Weights)
}

type Rand[T types.Float] interface {
	Float() T
}

type activationFunc[T types.Float] func(T) T

func NewActivateFunc[T types.Float](activation activationFunc[T]) ActivateFunc[[]T, T] {
	if activation == nil {
		panic(errors.WithMessage(errors2.InvalidParameterError, "neuron.NewActivateFunc: activation is nil"))
	}
	return func(features []T, weights []T) T {
		return activation(vector.DotProduct(features, weights))
	}
}

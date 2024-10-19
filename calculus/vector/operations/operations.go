package operations

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/vector/pool"
	"github.com/publiczny81/ml/errors"
)

type MetricsFunc[S ~[]T, T types.Float] func(S, S) T

func Add[V ~[]T, T types.Real](v1 V) (res types.VOperation[T]) {
	return func(v types.V[T]) {
		if len(v) != len(v1) {
			panic(errors.UnmatchedSizeOfVectorsError)
		}
		for i := range v {
			v[i] = v[i] + v1[i]
		}
	}
}

func Subtract[S ~[]T, T types.Real](v1, v2 S) (res types.VOperation[T]) {
	return func(v types.V[T]) {
		if len(v1) != len(v2) {
			panic(errors.UnmatchedSizeOfVectorsError)
		}
		for i := range v {
			v[i] = v1[i] - v2[i]
		}
	}
}

func Multiply[T types.Real](c T) (res types.VOperation[T]) {
	return func(v types.V[T]) {
		for i := range v {
			v[i] *= c
		}
	}
}

func Normalize[S ~[]T, T types.Float](metricsFunc MetricsFunc[S, T]) (res types.VOperation[T]) {
	return func(v types.V[T]) {
		var (
			zeros  = pool.Get[S](len(v))
			length = metricsFunc(zeros, S(v))
		)
		if length == 0 {
			return
		}
		v.Apply(Multiply[T](1 / length))
	}
}

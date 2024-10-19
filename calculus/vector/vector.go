package vector

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/vector/constructors"
	"github.com/publiczny81/ml/calculus/vector/operations"
	"github.com/publiczny81/ml/calculus/vector/pool"
	"github.com/publiczny81/ml/errors"
	"github.com/publiczny81/ml/utils/slices"
)

func Wrap[T types.Real](d []T) types.V[T] {
	return d
}

func Zeros[T types.Real](size int) types.V[T] {
	return pool.Get[types.V[T]](size)
}

func CopyOf[T types.Real](s []T) (r types.V[T]) {
	r = pool.Get[[]T](len(s))
	copy(r, s)
	return
}

func Exclude[S ~[]T, T types.Real](v S, idx int) (r S) {
	if idx < 0 {
		return v
	}
	if idx >= len(v) {
		return v
	}
	r = pool.Get[S](len(v) - 1)
	copy(r[:idx], v[:idx])
	copy(r[idx:], v[idx+1:])
	return
}

func Add[V ~[]T, T types.Real](v1, v2 V) (res V) {
	res = constructors.CopyOf(v1)
	Wrap(res).Apply(operations.Add(v2))
	return
}

func Apply[V ~[]T, T types.Real](v V, f func(int, T) T) {
	if f == nil {
		return
	}
	for i, e := range v {
		v[i] = f(i, e)
	}
}

func DotProduct[V ~[]T, T types.Real](v1, v2 V) (result T) {
	if len(v1) != len(v2) {
		panic(errors.UnmatchedSizeOfVectorsError)
	}
	for i, e := range v1 {
		result += e * v2[i]
	}
	return
}

func Product[M ~[][]T, V ~[]T, T types.Real](m M, v V) (result V) {
	result = pool.Get[V](len(m))
	for i, row := range m {
		result[i] = DotProduct(V(row), v)
	}
	return
}

func Length[V ~[]T, T types.Float](v V, metrics types.MetricsFunc[V, T]) (result T) {
	var zeros = pool.Get[V](len(v))
	return metrics(zeros, v)
}

func Subtract[V ~[]T, T types.Real](v1, v2 V) (res V) {
	res = pool.Get[V](len(v1))
	Wrap(res).Apply(operations.Subtract(v1, v2))
	return
}

func Multiply[V ~[]T, T types.Real](v1 V, c T) (res V) {
	var v = CopyOf(v1)
	v.Apply(operations.Multiply(c))
	res = V(v)
	return
}

func Normalize[V ~[]T, T types.Float](values V, metrics operations.MetricsFunc[V, T]) (result V) {
	var v = CopyOf(values)
	v.Apply(operations.Normalize(metrics))
	result = V(v)
	return
}

func Resize[V ~[]T, T types.Real](v V, size int) (r V) {
	return slices.Resize(v, size)
}

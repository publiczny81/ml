package constructors

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/vector/pool"
	"github.com/publiczny81/ml/errors"
)

func CopyOf[R ~[][]T, T types.Real](r R) (m R) {
	m = make(R, len(r))
	for i, v := range r {
		m[i] = pool.Get[[]T](len(v))
		copy(m[i], v)
	}
	return
}

func Zeros[T types.Real](shape ...int) (r [][]T) {
	switch len(shape) {
	case 0:
		panic(errors.InvalidSizeOfMatrixError)
	case 1:
		shape = append(shape, shape[0])
	}
	var (
		data       = pool.Get[[]T](shape[0] * shape[1])
		start, end = 0, shape[1]
	)
	r = make([][]T, shape[0])
	for i := range r {
		r[i] = data[start:end]
		start = end
		end += shape[1]
	}
	return
}

func Identity[T types.Real](size int) (r [][]T) {
	r = Zeros[T](size)
	for i := range r {
		r[i][i] = 1
	}
	return
}

func Wrap[R ~[][]T, T types.Real](m R) types.M[T] {
	return types.M[T](m)
}

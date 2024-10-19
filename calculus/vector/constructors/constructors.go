package constructors

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/vector/pool"
)

func Wrap[S ~[]T, T types.Real](d S) types.V[T] {
	if d == nil {
		return types.V[T]{}
	}
	return types.V[T](d)
}

func CopyOf[S ~[]T, T types.Real](s S) (r S) {
	r = pool.Get[[]T](len(s))
	copy(r, s)
	return
}

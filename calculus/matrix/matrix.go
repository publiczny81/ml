package matrix

import (
	"github.com/publiczny81/ml/calculus/matrix/concurrent"
	"github.com/publiczny81/ml/calculus/matrix/constructors"
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/utils"
	"github.com/publiczny81/ml/calculus/vector"
	"github.com/publiczny81/ml/errors"
)

func CopyOf[R ~[][]T, T types.Real](r R) (m types.M[T]) {
	return Wrap(constructors.CopyOf(r))
}

func Zeros[T types.Real](shape ...int) (m types.M[T]) {
	return constructors.Zeros[T](shape...)
}

func Identity[T types.Real](size int) (m types.M[T]) {
	return constructors.Identity[T](size)
}

func Wrap[R ~[][]T, T types.Real](m R) types.M[T] {
	return types.M[T](m)
}

func SetValues[R ~[][]T, T types.Real](values R) types.MOperation[T] {
	return func(m types.M[T]) {
		for i := range m {
			copy(m[i], values[i])
		}
	}
}

func ForEachColumn[T types.Real](f func(j int, c []T) []T) types.MOperation[T] {
	return func(m types.M[T]) {
		for j := 0; j < len(m[0]); j++ {
			var col = f(j, Column(m, j))
			for r := range col {
				(m)[r][j] = col[r]
			}
		}
	}
}

func SwapRows[T types.Real](i, j int) types.MOperation[T] {
	return func(m types.M[T]) {
		if i == j {
			return
		}
		m[i], m[j] = m[j], m[i]
	}
}

func SwapColumns[T types.Real](i, j int) types.MOperation[T] {
	return func(m types.M[T]) {
		if i == j {
			return
		}
		for k := range m {
			m[k][i], m[k][j] = m[k][j], m[k][i]
		}
		return
	}
}

func ForEach[T types.Real](f func(i, j int, t T) T) types.MOperation[T] {
	return func(m types.M[T]) {
		for i := range m {
			for j := range m[i] {
				m[i][j] = f(i, j, m[i][j])
			}
		}
		return
	}
}

func Add[R ~[][]T, T types.Real](m1, m2 R) R {
	if len(m1) != len(m2) {
		panic(errors.UnmatchedSizeOfMatricesError)
	}
	if len(m1[0]) != len(m2[0]) {
		panic(errors.UnmatchedSizeOfMatricesError)
	}

	var result = Zeros[T](len(m1), len(m1[0]))
	result.Apply(concurrent.ForEachRow[T](func(i int, r []T) []T {
		for j := range r {
			r[j] = m1[i][j] + m2[i][j]
		}
		return r
	}))

	return R(result)
}

func Apply[M ~[][]T, T types.Real](m M, operations ...types.MOperation[T]) {
	Wrap(m).Apply(operations...)
	return
}

func Column[M ~[][]T, T types.Real](m M, j int) (result []T) {
	result = make([]T, len(m))
	for i := range result {
		result[i] = m[i][j]
	}
	return
}

func Row[M ~[][]T, T types.Real](m M, i int) []T {
	return m[i]
}

func Subtract[S ~[][]T, T types.Real](m1, m2 S) (r S) {
	r = constructors.CopyOf(m1)
	Wrap(r).Apply(concurrent.Subtract(m2))
	return r
}

func Multiply[R ~[][]T, T types.Real](m R, c T) (result R) {
	result = constructors.CopyOf(m)
	Wrap(result).Apply(concurrent.Multiply(c))
	return result
}

func Transpose[R ~[][]T, T types.Real](m R) R {
	var result = make(R, len(m[0]))
	Wrap(result).Apply(concurrent.Transpose(m))
	return result
}

func Product[R ~[][]T, T types.Real](m1, m2 R) R {
	var result = Zeros[T](len(m1), len(m2[0]))
	var transposed = Transpose(m2)
	result.Apply(concurrent.ForEachRow(func(i int, r []T) []T {
		for j := range r {
			r[j] = vector.DotProduct(m1[i], transposed[j])
		}
		return r
	}))
	return R(result)
}

func ProductV[R ~[][]T, V ~[]T, T types.Real](m R, v []T) V {
	var result = vector.Zeros[T](len(v))
	result.Apply()
	return V(result)
}

func Det[R ~[][]T, T types.Float](m R) (result T) {
	if len(m) != len(m[0]) {
		panic(errors.InvalidSizeOfMatrixError)
	}
	switch len(m) {
	case 0:
	case 1:
		result = m[0][0]
	case 2:
		result = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	case 3:
		result = m[0][0]*m[1][1]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1] - m[0][2]*m[1][1]*m[2][0] - m[0][1]*m[1][0]*m[2][2] - m[0][0]*m[1][2]*m[2][1]
	default:
		var t, odd = UpperTriangular(m)
		result = 1
		for i := range t {
			result *= t[i][i]
		}
		if odd {
			result = -result
		}
	}
	return
}

func Minor[R ~[][]T, T types.Float](m R, i, j int) (result R) {
	var v = Zeros[T](len(m)-1, len(m[0])-1)
	v.Apply(concurrent.Minor(m, i, j))
	return R(v)
}

func Cofactor[R ~[][]T, T types.Float](m R) R {
	if len(m) != len(m[0]) {
		panic(errors.InvalidSizeOfMatrixError)
	}
	var result = Zeros[T](len(m), len(m[0]))
	result.Apply(concurrent.ForEachRow(func(i int, r []T) []T {
		vector.Apply(r, func(j int, t T) (value T) {
			value = Det(Minor(m, i, j))
			if (i+j)%2 != 0 {
				value = -value
			}
			return
		})
		return r
	}))

	return R(result)
}

func Adj[M ~[][]T, T types.Float](m M) (result M) {
	return Transpose(Cofactor(m))
}

func Inverse[M ~[][]T, T types.Float](m M) (result M, exists bool) {
	if len(m) != len(m[0]) {
		panic(errors.InvalidSizeOfMatrixError)
	}
	var det = Det(m)
	if utils.IsZero(det) {
		return
	}

	result = Multiply(Adj(m), 1.0/det)
	exists = true

	return
}

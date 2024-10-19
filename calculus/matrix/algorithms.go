package matrix

import (
	"github.com/publiczny81/ml/calculus/matrix/constructors"
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/utils"
	"github.com/publiczny81/ml/calculus/vector"
)

func GaussElimination[R ~[][]T, T types.Float](m R) (r R, p R) {
	r = constructors.CopyOf(m)
	p = constructors.Identity[T](len(r))
	var (
		result       = Wrap(r)
		permutations = Wrap(p)
	)

	rows := len(result)
	for row := range len(result) {
		var (
			maxEl  = result[row][row]
			maxRow = row
			swap   = false
		)

		for k := row + 1; k < len(result); k++ {
			if utils.Abs(result[k][row]) > utils.Abs(maxEl) {
				maxEl = result[k][row]
				maxRow = k
				swap = true
			}
		}

		if swap {
			result.Apply(SwapRows[T](maxRow, row))
			permutations.Apply(SwapRows[T](maxRow, row))
		}

		for k := row + 1; k < rows; k++ {
			c := -result[k][row] / result[row][row]
			vector.Apply(result[k][row:], func(i int, t T) T {
				if i == 0 {
					return 0
				}
				return t + c*result[row][i+row]
			})
		}
	}
	r = R(result)
	return
}

// UpperTriangular transforms the given matrix into an upper triangular matrix
// and returns the result and a boolean value indicating if the number of row swaps was odd.
func UpperTriangular[R ~[][]T, T types.Float](m R) (r R, odd bool) {
	var result = Wrap(CopyOf(m))
	rows := len(result)
	for row := range len(result) {

		maxEl := result[row][row]
		maxRow := row

		for k := row + 1; k < len(result); k++ {
			if utils.Abs(result[k][row]) > utils.Abs(maxEl) {
				maxEl = result[k][row]
				maxRow = k
			}
		}

		if maxRow != row {
			result.Apply(SwapRows[T](maxRow, row))
			odd = !odd
		}

		for k := row + 1; k < rows; k++ {
			c := -result[k][row] / result[row][row]
			vector.Apply(result[k][row:], func(i int, t T) T {
				if i == 0 {
					return 0
				}
				return t + c*result[row][i+row]
			})
		}
	}
	r = R(result)
	return
}

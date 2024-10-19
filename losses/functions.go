package losses

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/vector"
)

func MeanSquareError[T types.Float](actual []T, predicted []T) (partials []T, value T) {
	if len(actual) == 0 {
		return
	}
	partials = vector.Subtract(actual, predicted)
	value = vector.DotProduct(partials, partials) / T(len(actual))
	return
}

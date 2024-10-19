package metrics

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/utils"
	"github.com/publiczny81/ml/calculus/vector"
	"github.com/publiczny81/ml/calculus/vector/pool"
	"github.com/publiczny81/ml/utils/slices"
	"math"
)

const (
	Euclidean = "euclidean"
	Manhattan = "manhattan"
	Sum       = "sum"
)

var register = map[string]Metrics{
	Euclidean: {
		Name:     Euclidean,
		Function: EuclideanDistance[[]float64, float64],
	},
	Sum: {
		Name:     Sum,
		Function: SumDistance[[]float64, float64],
	},
	Manhattan: {
		Name:     Manhattan,
		Function: ManhattanFunc[[]float64, float64],
	},
}

type Func[S ~[]T, T types.Float] func(S, S) T

type Metrics struct {
	Name     string
	Function func([]float64, []float64) float64
}

func Get(metrics string) (m Metrics, found bool) {
	m, found = register[metrics]
	return
}

func EuclideanDistance[S ~[]T, T types.Float](x, y S) (value T) {
	var v = vector.Subtract(x, y)
	value = T(math.Sqrt(float64(vector.DotProduct(v, v))))
	pool.Put(v)
	return
}

func SumDistance[S ~[]T, T types.Float](x, y S) (value T) {
	var v = vector.Subtract(x, y)
	value = slices.Aggregate(v, value, func(acc T, e T) T {
		return acc + e
	})
	pool.Put(v)
	return
}

func ManhattanFunc[S ~[]T, T types.Float](x, y S) (result T) {
	var v = vector.Subtract(x, y)
	result = slices.AggregateWithIndex(v, result, func(acc T, i int, e T) T {
		return acc + utils.Abs(e)
	})
	pool.Put(v)
	return
}

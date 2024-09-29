package metrics

import "math"

func EuclideanDistance(x, y []float64) (value float64) {
	if len(x) != len(y) {
		panic("invalid vectors length")
	}

	for i := range x {
		value += math.Pow(x[i]-y[i], 2)
	}
	return math.Sqrt(value)
}

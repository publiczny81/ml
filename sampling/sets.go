package sampling

import (
	"context"
	"github.com/publiczny81/ml/calculus/vector"
	"github.com/publiczny81/ml/metrics"
)

// SplitSet splits source into 3 sets with given ratio.
// If ratio is not provided, then default is 0.7, 0.15, 0.15.
// If only one ratio is provided, then it is used for the first set and the rest is divided equally.

func SplitSet[E any](source Source[E], ratio ...float64) (s []Source[E], err error) {
	var (
		count      int
		start, end int
	)
	if count, err = source.Count(context.Background()); err != nil {
		return
	}
	for i, v := range normalizeRatio(ratio...) {
		if i == 2 {
			end = count
		} else {
			end += int(float64(count) * v)
		}
		s = append(s, MustNewLimitedSource(source, start, end))
		start = end
	}

	return
}

func normalizeRatio(ratio ...float64) (r []float64) {
	switch len(ratio) {
	case 0:
		r = []float64{0.7, 0.15, 0.15}
		return
	case 1:
		r = []float64{ratio[0], 1 - ratio[0]/2, 1 - ratio[0]/2}
	case 2:
		r = []float64{ratio[0], ratio[1], 1 - ratio[0] - ratio[1]}
	default:
		r = ratio[:3]
	}
	r = vector.Normalize(r, metrics.ManhattanFunc[[]float64])
	return
}

package neighbor

import (
	"github.com/publiczny81/ml/metrics"
	"math"
)

type Neighborhood func([]float64, []float64, int) float64

func (n Neighborhood) NeighborRate(me, neighbor []float64, epoch int) float64 {
	return n(me, neighbor, epoch)
}

type Radius func(int) float64

func (r Radius) Radius(epoch int) float64 {
	return r(epoch)
}

type radius interface {
	Radius(epoch int) float64
}

func Gaussian(metric metrics.Metrics, radius radius) Neighborhood {
	return func(me, neighbor []float64, epoch int) float64 {
		d := metric.Function(me, neighbor)
		r := radius.Radius(epoch)

		return math.Exp(-d * d / (2 * r * r))
	}
}

func Identity() Neighborhood {
	return func(me, neighbor []float64, _ int) float64 {
		for i, v := range me {
			if v != neighbor[i] {
				return 0
			}
		}
		return 1
	}
}

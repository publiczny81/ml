package som

import (
	"github.com/publiczny81/ml/learning"
	"github.com/publiczny81/ml/sampling"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestTrain(t *testing.T) {
	var (
		source = sampling.NewSliceSource([][]float64{
			{1, 0, 1, 0},
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 1, 1, 0},
		})
		sampler = sampling.New(source, func() sampling.Strategy[[]float64] {
			return sampling.NewSystematicalStrategyFromSource(source)
		})
		epochs     = 1
		lr         = learning.ConstantRate(0.6)
		network, _ = New(4, []int{2}, WithWeights([]float64{0.3, 0.5, 0.7, 0.2, 0.6, 0.7, 0.4, 0.3}))
		err        = Train(network, sampler, lr, new(NoneNeighbor), epochs)
		expects    = []float64{0.89, 0.08, 0.35, 0.03, 0.34, 0.95, 0.9, 0.29}
	)
	assert.NoError(t, err)
	assert.Condition(t, func() bool {
		var acceptErr = 0.005
		for i, w := range expects {
			if math.Abs(w-network.Weights[i]) > acceptErr {
				return false
			}
		}
		return true
	})
}

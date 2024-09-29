package som

import (
	"github.com/publiczny81/ml/learning"
	"github.com/publiczny81/ml/sampling"
	"github.com/stretchr/testify/assert"
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
		epochs     = 10
		network, _ = New(4, []int{2}, WithWeights([]float64{0.3, 0.5, 0.7, 0.2, 0.6, 0.5, 0.4, 0.2}))
		err        = Train(network, sampler, learning.ConstantRate(0.6), new(NoneNeighbor), epochs)
	)
	assert.NoError(t, err)
}

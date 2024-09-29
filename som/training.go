package som

import "github.com/publiczny81/ml/utils/slice"

type sampler interface {
	Next() []float64
	HasNext() bool
	Reset()
}

type learningRate interface {
	Rate(epoch int) float64
}

type neighborhoodRate interface {
	NeighborRate([]int, []int, int) float64
}

func Train(n *Network, sampler sampler, learningRate learningRate, neighborhood neighborhoodRate, epochs int) (err error) {
	var (
		train = func(learningRate float64, epoch int) (err error) {
			sampler.Reset()
			for sampler.HasNext() {
				var (
					sample = sampler.Next()
					bmu    = n.BestMatchingUnit(sample)
				)
				n.Neurons.IterateWithIndex(func(idx int, v neuron) bool {
					var factor = learningRate * neighborhood.NeighborRate(bmu, n.Neurons.Position(idx), epoch)
					if factor > 0 {
						slice.ApplyWithIndex(v, func(i int, w float64) float64 {
							return w + factor*(sample[i]-w)
						})
					}
					return true
				})
			}
			return
		}
	)

	for epoch := range epochs {
		if err = train(learningRate.Rate(epoch), epoch); err != nil {
			return
		}
	}
	return
}

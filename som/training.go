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

type neighborhood interface {
	NeighborRate([]int, []int, int) float64
}

func Train(n *Network, sampler sampler, learningRate learningRate, neighborhood neighborhood, epochs int) (err error) {
	for epoch := range epochs {
		if err = train(n, sampler, learningRate.Rate(epoch), neighborhood, epoch); err != nil {
			return
		}
	}
	return
}

func train(n *Network, sampler sampler, learningRates float64, neighborhood neighborhood, epoch int) (err error) {
	sampler.Reset()
	for sampler.HasNext() {
		if err = trainStep(n, sampler.Next(), learningRates, neighborhood, epoch); err != nil {
			return
		}
	}
	return
}

func trainStep(n *Network, sample []float64, learningRate float64, neighborhood neighborhood, epoch int) error {
	var bmu = n.BestMatchingUnit(sample)

	n.Neurons.IterateWithIndex(func(idx int, v neuron) bool {
		var factor = learningRate * neighborhood.NeighborRate(bmu, n.Neurons.Position(idx), epoch)
		slice.ApplyWithIndex(v, func(i int, w float64) float64 {
			return w + factor*(sample[i]-w)
		})
		return true
	})

	return nil
}

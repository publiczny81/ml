package som

type NoneNeighbor struct{}

func (nn NoneNeighbor) NeighborRate(me, neighborhood []int, _ int) float64 {
	for i, e := range me {
		if neighborhood[i] != e {
			return 0
		}
	}
	return 1
}

type OnlyClosestNeighbors struct {
	learningRate
}

func NewOnlyClosestNeighbors(rate learningRate) *OnlyClosestNeighbors {
	return &OnlyClosestNeighbors{
		learningRate: rate,
	}
}

func (ocn OnlyClosestNeighbors) NeighborRate(me, neighbor []int, epoch int) float64 {
	var count = 0
	for i, e := range me {
		if neighbor[i] == e {
			count++
		}
	}
	if count == len(me) {
		return 1
	}
	return ocn.Rate(epoch)
}

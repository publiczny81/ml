package mlp

//
//import (
//	"github.com/publiczny81/ml/ann/neuron"
//	"github.com/publiczny81/ml/functions"
//	"github.com/publiczny81/ml/learning"
//	"github.com/publiczny81/ml/losses"
//	"github.com/publiczny81/ml/sampling"
//	"github.com/stretchr/testify/suite"
//	"math/rand/v2"
//	"testing"
//)
//
//type BackPropagationTrainerSuite struct {
//	suite.Suite
//}
//
//func TestBackPropagationTrainer(t *testing.T) {
//	suite.Run(t, new(BackPropagationTrainerSuite))
//}
//
//type randStub struct {
//}
//
//func (s *randStub) Float() float64 {
//	return 2*rand.Float64() - 1
//}
//
//func (s *BackPropagationTrainerSuite) TestTrain() {
//	var (
//		net, err = New[float64]([]int{2, 2, 1},
//			WithActivateFunc(neuron.NewActivateFunc(functions.Sigmoid)),
//			WithRand(new(randStub)))
//		source = sampling.NewSliceSource([][][]float64{
//			{{0, 0}, {0}},
//			{{1, 0}, {1}},
//			{{0, 1}, {1}},
//			{{1, 1}, {0}},
//		})
//		epochs  = 10
//		trainer = BackPropagationTrainer[float64]{
//			sampler: sampling.New(source, func() sampling.Strategy[[][]float64] {
//				return sampling.NewSystematicalStrategyFromSource(source)
//			}),
//			learningRate: learning.LinearRateSchedule(epochs),
//			epochs:       epochs,
//			lossFunction: losses.MeanSquareError[float64],
//		}
//	)
//	s.NoError(err)
//	trainer.Train(net)
//}

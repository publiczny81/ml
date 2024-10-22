package som

import (
	"context"
	"github.com/publiczny81/ml/ann/som/neighbor"
	"github.com/publiczny81/ml/calculus/utils"
	"github.com/publiczny81/ml/learning"
	"github.com/publiczny81/ml/sampling"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TrainerSuite struct {
	suite.Suite
}

func TestTrainer(t *testing.T) {
	suite.Run(t, new(TrainerSuite))
}

func (s *TrainerSuite) TestTrain() {
	var initializerMock = new(mockInitializer)
	initializerMock.On("Initialize", mock.AnythingOfType("[]float64")).Run(func(args mock.Arguments) {
		values := []float64{0.3, 0.5, 0.7, 0.2, 0.6, 0.7, 0.4, 0.3}
		slice := args.Get(0).([]float64)

		for i := range slice {
			slice[i] = values[i]
		}
	})
	var (
		source = sampling.NewSliceSource([][]float64{
			{1, 0, 1, 0},
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 1, 1, 0},
		})
		sampler = sampling.New(source, new(sampling.SystematicalStrategy[[]float64]))
		epochs  = 1
		lr      = learning.ConstantRate(0.6)
		err     error
		network *Network

		expects = []float64{0.89, 0.08, 0.35, 0.03, 0.34, 0.95, 0.9, 0.29}
		trainer = NewTrainer(sampler, lr, neighbor.Identity(), WithInitializer(initializerMock))
	)
	network, err = New(4, []int{2})
	s.NoError(err)

	err = network.Init()
	s.NoError(err)

	err = trainer.Train(context.TODO(), network, epochs)
	s.NoError(err)

	s.Condition(func() bool {
		for i, w := range expects {
			v := utils.Round(network.Weights[i], 2)
			if v != w {
				return false
			}
		}
		return true
	})
}

type mockInitializer struct {
	mock.Mock
}

func (m *mockInitializer) Initialize(s []float64) {
	_ = m.Called(s)
}

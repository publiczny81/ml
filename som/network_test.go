package som

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NetworkSuite struct {
	suite.Suite
}

func TestNetwork(t *testing.T) {
	suite.Run(t, new(NetworkSuite))
}

func (s *NetworkSuite) TestNew() {
	var (
		tests = []struct {
			Name             string
			Factory          func() (*Network, error)
			ExpectedFeatures int
			ExpectedDim      []int
			ExpectedWeights  []float64
			ExpectedError    error
		}{
			{
				Name: "When passed invalid input then return error",
				Factory: func() (*Network, error) {
					return New(0, []int{2})
				},
				ExpectedError: InvalidParameterError,
			},
			{
				Name: "When passed invalid dim then return error",
				Factory: func() (*Network, error) {
					return New(4, nil)
				},
				ExpectedError: InvalidParameterError,
			},
			{
				Name: "When neither weights nor rand passed then return error",
				Factory: func() (*Network, error) {
					return New(4, []int{2, 1})
				},
				ExpectedError: InvalidParameterError,
			},
			{
				Name: "When passed rand then return valid network",
				Factory: func() (*Network, error) {
					var r = new(randMock)
					r.On("Float64").Return(float64(1))
					return New(4, []int{2, 2}, WithRand(r))
				},
				ExpectedFeatures: 4,
				ExpectedDim:      []int{2, 2},
				ExpectedWeights:  []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			{
				Name: "When passed weights then return valid network",
				Factory: func() (*Network, error) {
					return New(4, []int{2, 2}, WithWeights([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
				},
				ExpectedFeatures: 4,
				ExpectedDim:      []int{2, 2},
				ExpectedWeights:  []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var actual, err = test.Factory()
			if test.ExpectedError != nil {
				s.Error(err)
				s.ErrorContains(err, test.ExpectedError.Error())
				return
			}
			s.NoError(err)
			s.Equal(test.ExpectedDim, actual.Dim())
			s.Equal(test.ExpectedFeatures, actual.Features())
			s.Equal(test.ExpectedWeights, actual.Weights)
		})
	}
}

func (s *NetworkSuite) TestBestMatchingUnit() {
	var (
		network, _ = New(1, []int{2, 2, 2}, WithWeights([]float64{1, 2, 3, 4, 5, 6, 7, 8}))
		tests      = []struct {
			Name     string
			Input    []float64
			Expected []int
		}{
			{
				Name:     "When passed input closest first neuron then return position of first neuron",
				Input:    []float64{1},
				Expected: []int{0, 0, 0},
			},
			{
				Name:     "When passed input closest second neuron then return position of second neuron",
				Input:    []float64{2},
				Expected: []int{0, 0, 1},
			},
			{
				Name:     "When passed input closest eighth neuron then return position of second neuron",
				Input:    []float64{9},
				Expected: []int{1, 1, 1},
			},
		}
	)
	for _, test := range tests {
		s.Run(test.Name, func() {
			var actual = network.BestMatchingUnit(test.Input)
			s.Equal(test.Expected, actual)
		})
	}
}

type randMock struct {
	mock.Mock
}

func (r *randMock) Int() int {
	args := r.Called()
	return args.Int(0)
}

func (r *randMock) IntN(n int) int {
	args := r.Called(n)
	return args.Int(0)
}

func (r *randMock) Uint64() uint64 {
	args := r.Called()
	return args.Get(0).(uint64)
}

func (r *randMock) Float64() float64 {
	args := r.Called()
	return args.Get(0).(float64)
}

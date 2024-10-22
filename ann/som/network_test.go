package som

import (
	"github.com/publiczny81/ml/errors"
	"github.com/publiczny81/ml/metrics"
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
			Name          string
			Factory       func() (*Network, error)
			ExpectedError error
			Expected      *Network
		}{
			{
				Name: "When passed invalid topology then return error",
				Factory: func() (*Network, error) {
					return New(4, []int{2, 2, 2}, WithTopology("invalid"))
				},
				ExpectedError: errors.InvalidParameterValueError,
			},
			{
				Name: "When passed invalid metrics then return error",
				Factory: func() (*Network, error) {
					return New(4, []int{2, 2}, WithMetrics("invalid"))
				},
				ExpectedError: errors.InvalidParameterValueError,
			},
			{
				Name: "When shape and number of weights are incompatible then return error",
				Factory: func() (*Network, error) {
					return New(4, []int{2, 2}, WithWeights([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
				},
				ExpectedError: errors.InvalidParameterValueError,
			},
			{
				Name: "When all valid parameters provided then return valid network",
				Factory: func() (*Network, error) {
					return New(4, []int{2, 2}, WithTopology(TopologyHexagonal))
				},
				Expected: &Network{
					config: config{
						Features: 4,
						Shape:    []int{2, 2},
						Topology: TopologyHexagonal,
						Metrics:  metrics.Euclidean,
					},
				},
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
			s.Equal(test.Expected, actual)

		})
	}
}

func (s *NetworkSuite) TestInit() {
	var (
		tests = []struct {
			Name          string
			Factory       func() (*Network, error)
			Options       []Option
			Expected      []float64
			ExpectedError error
		}{
			{
				Name: "When passed invalid features then return error",
				Factory: func() (*Network, error) {
					return New(0, []int{2})
				},
				ExpectedError: errors.InvalidParameterValueError,
			},
			{
				Name: "When passed invalid shape then return error",
				Factory: func() (*Network, error) {
					return New(4, nil)
				},
				ExpectedError: errors.InvalidParameterValueError,
			},
			{
				Name: "When weights are not provided then initiate with zeros vector",
				Factory: func() (*Network, error) {
					return New(1, []int{6})
				},
				Expected: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			},
			{
				Name: "When weights are provided then return no error",
				Factory: func() (*Network, error) {
					return New(1, []int{6})
				},
				Options:  []Option{WithWeights([]float64{1, 2, 3, 4, 5, 6})},
				Expected: []float64{1, 2, 3, 4, 5, 6},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var (
				network, _ = test.Factory()
				err        = network.Init(test.Options...)
			)

			if test.ExpectedError != nil {
				s.Error(err)
				s.ErrorContains(err, test.ExpectedError.Error())
				return
			}
			s.NoError(err)
			s.Equal(test.Expected, network.Weights)
			s.NotEmpty(network.Neurons)
		})
	}
}

func (s *NetworkSuite) TestBestMatchingUnit() {
	var network, err = New(1, []int{6})
	s.NoError(err)

	err = network.Init(WithWeights([]float64{1, 2, 3, 4, 5, 6}))
	s.NoError(err)

	var (
		tests = []struct {
			Name     string
			Input    []float64
			Expected Point
		}{
			{
				Name:     "When passed input closest first neuron then return position of first neuron",
				Input:    []float64{0},
				Expected: Point{0},
			},
			{
				Name:     "When passed input closest second neuron then return position of second neuron",
				Input:    []float64{2},
				Expected: Point{1},
			},
			{
				Name:     "When passed input closest eighth neuron then return position of second neuron",
				Input:    []float64{9},
				Expected: Point{5},
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

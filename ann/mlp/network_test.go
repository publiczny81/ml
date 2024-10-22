package mlp

import (
	"context"
	"errors"
	"github.com/publiczny81/ml/activate"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
)

type OptionsSuite struct {
	suite.Suite
}

func TestOptions(t *testing.T) {
	suite.Run(t, new(OptionsSuite))
}

func (s *OptionsSuite) TestCountWeights() {
	var tests = []struct {
		Name     string
		Options  Options
		Expected int
	}{
		{
			Name:     "When layers are nil then return 0",
			Options:  Options{Input: 1},
			Expected: 0,
		},
		{
			Name:     "When layers are empty then return 0",
			Options:  Options{Input: 1, Layers: []LayerSpec{}},
			Expected: 0,
		},

		{
			Name:     "When input is 3 and one layer [4] then return 16",
			Options:  Options{Input: 3, Layers: []LayerSpec{{Neurons: 4}}},
			Expected: 16,
		},
		{
			Name:     "When input is 3 layers [4, 2] then return 26",
			Options:  Options{Input: 3, Layers: []LayerSpec{{Neurons: 4}, {Neurons: 2}}},
			Expected: 26,
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			s.Equal(test.Expected, test.Options.CountWeights())
		})
	}
}

type NetworkSuite struct {
	suite.Suite
}

func TestNetwork(t *testing.T) {
	suite.Run(t, new(NetworkSuite))
}

func (s *NetworkSuite) TestNew() {
	var tests = []struct {
		Name          string
		Factory       func() (*Network, error)
		ExpectedError error
		Expected      *Network
	}{
		{
			Name: "When passed invalid input then return network",
			Factory: func() (*Network, error) {
				return New(-1)
			},
			Expected: &Network{
				Options: Options{
					Input: -1,
				},
			},
		},
		{
			Name: "When passed valid layers then return valid network",
			Factory: func() (*Network, error) {
				return New(2, AddLayer(2, activate.Sigmoid))
			},
			Expected: &Network{
				Options: Options{
					Input:  2,
					Layers: []LayerSpec{{Neurons: 2, Activation: activate.Sigmoid}},
				},
			},
		},

		{
			Name: "When passed custom weights then return network",
			Factory: func() (*Network, error) {
				return New(2, AddLayer(2, activate.Sigmoid), WithWeights([]float64{1, 1, 1, 1, 1, 1, 1, 1}))
			},
			Expected: &Network{
				Options: Options{
					Input:   2,
					Layers:  []LayerSpec{{Neurons: 2, Activation: activate.Sigmoid}},
					Weights: []float64{1, 1, 1, 1, 1, 1, 1, 1},
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			var n, err = test.Factory()
			if test.ExpectedError != nil {
				s.Error(err)
				s.Equal(test.ExpectedError, err)
				return
			}
			s.NoError(err)
			s.Equal(test.Expected, n)
		})
	}
}

func (s *NetworkSuite) TestInit() {
	var tests = []struct {
		Name          string
		Factory       func() (*Network, error)
		Options       []Option
		Expected      *Network
		ExpectedError error
	}{
		{
			Name: "When init network with invalid input then return error",
			Factory: func() (*Network, error) {
				return New(-1)
			},
			ExpectedError: errors.New("network.Input=-1: invalid parameter value"),
		},
		{
			Name: "When init network with invalid layers then return error",
			Factory: func() (*Network, error) {
				return New(2)
			},
			ExpectedError: errors.New("network.Layers={}: invalid parameter value"),
		},
		{
			Name: "When init network with incorrect weights then return error",
			Factory: func() (*Network, error) {
				return New(2, AddLayer(2, activate.Sigmoid))
			},
			Options:       []Option{WithWeights([]float64{1, 2, 3, 4})},
			ExpectedError: errors.New("incompatible parameters values len(network.Weights)=4 and network.Layers=[{sigmoid 2}]: invalid parameter value"),
		},
		{
			Name: "When init network without initial weights then return no error",
			Factory: func() (*Network, error) {
				return New(2, AddLayer(3, activate.Sigmoid), AddLayer(2, activate.Sigmoid))
			},
			Expected: &Network{
				Options: Options{
					Input: 2,
					Layers: []LayerSpec{{Neurons: 3, Activation: activate.Sigmoid},
						{Neurons: 2, Activation: activate.Sigmoid}},
					Weights: []float64{
						0, 0, 0, 0, 0.0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					},
				},
				Layers: []layer{
					{
						Weights: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0},
						Input:   []float64{0, 0, 1.0},
						Output:  []float64{0, 0, 0},
					},
					{
						Weights: []float64{0, 0, 0, 0, 0, 0, 0, 0},
						Input:   []float64{0, 0, 0, 1.0},
						Output:  []float64{0, 0},
					},
				},
			},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			var (
				actual, _ = test.Factory()
				err       = actual.Init(test.Options...)
			)

			if test.ExpectedError != nil {
				s.Error(err)
				s.ErrorContains(err, test.ExpectedError.Error())
				return
			}
			s.NoError(err)
			s.Condition(func() (success bool) {
				if success = s.Equal(test.Expected.Options.Weights, actual.Options.Weights); !success {
					return
				}
				for i, l := range actual.Layers {
					if success = s.Equal(test.Expected.Layers[i].Weights, l.Weights); !success {
						return
					}
					if success = s.Equal(test.Expected.Layers[i].Input, l.Input); !success {
						return
					}
					if success = s.Equal(test.Expected.Layers[i].Output, l.Output); !success {
						return
					}
				}
				success = true
				return
			})
		})
	}
}

func (s *NetworkSuite) TestActivate() {
	var tests = []struct {
		Name          string
		Factory       func() (*Network, error)
		Input         []float64
		Expected      []float64
		ExpectedError error
	}{
		{
			Name: "When activate network with invalid input then return error",
			Factory: func() (*Network, error) {
				return New(3, AddLayer(2, activate.Sigmoid))
			},
			Input:         []float64{1, 2},
			ExpectedError: errors.New("len(input)=2): invalid parameter value"),
		},
		{
			Name: "When activate network with valid input then return output",
			Factory: func() (n *Network, err error) {
				n, err = New(2, AddLayer(3, activate.Sigmoid), AddLayer(2, activate.Sigmoid))
				err = n.Init(WithWeights([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
				return
			},
			Input:    []float64{1, 1},
			Expected: []float64{0.24190243276979864, 0.24190243276979864},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			var (
				n, _   = test.Factory()
				_, err = n.Activate(context.TODO(), test.Input)
			)
			if test.ExpectedError != nil {
				s.Error(err)
				s.ErrorContains(err, test.ExpectedError.Error())
				return
			}
			s.NoError(err)
			s.Equal(test.Expected, n.Layers[len(n.Layers)-1].Output)
		})
	}
}

func BenchmarkActivate(b *testing.B) {
	var size = 6 * 100
	var n, _ = New(100, AddLayer(600, activate.Sigmoid), AddLayer(3, activate.Sigmoid))
	var input []float64
	for i := 0; i < size; i++ {
		input = append(input, rand.Float64())
	}
	_ = n.Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = n.Activate(context.TODO(), input)
	}
}

type randMock struct {
	mock.Mock
}

func (m *randMock) Float64() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

package mlp

//
//import (
//	"github.com/publiczny81/ml/ann/neuron"
//	"github.com/publiczny81/ml/errors"
//	"github.com/publiczny81/ml/functions"
//	"github.com/stretchr/testify/mock"
//	"github.com/stretchr/testify/suite"
//	"testing"
//)
//
//type NetworkSuite struct {
//	suite.Suite
//}
//
//func TestNetwork(t *testing.T) {
//	suite.Run(t, new(NetworkSuite))
//}
//
//func (s *NetworkSuite) TestNew() {
//	var m = new(randMock)
//	m.On("Float").Return(float64(1))
//
//	var tests = []struct {
//		Name            string
//		Factory         func() (*Network[float64], error)
//		ExpectedWeights []float64
//		ExpectedNeurons []Layer[float64]
//		ExpectedError   error
//	}{
//		{
//			Name: "When passed empty layers then return error",
//			Factory: func() (*Network[float64], error) {
//				return New[float64](nil)
//			},
//			ExpectedError: errors.InvalidParameterError,
//		},
//		{
//			Name: "When passed with valid layers but with neither rand nor weights then return error",
//			Factory: func() (*Network[float64], error) {
//				return New[float64]([]int{3, 2, 1})
//			},
//			ExpectedError: errors.InvalidParameterError,
//		},
//		{
//			Name: "When passed with valid layers but with invalid weights length then return error",
//			Factory: func() (*Network[float64], error) {
//				return New([]int{3, 2, 1}, WithWeights([]float64{1, 2, 3}))
//			},
//			ExpectedError: errors.InvalidParameterError,
//		},
//		{
//			Name: "When passed with valid layers and with valid weights length then return the network",
//			Factory: func() (*Network[float64], error) {
//				return New([]int{3, 2, 1}, WithActivateFunc(neuron.NewActivateFunc(functions.Rectifier)), WithWeights([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
//			},
//			ExpectedWeights: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
//			ExpectedNeurons: []Layer[float64]{
//				{
//					&neuron.Neuron[float64]{
//						ActivateFunc: neuron.NewActivateFunc(functions.Sigmoid),
//						Weights:      []float64{1, 2, 3, 4},
//					},
//					&neuron.Neuron[float64]{
//						ActivateFunc: neuron.NewActivateFunc(functions.Sigmoid),
//						Weights:      []float64{5, 6, 7, 8},
//					},
//				},
//				{
//					&neuron.Neuron[float64]{
//						ActivateFunc: neuron.NewActivateFunc(functions.Sigmoid),
//						Weights:      []float64{9, 10, 11},
//					},
//				},
//			},
//		},
//		{
//			Name: "When passed with valid layers and rand then return the network",
//			Factory: func() (*Network[float64], error) {
//				return New[float64]([]int{3, 2, 1},
//					WithRand[float64](m),
//					WithActivateFunc(neuron.NewActivateFunc(functions.Rectifier)))
//			},
//			ExpectedWeights: []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
//			ExpectedNeurons: []Layer[float64]{
//				{
//					&neuron.Neuron[float64]{
//						ActivateFunc: neuron.NewActivateFunc(functions.Sigmoid),
//						Weights:      []float64{1, 1, 1, 1},
//					},
//					&neuron.Neuron[float64]{
//						ActivateFunc: neuron.NewActivateFunc(functions.Sigmoid),
//						Weights:      []float64{1, 1, 1, 1},
//					},
//				},
//				{
//					&neuron.Neuron[float64]{
//						ActivateFunc: neuron.NewActivateFunc(functions.Sigmoid),
//						Weights:      []float64{1, 1, 1},
//					},
//				},
//			},
//		},
//	}
//
//	for _, test := range tests {
//		s.Run(test.Name, func() {
//			var actual, err = test.Factory()
//			if test.ExpectedError != nil {
//				s.Error(err)
//				s.ErrorContains(err, test.ExpectedError.Error())
//				return
//			}
//			s.NoError(err)
//			s.NotNil(actual)
//			s.Equal(test.ExpectedWeights, actual.Weights)
//			s.Condition(func() (success bool) {
//				for i, l := range test.ExpectedNeurons {
//					if len(l) != len(actual.Neurons[i]) {
//						return
//					}
//					for j, n := range l {
//						if len(n.Weights) != len(actual.Neurons[i][j].Weights) {
//							return
//						}
//						for k, w := range n.Weights {
//							if w != actual.Neurons[i][j].Weights[k] {
//								return
//							}
//						}
//					}
//				}
//				success = true
//				return
//			})
//		})
//	}
//}
//
//func (s *NetworkSuite) TestProcessWithResult() {
//	var (
//		tests = []struct {
//			Name     string
//			Factory  func() *Network[float64]
//			Input    []float64
//			Expected [][]float64
//		}{
//			{
//				Name: "Given network with one hidden layer when passed valid input then return valid output",
//				Factory: func() (net *Network[float64]) {
//					net, _ = New([]int{2, 1},
//						WithWeights([]float64{1, 2, 3}),
//						WithActivateFunc(neuron.NewActivateFunc(func(f float64) float64 {
//							return f
//						})))
//					return
//				},
//				Input:    []float64{1, 1},
//				Expected: [][]float64{{1, 1}, {6}},
//			},
//			{
//				Name: "Given network with two hidden layers when passed valid input then return valid output",
//				Factory: func() (net *Network[float64]) {
//					net, _ = New([]int{2, 3, 2},
//						WithWeights([]float64{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}),
//						WithActivateFunc(neuron.NewActivateFunc(func(f float64) float64 {
//							return f
//						})))
//					return
//				},
//				Input:    []float64{1, 1},
//				Expected: [][]float64{{1, 1}, {3, 6, 9}, {76, 95}},
//			},
//		}
//	)
//
//	for _, test := range tests {
//		s.Run(test.Name, func() {
//			var (
//				net    = test.Factory()
//				actual = net.MakeResult(len(test.Input))
//			)
//			net.ProcessWithResult(test.Input, actual)
//			s.Equal(test.Expected, actual)
//		})
//	}
//}
//
//func (s *NetworkSuite) TestProcess() {
//	var (
//		tests = []struct {
//			Name     string
//			Factory  func() *Network[float64]
//			Input    []float64
//			Expected []float64
//		}{
//			{
//				Name: "Given network with one hidden layer when passed valid input then return valid output",
//				Factory: func() (net *Network[float64]) {
//					net, _ = New([]int{2, 1},
//						WithWeights([]float64{1, 2, 3}),
//						WithActivateFunc(neuron.NewActivateFunc(func(f float64) float64 {
//							return f
//						})))
//					return
//				},
//				Input:    []float64{1, 1},
//				Expected: []float64{6},
//			},
//			{
//				Name: "Given network with two hidden layers when passed valid input then return valid output",
//				Factory: func() (net *Network[float64]) {
//					net, _ = New([]int{2, 3, 2},
//						WithWeights([]float64{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}),
//						WithActivateFunc(neuron.NewActivateFunc(func(f float64) float64 {
//							return f
//						})))
//					return
//				},
//				Input:    []float64{1, 1},
//				Expected: []float64{76, 95},
//			},
//		}
//	)
//
//	for _, test := range tests {
//		s.Run(test.Name, func() {
//			var net = test.Factory()
//			var actual = net.Process(test.Input)
//			s.Equal(test.Expected, actual)
//		})
//	}
//}
//
//type randMock struct {
//	mock.Mock
//}
//
//func (m *randMock) Float() float64 {
//	args := m.Called()
//	return args.Get(0).(float64)
//}

package neuron

//
//import (
//	"github.com/stretchr/testify/mock"
//	"github.com/stretchr/testify/suite"
//	"testing"
//)
//
//type NeuronSuite struct {
//	suite.Suite
//}
//
//func TestNeuron(t *testing.T) {
//	suite.Run(t, new(NeuronSuite))
//}
//
//func (s *NeuronSuite) TestNew() {
//	var (
//		tests = []struct {
//			Name     string
//			Factory  func() *Neuron
//			Expected *Neuron
//		}{
//			{
//				Name: "When New without parameters then return default neuron",
//				Factory: func() *Neuron {
//					return New()
//				},
//				Expected: &Neuron{
//					Activation: DefaultActivationFunc,
//					Weights:    nil,
//				},
//			},
//			{
//				Name: "When New with input and rand then return neuron with randomized weights",
//				Factory: func() *Neuron {
//					var m = new(randMock)
//					m.On("Float64").Return(float64(1))
//					return New(WithInput(3, m))
//				},
//				Expected: &Neuron{
//					Activation: DefaultActivationFunc,
//					Weights:    []float64{1, 1, 1, 1},
//				},
//			},
//			{
//				Name: "When New with weights parameter then return neuron with the weights",
//				Factory: func() *Neuron {
//					var m = new(randMock)
//					m.On("Float64").Return(float64(1))
//					return New(WithWeights([]float64{1, 2, 3, 4}))
//				},
//				Expected: &Neuron{
//					Activation: DefaultActivationFunc,
//					Weights:    []float64{1, 2, 3, 4},
//				},
//			},
//		}
//	)
//	for _, test := range tests {
//		s.Run(test.Name, func() {
//			var actual = test.Factory()
//			s.Equal(test.Expected.Weights, actual.Weights)
//			s.NotNil(actual.Activation)
//		})
//	}
//}
//
//type randMock struct {
//	mock.Mock
//}
//
//func (m *randMock) Float64() float64 {
//	args := m.Called()
//	return args.Get(0).(float64)
//}

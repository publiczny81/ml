package neuron

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"math/rand/v2"
	"testing"
)

type NeuronSuite struct {
	suite.Suite
}

func TestNeuron(t *testing.T) {
	suite.Run(t, new(NeuronSuite))
}

func (s *NeuronSuite) TestNew() {
	var (
		actual   = New(5)
		expected = &Neuron{
			Activation: DefaultActivationFunc,
			Weights:    make([]float64, 5),
		}
	)
	s.Equal(expected.Weights, actual.Weights)
	s.Equal(expected.Bias, actual.Bias)
	s.NotNil(actual.Activation)
}

func (s *NeuronSuite) TestBuilder() {
	var (
		tests = []struct {
			Name      string
			Factory   func() *Neuron
			Condition func(t *testing.T, n *Neuron) bool
		}{
			{
				Name: "When called NewBuilder then return simple neuron",
				Factory: func() *Neuron {
					return NewBuilder(3).Build()
				},
				Condition: func(t *testing.T, n *Neuron) bool {
					assert.Equal(t, []float64{0, 0, 0}, n.Weights)
					assert.Equal(t, float64(0), n.Bias)
					assert.NotNil(t, n.Activation)
					return true
				},
			},
			{
				Name: "When called NewBuilder with ActivationFunc then new activation function assigned to neuron",
				Factory: func() *Neuron {
					return NewBuilder(5).WithActivationFunc(func(f float64) float64 {
						return 1
					}).Build()
				},
				Condition: func(t *testing.T, n *Neuron) bool {
					return n.Process([]float64{0, 0, 0, 0, 0}) == 1
				},
			},
			{
				Name: "When called NewBuilder with Bias then neuron has set bias",
				Factory: func() *Neuron {
					return NewBuilder(5).WithBias(1).Build()
				},
				Condition: func(t *testing.T, n *Neuron) bool {
					return n.Bias == 1
				},
			},
			{
				Name: "When called NewBuilder with Weights then neuron has set weights",
				Factory: func() *Neuron {
					return NewBuilder(5).WithWeights([]float64{1, 2, 3, 4, 5}).Build()
				},
				Condition: func(t *testing.T, n *Neuron) bool {
					assert.Equal(t, []float64{1, 2, 3, 4, 5}, n.Weights)
					return true
				},
			},
			{
				Name: "When called NewBuilder with RandomizedWeights then neuron has randomized weights and bias",
				Factory: func() *Neuron {
					var m = new(sourceMock)
					m.On("Uint64").Return(uint64(12))
					return NewBuilder(5).WithRandomizedWeights(rand.New(m)).Build()
				},
				Condition: func(t *testing.T, n *Neuron) bool {
					for _, w := range n.Weights {
						if w < 0 {
							return false
						}
						if w >= 1 {
							return false
						}
					}
					if n.Bias < 0 {
						return false
					}
					if n.Bias >= 1 {
						return false
					}
					return true
				},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var actual = test.Factory()
			s.True(test.Condition(s.T(), actual))
		})
	}
}

type sourceMock struct {
	mock.Mock
}

func (m *sourceMock) Uint64() uint64 {
	args := m.Called()
	return args.Get(0).(uint64)
}

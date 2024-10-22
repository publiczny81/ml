package neuron

import (
	"github.com/publiczny81/ml/errors"
	"github.com/publiczny81/ml/functions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestNewActivateFunc(t *testing.T) {
	var tests = []struct {
		Name     string
		Factory  func() ActivateFunc[[]float64, float64]
		Features []float64
		Weights  []float64
		Expected float64
		Error    error
	}{
		{
			Name: "When missing activation function then panic",
			Factory: func() ActivateFunc[[]float64, float64] {
				return NewActivateFunc[float64](nil)
			},
			Error: errors.InvalidParameterError,
		},
		{
			Name: "When length of features is equal to length of weights then process sum of multiplications",
			Factory: func() ActivateFunc[[]float64, float64] {
				return NewActivateFunc[float64](func(value float64) float64 { return value })
			},
			Features: []float64{1, 2, 3},
			Weights:  []float64{3, 2, 1},
			Expected: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.Error != nil {
				assert.Panics(t, func() {
					_ = test.Factory()
				})
				return
			}
			var actual = test.Factory()(test.Features, test.Weights)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

type NeuronSuite struct {
	suite.Suite
}

func TestNeuron(t *testing.T) {
	suite.Run(t, new(NeuronSuite))
}

func (s *NeuronSuite) TestNew() {
	var tests = []struct {
		Name            string
		Factory         func() *Neuron[float64]
		ExpectedWeights []float64
		ExpectedError   error
	}{
		{
			Name: "When missing activation function then panic",
			Factory: func() *Neuron[float64] {
				return New[float64](nil, []float64{1, 2, 3})
			},
			ExpectedError: errors.InvalidParameterError,
		},
		{
			Name: "When missing weights then panic",
			Factory: func() *Neuron[float64] {
				return New[float64](NewActivateFunc(functions.Sigmoid), nil)
			},
			ExpectedError: errors.InvalidParameterError,
		},
		{
			Name: "When all valid parameters provided then return valid neuron",
			Factory: func() *Neuron[float64] {
				return New(NewActivateFunc(functions.Sigmoid), []float64{1, 2, 3})
			},
			ExpectedWeights: []float64{1, 2, 3},
		},
	}
	for _, test := range tests {
		s.Run(test.Name, func() {
			if test.ExpectedError != nil {
				s.Panics(func() {
					_ = test.Factory()
				})
				return
			}
			var actual = test.Factory()
			s.NotNil(actual)
			s.NotNil(actual.ActivateFunc)
			s.Equal(test.ExpectedWeights, actual.Weights)
		})
	}
}

func (s *NeuronSuite) TestBuilder() {
	var tests = []struct {
		Name            string
		Factory         func() *Neuron[float64]
		ExpectedWeights []float64
		Error           error
	}{
		{
			Name: "When missing weights then panic",
			Factory: func() *Neuron[float64] {
				return NewBuilder[float64]().
					WithActivateFunc(NewActivateFunc[float64](functions.Rectifier)).Build()
			},
			Error: errors.InvalidParameterError,
		},
		{
			Name: "When missing activation function then panic",
			Factory: func() *Neuron[float64] {
				return NewBuilder[float64]().
					WithWeights([]float64{1, 2, 3}).
					Build()
			},
			Error: errors.InvalidParameterError,
		},
		{
			Name: "When missing rand then panic",
			Factory: func() *Neuron[float64] {
				return NewBuilder[float64]().
					WithFeatures(2, false, nil).
					Build()
			},
			Error: errors.InvalidParameterError,
		},
		{
			Name: "When features set to 2 then initialize neuron with two weights",
			Factory: func() *Neuron[float64] {
				var m = new(randMock)
				m.On("Float").Return(float64(1))
				return NewBuilder[float64]().
					WithActivateFunc(NewActivateFunc(functions.Rectifier)).
					WithFeatures(2, false, m).
					Build()
			},
			ExpectedWeights: []float64{1, 1},
		},
		{
			Name: "When features set to 2 with bias then initialize neuron with three weights",
			Factory: func() *Neuron[float64] {
				var m = new(randMock)
				m.On("Float").Return(float64(1))
				return NewBuilder[float64]().
					WithActivateFunc(NewActivateFunc(functions.Rectifier)).
					WithFeatures(2, true, m).
					Build()
			},
			ExpectedWeights: []float64{1, 1, 1},
		},
		{
			Name: "When weights provided then initialize neuron with these weights",
			Factory: func() *Neuron[float64] {
				var m = new(randMock)
				m.On("Float").Return(float64(1))
				return NewBuilder[float64]().
					WithActivateFunc(NewActivateFunc(functions.Rectifier)).
					WithWeights([]float64{1, 2, 3}).
					Build()
			},
			ExpectedWeights: []float64{1, 2, 3},
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			if test.Error != nil {
				s.Panics(func() {
					_ = test.Factory()
				})
				return
			}
			var neuron = test.Factory()
			s.NotNil(neuron)
			s.Equal(test.ExpectedWeights, neuron.Weights)
		})
	}
}

type randMock struct {
	mock.Mock
}

func (m *randMock) Float() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

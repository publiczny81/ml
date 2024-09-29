package sampling

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SystematicalStrategySuite struct {
	suite.Suite
}

func TestSystematicalStrategy(t *testing.T) {
	suite.Run(t, new(SystematicalStrategySuite))
}

func (s *SystematicalStrategySuite) TestNew() {
	var (
		tests = []struct {
			Name     string
			Factory  func() *SystematicalStrategy[float64]
			Expected *SystematicalStrategy[float64]
		}{
			{
				Name: "Created with limit",
				Factory: func() *SystematicalStrategy[float64] {
					return NewSystematicalStrategyWithLimit[float64](3)
				},
				Expected: &SystematicalStrategy[float64]{
					limit: 3,
				},
			},
			{
				Name: "Created with source",
				Factory: func() *SystematicalStrategy[float64] {
					m := new(sourceMock)
					m.On("Count").Return(3)
					return NewSystematicalStrategyFromSource[float64](m)
				},
				Expected: &SystematicalStrategy[float64]{
					limit: 3,
				},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var actual = test.Factory()
			s.Equal(test.Expected, actual)
		})
	}
}

func (s *SystematicalStrategySuite) TestNext() {
	var (
		tests = []struct {
			Name     string
			Strategy *SystematicalStrategy[float64]
			Source   func() *sourceMock
			Expected []float64
		}{
			{
				Name:     "When passed source with three values then all of then returned",
				Strategy: NewSystematicalStrategyWithLimit[float64](3),
				Source: func() (m *sourceMock) {
					m = new(sourceMock)
					m.On("Select", 0).Once().Return(float64(1))
					m.On("Select", 1).Once().Return(float64(2))
					m.On("Select", 2).Once().Return(float64(3))
					return
				},
				Expected: []float64{1, 2, 3},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var (
				source = test.Source()
				actual []float64
			)
			for {
				var value, found = test.Strategy.Next(source)
				if !found {
					break
				}
				actual = append(actual, value)
			}

			s.Equal(test.Expected, actual)
			source.AssertExpectations(s.T())
		})
	}
}

type RandomStrategySuite struct {
	suite.Suite
}

func TestRandomStrategy(t *testing.T) {
	var tests = []struct {
		Name     string
		Rand     func() *randMock
		Source   func() *sourceMock
		Expected float64
	}{
		{
			Name: "When called next then return next value from source and true and no more value available",
			Rand: func() (m *randMock) {
				m = new(randMock)
				m.On("IntN", 3).Return(1)
				return
			},
			Source: func() (m *sourceMock) {
				m = new(sourceMock)
				m.On("Count").Return(3)
				m.On("Select", 1).Return(2.5)
				return
			},
			Expected: 2.5,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				rand          = test.Rand()
				source        = test.Source()
				strategy      = NewRandomStrategy[float64](rand)
				actual, found = strategy.Next(source)
			)
			assert.True(t, found)
			assert.Equal(t, test.Expected, actual)

			_, found = strategy.Next(source)
			assert.False(t, found)
		})
	}
}

type randMock struct {
	mock.Mock
}

func (m *randMock) IntN(n int) int {
	args := m.Called(n)
	return args.Int(0)
}

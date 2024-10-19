package sampling

import (
	"context"
	"github.com/pkg/errors"
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

func (s *SystematicalStrategySuite) TestSamples() {
	var (
		ctx   = context.TODO()
		err   = errors.New("error")
		tests = []struct {
			Name     string
			Source   func() *sourceMock
			Expected []Sample[float64]
		}{
			{
				Name: "When source has 3 elements then strategy return all elements",
				Source: func() (m *sourceMock) {
					m = new(sourceMock)
					m.On("Count", ctx).Return(3, nil)
					m.On("Select", ctx, 0).Return(1.5, nil)
					m.On("Select", ctx, 1).Return(2.5, nil)
					m.On("Select", ctx, 2).Return(3.5, nil)
					return
				},
				Expected: []Sample[float64]{
					ValueOf(1.5),
					ValueOf(2.5),
					ValueOf(3.5),
				},
			},
			{
				Name: "When Count returns error then strategy return error",
				Source: func() (m *sourceMock) {
					m = new(sourceMock)
					m.On("Count", ctx).Return(0, err)
					return
				},
				Expected: []Sample[float64]{
					Error[float64](err),
				},
			},
			{
				Name: "When Select returns error then strategy return error",
				Source: func() (m *sourceMock) {
					m = new(sourceMock)
					m.On("Count", ctx).Return(3, nil)
					m.On("Select", ctx, 0).Return(float64(0), err)
					return
				},
				Expected: []Sample[float64]{
					Error[float64](err),
				},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var (
				source   = test.Source()
				strategy = &SystematicalStrategy[float64]{}
				actual   []Sample[float64]
			)
			for sample := range strategy.Samples(context.TODO(), source) {
				actual = append(actual, sample)
			}
			s.Equal(test.Expected, actual)
		})
	}
}

type RandomStrategySuite struct {
	suite.Suite
}

func TestRandomStrategy(t *testing.T) {
	suite.Run(t, new(RandomStrategySuite))
}

func (s *RandomStrategySuite) TestSamples() {
	var (
		ctx   = context.TODO()
		err   = errors.New("error")
		tests = []struct {
			Name     string
			Rand     func() *randMock
			Source   func() *sourceMock
			Expected []Sample[float64]
		}{
			{
				Name: "When source has 3 elements then strategy return one random element",
				Rand: func() (m *randMock) {
					m = new(randMock)
					m.On("IntN", 3).Return(1)
					return
				},
				Source: func() (m *sourceMock) {
					m = new(sourceMock)
					m.On("Count", ctx).Return(3, nil)
					m.On("Select", ctx, 1).Return(2.5, nil)
					return
				},
				Expected: []Sample[float64]{
					ValueOf(2.5),
				},
			},
			{
				Name: "When Count returns error then strategy return error",
				Rand: func() (m *randMock) {
					m = new(randMock)
					m.On("IntN", 3).Return(1)
					return
				},
				Source: func() (m *sourceMock) {
					m = new(sourceMock)
					m.On("Count", ctx).Return(0, err)
					return
				},
				Expected: []Sample[float64]{
					Error[float64](err),
				},
			},
			{
				Name: "When Select returns error then strategy return error",
				Rand: func() (m *randMock) {
					m = new(randMock)
					m.On("IntN", 3).Return(1)
					return
				},
				Source: func() (m *sourceMock) {
					m = new(sourceMock)
					m.On("Count", ctx).Return(3, nil)
					m.On("Select", ctx, 1).Return(float64(0), err)
					return
				},
				Expected: []Sample[float64]{
					Error[float64](err),
				},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var (
				rand     = test.Rand()
				source   = test.Source()
				strategy = NewRandomStrategy[float64](rand)
				actual   []Sample[float64]
			)
			for sample := range strategy.Samples(context.TODO(), source) {
				actual = append(actual, sample)
			}
			s.Equal(test.Expected, actual)
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

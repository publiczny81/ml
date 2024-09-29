package sampling

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SamplerSuite struct {
	suite.Suite
}

func TestSampler(t *testing.T) {
	suite.Run(t, new(SamplerSuite))
}

func (s *SamplerSuite) TestNew() {
	var sampler = New[float64](new(sourceMock), func() Strategy[float64] {
		return new(strategyMock)
	})
	s.NotNil(sampler)
	s.NotNil(sampler.newStrategy)
	s.NotNil(sampler.source)
}

func (s *SamplerSuite) TestReset() {
	var (
		strategy = func() (m *strategyMock) {
			m = new(strategyMock)
			m.On("Next", mock.AnythingOfType("*sampling.sourceMock")).Once().Return(float64(1), true)
			m.On("Next", mock.AnythingOfType("*sampling.sourceMock")).Once().Return(float64(-1), false)
			return
		}()
		actual = New[float64](new(sourceMock), func() Strategy[float64] {
			return strategy
		})
	)
	actual.Reset()

	s.True(actual.HasNext())
	s.Equal(float64(1), actual.Next())
	s.False(actual.HasNext())
	strategy.AssertExpectations(s.T())
}

func (s *SamplerSuite) TestHasNext() {
	var (
		strategy = func() (m *strategyMock) {
			m = new(strategyMock)
			m.On("Next", mock.AnythingOfType("*sampling.sourceMock")).Once().Return(float64(1), true)
			m.On("Next", mock.AnythingOfType("*sampling.sourceMock")).Once().Return(float64(-1), false)
			return
		}()
		actual = New[float64](new(sourceMock), func() Strategy[float64] {
			return strategy
		})
	)
	actual.Reset()

	s.True(actual.HasNext())
	s.Equal(float64(1), actual.Next())
	s.False(actual.HasNext())
	strategy.AssertExpectations(s.T())
}

func (s *SamplerSuite) TestNext() {
	var (
		strategy = func() (m *strategyMock) {
			m = new(strategyMock)
			m.On("Next", mock.AnythingOfType("*sampling.sourceMock")).Once().Return(float64(1), true)
			m.On("Next", mock.AnythingOfType("*sampling.sourceMock")).Once().Return(float64(-1), false)
			return
		}()
		actual = New[float64](new(sourceMock), func() Strategy[float64] {
			return strategy
		})
	)
	actual.Reset()

	s.True(actual.HasNext())
	s.Equal(float64(1), actual.Next())
	s.False(actual.HasNext())
	strategy.AssertExpectations(s.T())
}

type sourceMock struct {
	mock.Mock
}

func (s *sourceMock) Count() int {
	args := s.Called()
	return args.Int(0)
}

func (s *sourceMock) Select(i int) float64 {
	args := s.Called(i)
	return args.Get(0).(float64)
}

type strategyMock struct {
	mock.Mock
}

func (s *strategyMock) Next(source Source[float64]) (value float64, found bool) {
	args := s.Called(source)
	value = args.Get(0).(float64)
	found = args.Bool(1)
	return
}

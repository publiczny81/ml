package sampling

import (
	"context"
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
	var sampler = New[float64](new(sourceMock), new(strategyMock))
	s.NotNil(sampler)
	s.NotNil(sampler.source)
}

func (s *SamplerSuite) TestSamples() {
	var (
		ctx      = context.TODO()
		source   = new(sourceMock)
		strategy = new(strategyMock)
		sampler  = New[float64](source, strategy)
		ch       = make(chan Sample[float64])
	)
	strategy.On("Samples", ctx, source).Return(ch)
	sampler.Samples(ctx)
	strategy.AssertExpectations(s.T())
	close(ch)
}

type sourceMock struct {
	mock.Mock
}

func (s *sourceMock) Count(ctx context.Context) (int, error) {
	args := s.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (s *sourceMock) Select(ctx context.Context, i int) (float64, error) {
	args := s.Called(ctx, i)
	return args.Get(0).(float64), args.Error(1)
}

type strategyMock struct {
	mock.Mock
}

func (s *strategyMock) Samples(ctx context.Context, source Source[float64]) <-chan Sample[float64] {
	args := s.Called(ctx, source)
	return args.Get(0).(chan Sample[float64])
}

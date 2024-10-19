package pool

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PoolSuite struct {
	suite.Suite
}

func TestPool(t *testing.T) {
	suite.Run(t, new(PoolSuite))
}

func (s *PoolSuite) TestNew() {
	var (
		actual = New(func() []float64 {
			return make([]float64, 0, 10)
		})
		value    = actual.Get()
		expected = make([]float64, 0, 10)
	)
	s.Equal(expected, value)
	s.Equal(10, cap(value))
	actual.Put(value)
}

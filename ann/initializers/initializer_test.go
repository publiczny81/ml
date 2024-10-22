package initializers

import (
	"github.com/stretchr/testify/suite"
	rand2 "math/rand"
	"testing"
	"time"
)

type InitializerSuite struct {
	suite.Suite
}

func TestInitializer(t *testing.T) {
	suite.Run(t, new(InitializerSuite))
}

func (s *InitializerSuite) TestDistribution() {
	var (
		actual = &Initializer{
			distribution: Normal,
			rand: func() float64 {
				return 1
			},
		}
	)
	s.Equal(Normal, actual.Distribution())
}

func (s *InitializerSuite) TestInitialize() {
	var (
		slice  = make([]float64, 10)
		actual = &Initializer{
			distribution: Normal,
			rand: func() float64 {
				return 1
			},
		}
	)
	actual.Initialize(slice)
	for i := range slice {
		s.Equal(1.0, slice[i])
	}
}

func (s *InitializerSuite) TestConstructors() {
	var tests = []struct {
		Name    string
		Factory func() *Initializer
	}{
		{
			Name: Normal,
			Factory: func() *Initializer {
				return NewNormal(rand2.New(rand2.NewSource(time.Now().UnixNano())))
			},
		},
		{
			Name: Uniform,
			Factory: func() *Initializer {
				return NewUniform(rand2.New(rand2.NewSource(time.Now().UnixNano())))
			},
		},
		{
			Name: GlorotNormal,
			Factory: func() *Initializer {
				return NewGlorotNormal(rand2.New(rand2.NewSource(time.Now().UnixNano())), 1, 1)
			},
		},
		{
			Name: GlorotUniform,
			Factory: func() *Initializer {
				return NewGlorotUniform(rand2.New(rand2.NewSource(time.Now().UnixNano())), 1, 1)
			},
		},
		{
			Name: HeNormal,
			Factory: func() *Initializer {
				return NewHeNormal(rand2.New(rand2.NewSource(time.Now().UnixNano())), 1)
			},
		},
		{
			Name: HeUniform,
			Factory: func() *Initializer {
				return NewHeUniform(rand2.New(rand2.NewSource(time.Now().UnixNano())), 1)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			var actual = test.Factory()
			s.NotNil(actual)
			s.Equal(test.Name, actual.Distribution())
			s.NotNil(actual.rand)
		})
	}
}

package initializers

import (
	"math"
)

const (
	Normal        = "normal"
	Uniform       = "uniform"
	GlorotNormal  = "glorot_normal"
	GlorotUniform = "glorot_uniform"
	HeNormal      = "he_normal"
	HeUniform     = "he_uniform"
)

type uniform interface {
	// Float64 returns a uniform random float64 in [0, 1)
	Float64() float64
}

type normal interface {
	// NormFloat64 returns a normally distributed random float64 with mean 0 and standard deviation 1
	NormFloat64() float64
}

type rand func() float64

type Initializer struct {
	distribution string
	rand         rand
}

func (ir *Initializer) Distribution() string {
	return ir.distribution
}

func (ir *Initializer) Initialize(s []float64) {
	for i := range s {
		s[i] = ir.rand()
	}
}

func NewNormal(n normal) *Initializer {
	return &Initializer{
		distribution: Normal,
		rand: func() float64 {
			return n.NormFloat64()
		},
	}
}

func NewUniform(u uniform) *Initializer {
	return &Initializer{
		distribution: Uniform,
		rand: func() float64 {
			return 2*u.Float64() - 1
		},
	}
}

func NewGlorotNormal(n normal, input, output int) *Initializer {
	stdDev := math.Sqrt(2.0 / float64(input+output))
	return &Initializer{
		distribution: GlorotNormal,
		rand: func() float64 {
			return n.NormFloat64() * stdDev
		},
	}
}

func NewGlorotUniform(u uniform, input, output int) *Initializer {
	factor := math.Sqrt(6.0 / float64(input+output))
	return &Initializer{
		distribution: GlorotUniform,
		rand: func() float64 {
			return factor * (u.Float64()*2 - 1)
		},
	}
}

func NewHeNormal(n normal, input int) *Initializer {
	stdDev := math.Sqrt(2.0 / float64(input))
	return &Initializer{
		distribution: HeNormal,
		rand: func() float64 {
			return n.NormFloat64() * stdDev
		},
	}
}

func NewHeUniform(u uniform, input int) *Initializer {
	factor := math.Sqrt(6.0 / float64(input))
	return &Initializer{
		distribution: HeUniform,
		rand: func() float64 {
			return factor * (u.Float64()*2 - 1)
		},
	}
}

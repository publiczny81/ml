package utils

import (
	"math/rand/v2"
)

var Rand = new(randomizer)

type randomizer struct{}

func (r *randomizer) Float64() float64 {
	return rand.Float64()*2 - 1
}

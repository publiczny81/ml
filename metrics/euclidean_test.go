package metrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			X        []float64
			Y        []float64
			Expected float64
		}{
			{
				Name:     "one-dimension vectors",
				X:        []float64{0},
				Y:        []float64{1},
				Expected: 1,
			},
			{
				Name:     "two-dimension vectors",
				X:        []float64{0, 0},
				Y:        []float64{1, 1},
				Expected: 1.4142135623730951,
			},
		}
	)
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = EuclideanDistance(test.X, test.Y)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestEuclideanDistancePanic(t *testing.T) {

	assert.Panics(t, func() {
		_ = EuclideanDistance([]float64{0}, []float64{0, 1})
	}, "invalid vectors length")
}

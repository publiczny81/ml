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

func TestSum(t *testing.T) {
	var tests = []struct {
		x, y []float64
		want float64
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, 0},
		{[]float64{1, 2, 3}, []float64{1, 2, 4}, -1},
		{[]float64{1, 2, 3}, []float64{0, 0, 0}, 6},
	}

	for _, test := range tests {
		var actual = SumDistance(test.x, test.y)
		assert.Equal(t, test.want, actual)
	}
}

func TestGet(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Metrics  string
			Expected Metrics
			Found    bool
		}{
			{
				Name:     "euclidean",
				Metrics:  Euclidean,
				Expected: register[Euclidean],
				Found:    true,
			},
			{
				Name:     "manhattan",
				Metrics:  Manhattan,
				Expected: register[Manhattan],
				Found:    true,
			},
			//{
			//	Name:     "sum",
			//	Metrics:  Sum,
			//	Expected: register[Sum],
			//	Found:    true,
			//},
			{
				Name:  "not found",
				Found: false,
			},
		}
	)
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual, found = Get(test.Metrics)
			assert.Equal(t, test.Expected.Name, actual.Name)
			assert.Equal(t, test.Found, found)
		})
	}
}

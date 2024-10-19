package operations

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/vector/constructors"
	"github.com/publiczny81/ml/calculus/vector/pool"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		Name   string
		V1     []float64
		V2     []float64
		Result []float64
	}{
		{
			Name:   "When vectors have the same size then return their sum",
			V1:     []float64{1.0, 2.0, 3.0},
			V2:     []float64{4.0, 5.0, 6.0},
			Result: []float64{5.0, 7.0, 9.0},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = constructors.CopyOf(test.V1)
			constructors.Wrap(actual).Apply(Add(test.V2))
			assert.Equal(t, test.Result, actual)
		})
	}
}

func TestAddPanics(t *testing.T) {
	var tests = []struct {
		Name  string
		V1    []float64
		V2    []float64
		Error string
	}{
		{
			Name:  "When vectors have different sizes then panic",
			V1:    []float64{1.0, 2.0},
			V2:    []float64{1.0},
			Error: "unmatched size of vectors",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.PanicsWithError(t, test.Error, func() {
				v := constructors.CopyOf(test.V1)
				constructors.Wrap(v).Apply(Add(test.V2))
			})
		})
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		Name   string
		V1     []float64
		V2     []float64
		Result []float64
	}{
		{
			Name:   "When vectors have the same size then return their difference",
			V1:     []float64{1.0, 2.0, 3.0},
			V2:     []float64{4.0, 5.0, 6.0},
			Result: []float64{-3.0, -3.0, -3.0},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = pool.Get[[]float64](len(test.V1))
			types.V[float64](actual).Apply(Subtract(test.V1, test.V2))
			assert.Equal(t, test.Result, actual)
		})
	}
}

func TestSubtractPanics(t *testing.T) {
	var tests = []struct {
		Name  string
		V1    []float64
		V2    []float64
		Error string
	}{
		{
			Name:  "When vectors have different sizes then panic",
			V1:    []float64{1.0, 2.0},
			V2:    []float64{1.0},
			Error: "unmatched size of vectors",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.PanicsWithError(t, test.Error, func() {
				types.V[float64](pool.Get[[]float64](len(test.V1))).Apply(Subtract(test.V1, test.V2))
			})
		})
	}
}

func TestMultiply(t *testing.T) {
	var tests = []struct {
		Name   string
		V      []float64
		C      float64
		Result []float64
	}{
		{
			Name:   "When vector is multiplied by a constant then return the result",
			V:      []float64{1.0, 2.0, 3.0},
			C:      2.0,
			Result: []float64{2.0, 4.0, 6.0},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = pool.Get[[]float64](len(test.V))
			copy(actual, test.V)
			types.V[float64](actual).Apply(Multiply(test.C))
			assert.Equal(t, test.Result, actual)
		})
	}
}

func TestNormalize(t *testing.T) {
	var metrics = func(v1, v2 []float64) (result float64) {
		for i, e := range v2 {
			result += (v1[i] - e) * (v1[i] - e)
		}
		result = math.Sqrt(result)
		return
	}

	var tests = []struct {
		Name   string
		V      []float64
		Result []float64
	}{
		{
			Name:   "When vector is normalized then return the result",
			V:      []float64{1.0, 2.0, 3.0},
			Result: []float64{0.2672612419124244, 0.5345224838248488, 0.8017837257372732},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = pool.Get[[]float64](len(test.V))
			copy(actual, test.V)
			types.V[float64](actual).Apply(Normalize(metrics))
			assert.InDeltaSlice(t, test.Result, actual, 1e-14)
		})
	}
}

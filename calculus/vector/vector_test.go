package vector

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestZeros(t *testing.T) {
	var tests = []struct {
		Name   string
		Length int
	}{
		{
			Name:   "When length is 0 then new vector is empty",
			Length: 0,
		},
		{
			Name:   "When length is 3 then new vector has 3 elements",
			Length: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			v := Zeros[float64](test.Length)
			assert.Equal(t, test.Length, len(v))
		})
	}
}

func TestWrap(t *testing.T) {
	var tests = []struct {
		Name string
		S    []float64
	}{
		{
			Name: "When slice is empty then new vector is empty",
			S:    []float64{},
		},
		{
			Name: "When slice has 3 elements then new vector has 3 elements",
			S:    []float64{1.0, 2.0, 3.0},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			v := Wrap(test.S)
			assert.Equal(t, types.V[float64](test.S), v)
		})
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		Name     string
		V1       types.V[float64]
		V2       types.V[float64]
		Expected types.V[float64]
		Error    string
	}{
		{
			Name:     "When adding two vectors then return sum of the vectors",
			V1:       types.V[float64]{1.0, 2.0, 3.0},
			V2:       types.V[float64]{4.0, 5.0, 6.0},
			Expected: types.V[float64]{5.0, 7.0, 9.0},
		},
		{
			Name:     "When adding two empty vectors then return empty vector",
			V1:       types.V[float64]{},
			V2:       types.V[float64]{},
			Expected: types.V[float64]{},
		},
		{
			Name:  "When adding two vectors of different lengths then panic",
			V1:    types.V[float64]{1.0, 2.0},
			V2:    types.V[float64]{1.0},
			Error: "unmatched size of vectors",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.Error != "" {
				assert.PanicsWithError(t, test.Error, func() {
					_ = Add(test.V1, test.V2)
				})
				return
			}
			assert.Equal(t, test.Expected, Add(test.V1, test.V2))
		})
	}
}

func TestApply(t *testing.T) {
	var tests = []struct {
		Name     string
		Actual   types.V[float64]
		Function func(int, float64) float64
		Expected types.V[float64]
	}{
		{
			Name:   "When function called on each element then each element updated",
			Actual: types.V[float64]{1.0, 2.0, 3.0},
			Function: func(i int, e float64) float64 {
				return e + 1.0
			},
			Expected: types.V[float64]{2.0, 3.0, 4.0},
		},
		{
			Name:     "When function is nil then no change",
			Actual:   types.V[float64]{1.0, 2.0, 3.0},
			Function: nil,
			Expected: types.V[float64]{1.0, 2.0, 3.0},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			Apply(test.Actual, test.Function)
			assert.Equal(t, test.Expected, test.Actual)
		})
	}
}

func TestDotProduct(t *testing.T) {
	var tests = []struct {
		Name     string
		V1       types.V[float64]
		V2       types.V[float64]
		Expected float64
		Error    string
	}{
		{
			Name:     "When vectors are 3-element then return dot product",
			V1:       types.V[float64]{1.0, 2.0, 3.0},
			V2:       types.V[float64]{4.0, 5.0, 6.0},
			Expected: 32.0,
		},
		{
			Name:     "When vectors are empty then return 0",
			V1:       types.V[float64]{},
			V2:       types.V[float64]{},
			Expected: 0.0,
		},
		{
			Name:  "When vectors have different lengths then panic",
			V1:    types.V[float64]{1.0, 2.0},
			V2:    types.V[float64]{1.0},
			Error: "unmatched size of vectors",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.Error != "" {
				assert.PanicsWithError(t, test.Error, func() {
					_ = DotProduct(test.V1, test.V2)
				})
				return
			}
			assert.Equal(t, test.Expected, DotProduct(test.V1, test.V2))
		})
	}
}
func TestResize(t *testing.T) {
	var tests = []struct {
		Name     string
		Resize   func() types.V[float64]
		Expected types.V[float64]
	}{
		{
			Name: "When when resizing to 0 then new vector is empty",
			Resize: func() types.V[float64] {
				return Resize(types.V[float64]{1.0, 2.0, 3.0}, 0)
			},
			Expected: types.V[float64]{},
		},
		{
			Name: "When resizing to 3 then new vector has 3 elements",
			Resize: func() types.V[float64] {
				return Resize(types.V[float64]{1.0, 2.0, 3.0, 4.0}, 3)
			},
			Expected: types.V[float64]{1.0, 2.0, 3.0},
		},
		{
			Name: "When resizing to 5 then new vector has 5 elements",
			Resize: func() types.V[float64] {
				return Resize(types.V[float64]{1.0, 2.0, 3.0}, 5)
			},
			Expected: types.V[float64]{1.0, 2.0, 3.0, 0.0, 0.0},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			v := test.Resize()
			assert.Equal(t, test.Expected, v)
		})
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		Name     string
		V1       types.V[float64]
		V2       types.V[float64]
		Expected types.V[float64]
		Error    string
	}{
		{
			Name:     "When subtracting two vectors then return difference of the vectors",
			V1:       types.V[float64]{1.0, 2.0, 3.0},
			V2:       types.V[float64]{4.0, 5.0, 6.0},
			Expected: types.V[float64]{-3.0, -3.0, -3.0},
		},
		{
			Name:     "When subtracting two empty vectors then return empty vector",
			V1:       types.V[float64]{},
			V2:       types.V[float64]{},
			Expected: nil,
		},
		{
			Name:  "When subtracting two vectors of different lengths then panic",
			V1:    types.V[float64]{1.0, 2.0},
			V2:    types.V[float64]{1.0},
			Error: "unmatched size of vectors",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.Error != "" {
				assert.PanicsWithError(t, test.Error, func() {
					_ = Subtract(test.V1, test.V2)
				})
				return
			}
			assert.Equal(t, test.Expected, Subtract(test.V1, test.V2))
		})
	}
}

func TestMultiply2(t *testing.T) {
	var tests = []struct {
		Name     string
		V        types.V[float64]
		C        float64
		Expected types.V[float64]
	}{
		{
			Name:     "When multiplying vector by 2 then return vector with elements doubled",
			V:        types.V[float64]{1.0, 2.0, 3.0},
			C:        2.0,
			Expected: types.V[float64]{2.0, 4.0, 6.0},
		},
		{
			Name:     "When multiplying empty vector by 2 then return empty vector",
			V:        types.V[float64]{},
			C:        2.0,
			Expected: []float64{},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expected, Multiply(test.V, test.C))
		})
	}
}

func TestNormalize(t *testing.T) {
	var metrics = func(v1, v2 types.V[float64]) (result float64) {
		for i, e := range v2 {
			result += (v1[i] - e) * (v1[i] - e)
		}
		return math.Sqrt(result)
	}
	var tests = []struct {
		Name     string
		V        types.V[float64]
		Expected types.V[float64]
	}{
		{
			Name:     "When normalizing vector then return vector with length 1",
			V:        types.V[float64]{1.0, 2.0, 3.0},
			Expected: types.V[float64]{0.2672612419124244, 0.5345224838248488, 0.8017837257372732},
		},
		{
			Name:     "When normalizing empty vector then return empty vector",
			V:        types.V[float64]{},
			Expected: types.V[float64]{},
		},
		{
			Name:     "When normalizing vector with 0 length then return same vector",
			V:        types.V[float64]{0.0, 0.0, 0.0},
			Expected: types.V[float64]{0.0, 0.0, 0.0},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expected, Normalize(test.V, metrics))
		})
	}
}

func TestExclude(t *testing.T) {
	var tests = []struct {
		Name     string
		V        types.V[float64]
		Index    int
		Expected types.V[float64]
	}{
		{
			Name:     "When excluding element at index 1 then return vector without element at index 1",
			V:        types.V[float64]{1.0, 2.0, 3.0},
			Index:    1,
			Expected: types.V[float64]{1.0, 3.0},
		},
		{
			Name:     "When excluding element at index 0 then return vector without element at index 0",
			V:        types.V[float64]{1.0, 2.0, 3.0},
			Index:    0,
			Expected: types.V[float64]{2.0, 3.0},
		},
		{
			Name:     "When excluding element at index 2 then return vector without element at index 2",
			V:        types.V[float64]{1.0, 2.0, 3.0},
			Index:    2,
			Expected: types.V[float64]{1.0, 2.0},
		},
		{
			Name:     "When excluding element at index -1 then return same vector",
			V:        types.V[float64]{1.0, 2.0, 3.0},
			Index:    -1,
			Expected: types.V[float64]{1.0, 2.0, 3.0},
		},
		{
			Name:     "When excluding element at index 3 then return same vector",
			V:        types.V[float64]{1.0, 2.0, 3.0},
			Index:    3,
			Expected: types.V[float64]{1.0, 2.0, 3.0},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expected, Exclude(test.V, test.Index))
		})
	}
}

package concurrent

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForEachColumn(t *testing.T) {
	var (
		m         = zeros[float64](1000)
		operation = ForEachColumn[float64](func(j int, c []float64) []float64 {
			for i := range c {
				c[i] = float64(j) * 2
			}
			return c
		})
	)
	m.Apply(operation)
	assert.Condition(t, func() (success bool) {
		for i := range m {
			for j, v := range m[i] {
				if v != float64(j*2) {
					return false
				}
			}
		}
		return true
	})
}

func TestForEachRow(t *testing.T) {
	var (
		m         = zeros[float64](1000)
		operation = ForEachRow[float64](func(j int, r []float64) []float64 {
			for i := range r {
				r[i] = float64(j) * 2
			}
			return r
		})
	)
	m.Apply(operation)
	assert.Condition(t, func() (success bool) {
		for i := range m {
			for _, v := range m[i] {
				if v != float64(i*2) {
					return false
				}
			}
		}
		return true
	})
}

func zeros[T types.Real](shape ...int) (m types.M[T]) {
	switch len(shape) {
	case 0:
		panic(errors.InvalidSizeOfMatrixError)
	case 1:
		shape = append(shape, shape[0])
	}
	var (
		data       = make([]T, shape[0]*shape[1])
		start, end = 0, shape[1]
	)
	m = make(types.M[T], shape[0])
	for i := range m {
		m[i] = data[start:end]
		start = end
		end += shape[1]
	}
	return
}

func TestAdd(t *testing.T) {
	var (
		m      = zeros[float64](10)
		actual = zeros[float64](10)
	)
	m[2][2] = 3

	actual.Apply(Add(m))
	assert.Condition(t, func() (success bool) {
		actual[2][2] = 3
		return true
	})
}

func TestSubtract(t *testing.T) {
	var (
		m      = zeros[float64](10)
		actual = zeros[float64](10)
	)
	m[2][2] = 3

	actual.Apply(Subtract(m))
	assert.Condition(t, func() (success bool) {
		actual[2][2] = -3
		return true
	})
}

func TestMultiply(t *testing.T) {
	var (
		actual = zeros[float64](10)
	)
	actual[2][2] = 3

	actual.Apply(Multiply(float64(2)))
	assert.Condition(t, func() (success bool) {
		actual[2][2] = 6
		return true
	})
}

func TestTranspose(t *testing.T) {
	var tests = []struct {
		Name     string
		M1       types.M[int]
		Actual   types.M[int]
		Expected types.M[int]
	}{
		{
			Name:     "When matrices are 2x2 then return transposed matrix",
			M1:       types.M[int]{{1, 2}, {3, 4}},
			Actual:   zeros[int](2, 2),
			Expected: types.M[int]{{1, 3}, {2, 4}},
		},
		{
			Name:     "When matrices are 3x3 then return transposed matrix",
			M1:       types.M[int]{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			Actual:   zeros[int](3, 3),
			Expected: types.M[int]{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			Name:     "When matrices are 3x2 then return transposed matrix",
			M1:       types.M[int]{{1, 2}, {3, 4}, {5, 6}},
			Actual:   zeros[int](2, 3),
			Expected: types.M[int]{{1, 3, 5}, {2, 4, 6}},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Actual.Apply(Transpose(test.M1))
			assert.Equal(t, test.Expected, test.Actual)
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	var (
		m      = zeros[float64](10000)
		actual = zeros[float64](10000)
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual.Apply(Add(m))
	}
}

func BenchmarkAddConcurrent(b *testing.B) {
	var (
		m      = zeros[float64](10000)
		actual = zeros[float64](10000)
	)
	b.ResetTimer()
	for range b.N {
		actual.Apply(Add(m))
	}
}

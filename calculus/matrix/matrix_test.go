package matrix

import (
	"github.com/publiczny81/ml/calculus/matrix/concurrent"
	"github.com/publiczny81/ml/calculus/types"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSetValues(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Values   types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When setting values of 2x2 matrix then return matrix with values set",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Values:   [][]float64{{5, 6}, {7, 8}},
				Expected: types.M[float64]{{5, 6}, {7, 8}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			SetValues(test.Values)(test.M)
			assert.Equal(t, test.Expected, test.M)
		})
	}
}

func TestForEachRow(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When applying operation to each row of 2x2 matrix then return matrix with operation applied",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Expected: types.M[float64]{{2, 4}, {6, 8}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			concurrent.ForEachRow[float64](func(i int, r []float64) []float64 {
				for j := range r {
					r[j] *= 2
				}
				return r
			})(test.M)
			assert.Equal(t, test.Expected, test.M)
		})
	}
}

func TestForEachColumn(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When applying operation to each column of 2x2 matrix then return matrix with operation applied",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Expected: types.M[float64]{{2, 4}, {6, 8}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ForEachColumn[float64](func(j int, c []float64) []float64 {
				for i := range c {
					c[i] *= 2
				}
				return c
			})(test.M)
			assert.Equal(t, test.Expected, test.M)
		})
	}
}

func TestSwapRows(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Row1     int
			Row2     int
			Expected types.M[float64]
		}{
			{
				Name:     "When swapping rows of 2x2 matrix then return matrix with rows swapped",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Row1:     0,
				Row2:     1,
				Expected: types.M[float64]{{3, 4}, {1, 2}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			SwapRows[float64](test.Row1, test.Row2)(test.M)
			assert.Equal(t, test.Expected, test.M)
		})
	}
}

func TestSwapColumns(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Col1     int
			Col2     int
			Expected types.M[float64]
		}{
			{
				Name:     "When swapping columns of 2x2 matrix then return matrix with columns swapped",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Col1:     0,
				Col2:     1,
				Expected: types.M[float64]{{2, 1}, {4, 3}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			SwapColumns[float64](test.Col1, test.Col2)(test.M)
			assert.Equal(t, test.Expected, test.M)
		})
	}
}

func TestAdd(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M1       types.M[float64]
			M2       types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When matrices are 2x2 then result is 2x2 with elements sum",
				M1:       types.M[float64]{{1, 2}, {3, 4}},
				M2:       types.M[float64]{{5, 6}, {7, 8}},
				Expected: types.M[float64]{{6, 8}, {10, 12}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := Add(test.M1, test.M2)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestSubtract(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M1       types.M[float64]
			M2       types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When matrices are 2x2 then result is 2x2 with elements difference",
				M1:       types.M[float64]{{1, 2}, {3, 4}},
				M2:       types.M[float64]{{5, 6}, {7, 8}},
				Expected: types.M[float64]{{-4, -4}, {-4, -4}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := Subtract(test.M1, test.M2)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestMultiply(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			C        float64
			Expected types.M[float64]
		}{
			{
				Name:     "When multiplying matrix by 2 then all elements are doubled",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				C:        2,
				Expected: types.M[float64]{{2, 4}, {6, 8}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := Multiply(test.M, test.C)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestProduct(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M1       types.M[float64]
			M2       types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When multiplying 1x1 by 1x1 then result is 1x1 with elements product",
				M1:       types.M[float64]{{1}},
				M2:       types.M[float64]{{2}},
				Expected: types.M[float64]{{2}},
			},
			{
				Name:     "When multiplying 1x2 by 2x1 then result is 1x1 with elements product",
				M1:       types.M[float64]{{1, 2}},
				M2:       types.M[float64]{{3}, {4}},
				Expected: types.M[float64]{{11}},
			},
			{
				Name:     "When multiplying 2x1 by 1x2 then result is 2x2 with elements product",
				M1:       types.M[float64]{{1}, {2}},
				M2:       types.M[float64]{{3, 4}},
				Expected: types.M[float64]{{3, 4}, {6, 8}},
			},
			{
				Name:     "When multiplying 2x2 by 2x2 then result is 2x2 with elements product",
				M1:       types.M[float64]{{1, 2}, {3, 4}},
				M2:       types.M[float64]{{5, 6}, {7, 8}},
				Expected: types.M[float64]{{19, 22}, {43, 50}},
			},
			{
				Name: "When multiplying 2x3 by 3x2 then result is 2x2 with elements product",
				M1:   types.M[float64]{{1, 2, 3}, {4, 5, 6}},
				M2:   types.M[float64]{{7, 8}, {9, 10}, {11, 12}},
				Expected: types.M[float64]{
					{58, 64},
					{139, 154},
				},
			},
			{
				Name: "When multiplying 3x2 by 2x3 then result is 3x3 with elements product",
				M1:   types.M[float64]{{1, 2}, {3, 4}, {5, 6}},
				M2:   types.M[float64]{{7, 8, 9}, {10, 11, 12}},
				Expected: types.M[float64]{
					{27, 30, 33},
					{61, 68, 75},
					{95, 106, 117},
				},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := Product(test.M1, test.M2)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestTranspose(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When transposing 1x1 matrix then result is 1x1 with elements transposed",
				M:        types.M[float64]{{1}},
				Expected: types.M[float64]{{1}},
			},
			{
				Name:     "When transposing 2x2 matrix then result is 2x2 with elements transposed",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Expected: types.M[float64]{{1, 3}, {2, 4}},
			},
			{
				Name:     "When transposing 2x3 matrix then result is 3x2 with elements transposed",
				M:        types.M[float64]{{1, 2, 3}, {4, 5, 6}},
				Expected: types.M[float64]{{1, 4}, {2, 5}, {3, 6}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := Transpose(test.M)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestDet(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected float64
		}{
			{
				Name:     "When calculating determinant of 1x1 matrix then return determinant",
				M:        types.M[float64]{{1}},
				Expected: 1,
			},
			{
				Name:     "When calculating determinant of 2x2 matrix then return determinant",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Expected: -2,
			},
			{
				Name:     "When calculating determinant of 3x3 matrix then return determinant",
				M:        types.M[float64]{{4, 2, 2}, {2, 1, 4}, {2, 2, 0}},
				Expected: -12,
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := Det(test.M)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestMinor(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Row      int
			Col      int
			Expected types.M[float64]
		}{
			{
				Name:     "When calculating minor of 2x2 matrix then return minor",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Row:      0,
				Col:      0,
				Expected: types.M[float64]{{4}},
			},
			{
				Name:     "When calculating minor [0,0] of 3x3 matrix then return minor",
				M:        types.M[float64]{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				Row:      0,
				Col:      0,
				Expected: types.M[float64]{{5, 6}, {8, 9}},
			},
			{
				Name:     "When calculating minor [1,1] of 3x3 matrix then return minor",
				M:        types.M[float64]{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				Row:      1,
				Col:      1,
				Expected: types.M[float64]{{1, 3}, {7, 9}},
			},
			{
				Name:     "When calculating minor [1,2] of 3x3 matrix then return minor",
				M:        types.M[float64]{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				Row:      1,
				Col:      2,
				Expected: types.M[float64]{{1, 2}, {7, 8}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := Minor(test.M, test.Row, test.Col)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestCofactor(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When calculating cofactor of 2x2 matrix then return cofactor",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Expected: types.M[float64]{{4, -3}, {-2, 1}},
			},
			{
				Name:     "When calculating cofactor of 3x3 matrix then return cofactor",
				M:        types.M[float64]{{4, 2, 2}, {2, 1, 4}, {2, 2, 0}},
				Expected: types.M[float64]{{-8, 8, 2}, {4, -4, -4}, {6, -12, 0}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := Cofactor(test.M)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestInverse(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected types.M[float64]
			Exists   bool
		}{
			{
				Name:   "When inverting 2x2 matrix then return inverse",
				M:      types.M[float64]{{1, 2}, {3, 4}},
				Exists: true,
			},
			{
				Name:   "When inverting 3x3 matrix then return inverse",
				M:      types.M[float64]{{4, 2, 2}, {2, 1, 4}, {2, 2, 0}},
				Exists: true,
			},
			{
				Name:   "When inverting 2x2 matrix with zero determinant then return no inverse",
				M:      types.M[float64]{{1, 2}, {2, 4}},
				Exists: false,
			},
			{
				Name:   "When inverting 3x3 matrix with zero determinant then return no inverse",
				M:      types.M[float64]{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				Exists: false,
			},
			{
				Name:   "When inverting 4x4 matrix with zero determinant then return no inverse",
				M:      types.M[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
				Exists: false,
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual, exists := Inverse(test.M)
			assert.Equal(t, test.Exists, exists)
			if !exists {
				assert.Nil(t, actual)
				return
			}
			assert.Equal(t, Product(test.M, actual), Identity[float64](len(test.M)))
			assert.Equal(t, Product(actual, test.M), Identity[float64](len(test.M)))
		})
	}
}

func BenchmarkForEachColumnConcurrent(b *testing.B) {
	var m = Zeros[float64](1000)
	for range b.N {
		m.Apply(concurrent.ForEachColumn(func(j int, c []float64) []float64 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := range c {
				c[i] = r.Float64()
			}
			return c
		}))
	}
}

func BenchmarkForEachColumn(b *testing.B) {
	var m = Zeros[float64](1000)
	for range b.N {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		m.Apply(ForEachColumn(func(j int, c []float64) []float64 {
			for i := range c {
				c[i] = r.Float64()
			}
			return c
		}))
	}
}

func BenchmarkForEach(b *testing.B) {
	var m = Zeros[float64](1000)
	for range b.N {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		m.Apply(ForEach(func(i, j int, t float64) float64 {
			return r.Float64()
		}))
	}
}

func BenchmarkLoop(b *testing.B) {
	var m = Zeros[float64](1000)
	for range b.N {
		var wg sync.WaitGroup
		for i := range m {
			wg.Add(1)
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			go func() {
				for j := range m[i] {
					m[i][j] = r.Float64()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

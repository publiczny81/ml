package matrix

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGaussElimination(t *testing.T) {
	var (
		tests = []struct {
			Name                string
			M                   types.M[float64]
			Expected            types.M[float64]
			ExpectedPermutation types.M[float64]
		}{
			{
				Name:                "When matrix is 2x2 then result is 2x2 with elements in upper triangular form",
				M:                   types.M[float64]{{1, 1}, {3, -2}},
				Expected:            types.M[float64]{{3, -2}, {0, 1.6666666666666665}},
				ExpectedPermutation: types.M[float64]{{0, 1}, {1, 0}},
			},
			{
				Name:                "When matrix is 3x3 then result is 3x3 with elements in upper triangular form",
				M:                   types.M[float64]{{2, 1, 4}, {2, 2, 0}, {4, 2, 2}},
				Expected:            types.M[float64]{{4, 2, 2}, {0, 1, -1}, {0, 0, 3}},
				ExpectedPermutation: types.M[float64]{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual, permutation := GaussElimination(test.M)
			assert.Equal(t, test.Expected, actual)
			assert.Equal(t, test.ExpectedPermutation, permutation)
		})
	}
}

func TestUpperTriangular(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected types.M[float64]
			Odd      bool
		}{
			{
				Name:     "When matrix is 2x2 then result is 2x2 with elements in upper triangular form",
				M:        types.M[float64]{{1, 1}, {3, -2}},
				Expected: types.M[float64]{{3, -2}, {0, 1.6666666666666665}},
				Odd:      true,
			},
			{
				Name:     "When matrix is 3x3 then result is 3x3 with elements in upper triangular form",
				M:        types.M[float64]{{2, 1, 4}, {2, 2, 0}, {4, 2, 2}},
				Expected: types.M[float64]{{4, 2, 2}, {0, 1, -1}, {0, 0, 3}},
				Odd:      true,
			},
			{
				Name:     "Test 3",
				M:        types.M[float64]{{4, 0, 1}, {1, 2, 1}, {2, 1, 3}},
				Expected: types.M[float64]{{4, 0, 1}, {0, 2, 0.75}, {0, 0, 2.5 - 0.5*0.75}},
				Odd:      false,
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual, odd := UpperTriangular(test.M)
			assert.Equal(t, test.Expected, actual)
			assert.Equal(t, test.Odd, odd)
		})
	}
}

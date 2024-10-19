package constructors

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MSuite struct {
	suite.Suite
}

func TestM(t *testing.T) {
	suite.Run(t, new(MSuite))
}

func (s *MSuite) TestCopyOf() {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected types.M[float64]
		}{
			{
				Name:     "When copying 2x2 matrix then return copy of matrix",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Expected: types.M[float64]{{1, 2}, {3, 4}},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			actual := CopyOf(test.M)
			s.Equal(test.Expected, actual)
		})
	}
}

func (s *MSuite) TestZeros() {
	var (
		tests = []struct {
			Name     string
			Args     []int
			Expected [][]float64
		}{
			{
				Name:     "When called with single argument 2 then new matrix is 2x2 with all zeros",
				Args:     []int{2},
				Expected: [][]float64{{0, 0}, {0, 0}},
			},
			{
				Name:     "When called with 2x3 then new matrix is 2x3 with all zeros",
				Args:     []int{2, 3},
				Expected: [][]float64{{0, 0, 0}, {0, 0, 0}},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			actual := Zeros[float64](test.Args...)
			s.Equal(test.Expected, actual)
		})
	}
}

func (s *MSuite) TestIdentity() {
	var (
		tests = []struct {
			Name     string
			Size     int
			Expected [][]float64
		}{
			{
				Name:     "When called with 2 then new matrix is 2x2 identity matrix",
				Size:     2,
				Expected: [][]float64{{1, 0}, {0, 1}},
			},
			{
				Name:     "When called with 3 then new matrix is 3x3 identity matrix",
				Size:     3,
				Expected: [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			actual := Identity[float64](test.Size)
			s.Equal(test.Expected, actual)
		})
	}
}

func TestShape(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			M        types.M[float64]
			Expected [2]int
		}{
			{
				Name:     "When matrix is 2x2 then return 2x2",
				M:        types.M[float64]{{1, 2}, {3, 4}},
				Expected: [2]int{2, 2},
			},
			{
				Name:     "When matrix is 2x3 then return 2x3",
				M:        types.M[float64]{{1, 2, 3}, {4, 5, 6}},
				Expected: [2]int{2, 3},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rows, cols := test.M.Shape()
			assert.Equal(t, test.Expected, [2]int{rows, cols})
		})
	}
}

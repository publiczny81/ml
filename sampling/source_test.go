package sampling

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SliceSourceSuite struct {
	suite.Suite
}

func TestSlice(t *testing.T) {
	suite.Run(t, new(SliceSourceSuite))
}

func (s *SliceSourceSuite) TestCount() {
	var tests = []struct {
		Name     string
		Source   *SliceSource[[]float64, float64]
		Expected int
	}{
		{
			Name:     "When slice is nil then return 0",
			Source:   NewSliceSource[[]float64, float64](nil),
			Expected: 0,
		},
		{
			Name:     "When slice has 3 elements then return 3",
			Source:   NewSliceSource([]float64{1, 2, 3}),
			Expected: 3,
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			s.Equal(test.Expected, test.Source.Count())
		})
	}
}

func (s *SliceSourceSuite) TestSelect() {
	var tests = []struct {
		Name     string
		Source   *SliceSource[[][]float64, []float64]
		Index    int
		Expected []float64
	}{
		{
			Name:     "When slice is nil then return nil",
			Source:   NewSliceSource[[][]float64, []float64](nil),
			Index:    0,
			Expected: nil,
		},
		{
			Name:     "When slice has 3 elements and index is within range then return valid element",
			Source:   NewSliceSource([][]float64{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}),
			Index:    1,
			Expected: []float64{2, 2, 2},
		},
		{
			Name:     "When slice has 3 elements and index is negative then return nil",
			Source:   NewSliceSource([][]float64{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}),
			Index:    -1,
			Expected: nil,
		},
		{
			Name:     "When slice has 3 elements and index is above range then return nil",
			Source:   NewSliceSource([][]float64{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}),
			Index:    3,
			Expected: nil,
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			s.Equal(test.Expected, test.Source.Select(test.Index))
		})
	}
}

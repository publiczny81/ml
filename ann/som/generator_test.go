package som

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearGenerator(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Limit    int
			Expected []Point
		}{
			{
				Name:     "When limit is negative then return no elements",
				Limit:    -1,
				Expected: nil,
			},
			{
				Name:     "When limit 0 then return no elements",
				Limit:    0,
				Expected: nil,
			},
			{
				Name:     "When limit is positive then return number of elements equals to limit",
				Limit:    3,
				Expected: []Point{{0}, {1}, {2}},
			},
		}
	)
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				actual []Point
			)
			for point := range NewLinearGenerator(test.Limit) {
				actual = append(actual, point)
			}
			assert.ElementsMatch(t, test.Expected, actual)
		})
	}
}

func TestRectangularGenerator(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Shape    []int
			Expected []Point
		}{
			{
				Name:  "When shape is negative then return no elements",
				Shape: []int{-1},
			},
			{
				Name:  "When shape is 0 then return no elements",
				Shape: []int{0, 0},
			},
			{
				Name:     "When shape is 1 then return 1x1 grid",
				Shape:    []int{1},
				Expected: []Point{{0, 0}},
			},
			{
				Name:     "When shape is 2X3 then return points representing 2x3 grid",
				Shape:    []int{2, 3},
				Expected: []Point{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {1, 1}, {2, 1}},
			},
		}
	)
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				actual []Point
			)
			for point := range NewRectangularGenerator(test.Shape...) {
				actual = append(actual, point)
			}
			assert.ElementsMatch(t, test.Expected, actual)
		})
	}
}

func TestNewHexagonalGenerator(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Shape    []int
			Expected []Point
		}{
			{
				Name:  "When shape is negative then return no elements",
				Shape: []int{-1},
			},
			{
				Name:  "When shape is 0 then return no elements",
				Shape: []int{0, 0},
			},
			{
				Name:     "When shape is 1 then return 1x1 grid",
				Shape:    []int{1},
				Expected: []Point{{0.5, 0}},
			},
			{
				Name:  "When shape is 2X3 then return points representing 2x3 grid",
				Shape: []int{2, 3},
				Expected: []Point{{0.5, 0}, {2, 1.7320508075688772 / 2}, {3.5, 0},
					{0.5, 1.7320508075688772}, {2, 1.5 * 1.7320508075688772}, {3.5, 1.7320508075688772}},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				actual []Point
			)
			for point := range NewHexagonalGenerator(test.Shape...) {
				actual = append(actual, point)
			}
			assert.ElementsMatch(t, test.Expected, actual)
		})
	}
}

//
//func TestHexagonalGenerator(t *testing.T) {
//	var (
//		generator = NewHexagonalGenerator(1)
//		tests     = []struct {
//			Name     string
//			P1       Point
//			P2       Point
//			Expected []Point
//		}{
//			{
//				Name:     "When p2 is nil then return 3 neighbours",
//				P1:       []float64{0, 0},
//				P2:       nil,
//				Expected: []Point{{1, 0}, {-0.5, 0.87}, {-0.5, -0.87}},
//			},
//			{
//				Name:     "When p1 and p2 is given then return 3 neighbours",
//				P1:       []float64{0, 0},
//				P2:       []float64{1, 0},
//				Expected: []Point{{-0.5, 0.87}, {-0.5, -0.87}},
//			},
//			{
//				Name:     "When another p1 and p2 is given then return 2 neighbours",
//				P1:       []float64{1, 0},
//				P2:       []float64{0, 0},
//				Expected: []Point{{1.5, -0.87}, {1.5, 0.87}},
//			},
//		}
//	)
//	for _, test := range tests {
//		t.Run(test.Name, func(t *testing.T) {
//			var (
//				actual []Point
//			)
//			for point := range generator() {
//				point[0] = utils.Round(point[0], 2)
//				point[1] = utils.Round(point[1], 2)
//				actual = append(actual, point)
//			}
//			assert.ElementsMatch(t, test.Expected, actual)
//		})
//	}
//}

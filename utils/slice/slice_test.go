package slice

import (
	"github.com/publiczny81/ml/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterate(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Actual   []*int
			Function func(*int) bool
			Expected []*int
		}{
			{
				Name:   "When function called on each element then each element updated",
				Actual: []*int{utils.Pointer(0), utils.Pointer(0), utils.Pointer(0)},
				Function: func(i *int) bool {
					*i = 1
					return true
				},
				Expected: []*int{utils.Pointer(1), utils.Pointer(1), utils.Pointer(1)},
			},
			{
				Name:   "When function returns false then it stops iterating",
				Actual: []*int{utils.Pointer(3), utils.Pointer(1), utils.Pointer(2)},
				Function: func(i *int) bool {
					if *i == 1 {
						return false
					}
					*i = 0
					return true
				},
				Expected: []*int{utils.Pointer(0), utils.Pointer(1), utils.Pointer(2)},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			Iterate(test.Actual, test.Function)
			assert.Equal(t, test.Expected, test.Actual)
		})
	}
}

func TestIterateWithIndex(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Actual   []*int
			Function func(int, *int) bool
			Expected []*int
		}{
			{
				Name:   "When function called on each element then each element updated",
				Actual: []*int{utils.Pointer(0), utils.Pointer(0), utils.Pointer(0)},
				Function: func(idx int, i *int) bool {
					if idx > 0 {
						*i = idx
					}
					return true
				},
				Expected: []*int{utils.Pointer(0), utils.Pointer(1), utils.Pointer(2)},
			},
			{
				Name:   "When function returns false then it stops iterating",
				Actual: []*int{utils.Pointer(0), utils.Pointer(0), utils.Pointer(0)},
				Function: func(idx int, i *int) bool {
					if idx > 1 {
						return false
					}
					*i = idx
					return true
				},
				Expected: []*int{utils.Pointer(0), utils.Pointer(1), utils.Pointer(0)},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			IterateWithIndex(test.Actual, test.Function)
			assert.Equal(t, test.Expected, test.Actual)
		})
	}
}

func TestAggregate(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Slice    []int
			Initial  int
			Expected int
		}{
			{
				Name:     "Given sum function when called with empty slice then return initial value",
				Initial:  3,
				Expected: 3,
			},
			{
				Name:     "Given sum function when called with 3-element slice then return sum of the elements",
				Slice:    []int{1, 2, 3},
				Initial:  0,
				Expected: 6,
			},
			{
				Name:     "Given sum function when called with 3-element slice and initial value then return sum of all elements",
				Slice:    []int{1, 2, 3},
				Initial:  1,
				Expected: 7,
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = Aggregate(test.Slice, test.Initial, func(i int, i2 int) int {
				return i + i2
			})

			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestApply(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Actual   []*int
			Function func(*int) *int
			Expected []*int
		}{
			{
				Name:   "When function called on each element then each element updated",
				Actual: []*int{utils.Pointer(0), utils.Pointer(0), utils.Pointer(0)},
				Function: func(i *int) *int {
					*i = 1
					return i
				},
				Expected: []*int{utils.Pointer(1), utils.Pointer(1), utils.Pointer(1)},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			Apply(test.Actual, test.Function)
			assert.Equal(t, test.Expected, test.Actual)
		})
	}
}

func TestApplyWithIndex(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Actual   []*int
			Function func(int, *int) *int
			Expected []*int
		}{
			{
				Name:   "When function called on each element then each element updated",
				Actual: []*int{utils.Pointer(0), utils.Pointer(0), utils.Pointer(0)},
				Function: func(idx int, i *int) *int {
					if idx > 0 {
						*i = idx
					}
					return i
				},
				Expected: []*int{utils.Pointer(0), utils.Pointer(1), utils.Pointer(2)},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ApplyWithIndex(test.Actual, test.Function)
			assert.Equal(t, test.Expected, test.Actual)
		})
	}
}

package array

import (
	"github.com/publiczny81/ml/utils/slice"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type ArraySuite struct {
	suite.Suite
}

func TestArray(t *testing.T) {
	suite.Run(t, new(ArraySuite))
}

func (s *ArraySuite) TestNew() {
	var (
		tests = []struct {
			Name     string
			Factory  func() *Array[float64]
			Expected *Array[float64]
		}{
			{
				Name: "When initialized with no dimensions then matrix data is empty",
				Factory: func() *Array[float64] {
					return New[float64]()
				},
				Expected: &Array[float64]{},
			},
			{
				Name: "When initialized with one dimension then matrix data is vector",
				Factory: func() *Array[float64] {
					return New[float64](6)
				},
				Expected: &Array[float64]{
					dim:  []int{6},
					data: make([]float64, 6),
				},
			},
			{
				Name: "When initialized with two dimension then matrix data is matrix",
				Factory: func() *Array[float64] {
					return New[float64](6, 4)
				},
				Expected: &Array[float64]{
					dim:  []int{6, 4},
					data: make([]float64, 24),
				},
			},
			{
				Name: "When initialized with three dimension then matrix data is 3d matrix",
				Factory: func() *Array[float64] {
					return New[float64](6, 4, 2)
				},
				Expected: &Array[float64]{
					dim:  []int{6, 4, 2},
					data: make([]float64, 48),
				},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var actual = test.Factory()
			s.True(equalArray(test.Expected, actual))
			//s.Equal(test.Expected.dim, actual.dim)
			//s.Equal(test.Expected.data, actual.data)
		})
	}
}

func (s *ArraySuite) TestBuilder() {
	var (
		tests = []struct {
			Name     string
			Factory  func() *Array[int]
			Expected *Array[int]
		}{
			{
				Name: "When build only dimension then return uninitialized matrix",
				Factory: func() *Array[int] {
					return NewBuilder[int](6, 4, 2).Build()
				},
				Expected: &Array[int]{
					dim:  []int{6, 4, 2},
					data: make([]int, 48),
				},
			},
			{
				Name: "When build with initFunc then return matrix initialized with value from the initFunc",
				Factory: func() *Array[int] {
					return NewBuilder[int](3, 2, 2).
						WithInitFunc(func(idx int) int {
							return idx + 1
						}).Build()
				},
				Expected: &Array[int]{
					dim:  []int{3, 2, 2},
					data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				},
			},
			{
				Name: "When build with applyFunc then return matrix initialized with value from the applyFunc",
				Factory: func() *Array[int] {
					return NewBuilder[int](3, 2, 2).
						WithApplyFunc(func(idx int, value int) int {
							return value + 1
						}).Build()
				},
				Expected: &Array[int]{
					dim:  []int{3, 2, 2},
					data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				},
			},
			{
				Name: "When build with data then return matrix initialized with data",
				Factory: func() *Array[int] {
					return NewBuilder[int](3, 2, 2).
						WithData([]int{1, 2, 3, 4, 5, 6}).Build()
				},
				Expected: &Array[int]{
					dim:  []int{3, 2, 2},
					data: []int{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0},
				},
			},
		}
	)
	for _, test := range tests {
		s.Run(test.Name, func() {
			var actual = test.Factory()
			s.True(equalArray(test.Expected, actual))
		})
	}
}

func equalArray[T any](a *Array[T], a2 *Array[T]) (equal bool) {
	if a == a2 {
		equal = true
		return true
	}
	if a == nil {
		return false
	}
	if a2 == nil {
		return false
	}
	for idx, e := range a.dim {
		if a2.dim[idx] == e {
			continue
		}
		return false
	}
	return reflect.DeepEqual(a.data, a2.data)
}

func (s *ArraySuite) TestIndexPosition() {
	type subTest struct {
		Name          string
		Index         int
		Position      []int
		IndexError    error
		PositionError error
	}
	var (
		tests = []struct {
			Name     string
			Dim      []int
			SubTests []subTest
		}{
			{
				Name: "Given 3-dimension array with one element",
				Dim:  []int{1, 1, 1},
				SubTests: []subTest{
					{
						Index:    0,
						Position: []int{0, 0, 0},
					},
				},
			},
			{
				Name: "Given 3-dimension array",
				Dim:  []int{3, 2, 2},
				SubTests: []subTest{
					{
						Name:     "When query for first element then index corresponds to position",
						Index:    0,
						Position: []int{0, 0, 0},
					},
					{
						Name:     "When query for second element then index corresponds to position",
						Index:    1,
						Position: []int{0, 0, 1},
					},
					{
						Name:     "When query for third element then index corresponds to position",
						Index:    2,
						Position: []int{0, 1, 0},
					},
					{
						Name:     "When query for forth element then index corresponds to position",
						Index:    3,
						Position: []int{0, 1, 1},
					},
					{
						Name:     "When query for fifth element then index corresponds to position",
						Index:    4,
						Position: []int{1, 0, 0},
					},
					{
						Name:     "When query for sixth element then index corresponds to position",
						Index:    5,
						Position: []int{1, 0, 1},
					},
					{
						Name:     "When query for seventh element then index corresponds to position",
						Index:    6,
						Position: []int{1, 1, 0},
					},
					{
						Name:     "When query for eighth element then index corresponds to position",
						Index:    7,
						Position: []int{1, 1, 1},
					},
					{
						Name:     "When query for ninth element then index corresponds to position",
						Index:    8,
						Position: []int{2, 0, 0},
					},
					{
						Name:       "When passed invalid position dimension then panic with error",
						Index:      11,
						Position:   []int{3, 2, 2, 1},
						IndexError: InvalidDimensionError,
					},
					{
						Name:       "When passed invalid position then return panic with error",
						Index:      11,
						Position:   []int{2, 1, 2},
						IndexError: IndexOutOfRangeError,
					},
					{
						Name:          "When passed invalid index then panic with error",
						Index:         12,
						Position:      []int{2, 1, 1},
						PositionError: IndexOutOfRangeError,
					},
				},
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var index, position = MakeIndexPositionFunc(test.Dim...)
			slice.Iterate(test.SubTests, func(test subTest) bool {
				s.Run(test.Name, func() {
					if test.IndexError != nil {
						s.PanicsWithError(test.IndexError.Error(), func() {
							_ = index(test.Position...)
						})
						return
					}
					if test.PositionError != nil {
						s.PanicsWithError(test.PositionError.Error(), func() {
							_ = position(test.Index)
						})
						return
					}
					var (
						actualIndex    = index(test.Position...)
						actualPosition = position(test.Index)
					)
					s.Equal(test.Index, actualIndex)
					s.Equal(test.Position, actualPosition)
				})
				return true
			})
		})
	}
}

func (s *ArraySuite) TestSize() {
	var (
		tests = []struct {
			Name     string
			Array    *Array[int]
			Expected int
		}{
			{
				Name:     "Given empty array then return zero",
				Array:    New[int](0, 0),
				Expected: 0,
			},
			{
				Name:     "Given one-dimension array then return its length",
				Array:    New[int](2),
				Expected: 2,
			},
			{
				Name:     "Given two-dimension array then return number of elements",
				Array:    New[int](3, 2),
				Expected: 6,
			},
		}
	)

	for _, test := range tests {
		s.Run(test.Name, func() {
			var actual = test.Array.Size()
			s.Equal(test.Expected, actual)
		})
	}
}

func (s *ArraySuite) TestBackedData() {
	var (
		actual = NewBuilder[int](5).WithInitFunc(func(idx int) int {
			return idx
		}).Build().BackedData()
		expected = []int{0, 1, 2, 3, 4}
	)

	s.Equal(expected, actual)
}

func (s *ArraySuite) TestDim() {
	var (
		actual   = New[int](5, 4, 3, 2, 1).Dim()
		expected = []int{5, 4, 3, 2, 1}
	)

	s.Equal(expected, actual)
}

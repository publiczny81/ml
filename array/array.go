package array

import (
	"github.com/pkg/errors"
	"github.com/publiczny81/ml/utils"
	"github.com/publiczny81/ml/utils/slices"
)

type Array[T any] struct {
	dim      []int
	data     []T
	Index    func(idx ...int) int
	Position func(int) []int
}

func MakeIndexPositionFunc(dim ...int) (index func(...int) int, position func(int) []int) {
	var (
		shift  = make([]int, len(dim))
		length = 1
	)

	slices.IterateWithIndex(dim, func(idx, d int) bool {
		length *= d
		shift[idx] = slices.Aggregate(dim[idx+1:], 1, func(i int, i2 int) int {
			return i * i2
		})
		return true
	})

	index = func(pos ...int) (index int) {
		if len(pos) != len(dim) {
			panic(errors.WithStack(InvalidDimensionError))
		}
		for i, d := range pos {
			if d >= dim[i] {
				panic(errors.WithStack(IndexOutOfRangeError))
			}
			index += d * shift[i]
		}
		return
	}

	position = func(idx int) (position []int) {
		if idx >= length {
			panic(errors.WithStack(IndexOutOfRangeError))
		}

		for _, v := range shift {
			position = append(position, idx/v)
			idx = idx % v
		}
		return
	}
	return
}

type InitFunc[T any] func(idx int) T

type ApplyFunc[T any] func(idx int, value T) T

func New[T any](dim ...int) (m *Array[T]) {
	if len(dim) == 0 {
		m = &Array[T]{}
		return
	}

	m = &Array[T]{
		dim: dim,
		data: make([]T, slices.Aggregate(dim, 1, func(i int, i2 int) int {
			return i * i2
		})),
	}
	m.Index, m.Position = MakeIndexPositionFunc(dim...)
	return
}

type Builder[T any] func() *Array[T]

func NewBuilder[T any](dim ...int) Builder[T] {
	return func() *Array[T] {
		return New[T](dim...)
	}
}

func (b Builder[T]) WithInitFunc(initFunc InitFunc[T]) Builder[T] {
	return func() (m *Array[T]) {
		m = b()
		slices.IterateWithIndex(m.data, func(i int, t T) bool {
			m.data[i] = initFunc(i)
			return true
		})
		return
	}
}

func (b Builder[T]) WithApplyFunc(f ApplyFunc[T]) Builder[T] {
	return func() (m *Array[T]) {
		m = b()
		slices.ApplyWithIndex(m.data, f)
		return
	}
}

func (b Builder[T]) WithData(data []T) Builder[T] {
	return func() (m *Array[T]) {
		m = b()
		var limit = min(len(data), len(m.data))
		copy(m.data[:limit], data[:limit])
		return
	}
}

func (b Builder[T]) Build() *Array[T] {
	return b()
}

func (a *Array[T]) Dim() []int {
	return a.dim
}

func (a *Array[T]) BackedData() []T {
	return a.data
}

func (a *Array[T]) Size() (s int) {
	if len(a.dim) == 0 {
		return
	}
	seq := utils.NewSeq(a.dim)
	return seq.Aggregate(1, func(i int, i2 int) int {
		return i * i2
	})
}

func (a *Array[T]) Get(index ...int) (v T) {
	return a.data[a.Index(index...)]
}

func (a *Array[T]) Set(value T, index ...int) {
	a.data[a.Index(index...)] = value
}

func (a *Array[T]) Iterate(f func(v T) bool) {
	slices.Iterate(a.data, f)
}

func (a *Array[T]) IterateWithIndex(f func(idx int, v T) bool) {
	slices.IterateWithIndex(a.data, f)
}

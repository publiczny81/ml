package sampling

import (
	"context"
	"github.com/publiczny81/ml/errors"
	"reflect"
)

type SliceSource[S ~[]E, E any] struct {
	slice S
}

func NewSliceSource[S ~[]E, E any](s S) *SliceSource[S, E] {
	return &SliceSource[S, E]{
		slice: s,
	}
}

func (s *SliceSource[S, E]) Count(context.Context) (int, error) {
	return len(s.slice), nil
}

func (s *SliceSource[S, E]) Select(_ context.Context, idx int) (e E, err error) {
	if s.slice == nil {
		return
	}
	if idx < 0 {
		return
	}
	if idx >= len(s.slice) {
		return
	}
	e = s.slice[idx]
	return
}

type LimitedSource[E any] struct {
	source Source[E]
	from   int
	to     int
}

func NewLimitedSource[E any](source Source[E], from, to int) (s *LimitedSource[E], err error) {
	if source == nil || reflect.ValueOf(source).IsZero() {
		err = errors.WithMessage(errors.InvalidParameterError, "NewLimitedSource: source is nil")
		return
	}
	if from < 0 {
		err = errors.WithMessage(errors.InvalidParameterError, "NewLimitedSource: from is negative")
		return
	}
	if to < 0 {
		err = errors.WithMessage(errors.InvalidParameterError, "NewLimitedSource: to is negative")
		return
	}
	if to < from {
		err = errors.WithMessage(errors.InvalidParameterError, "NewLimitedSource: to is less than from")
		return
	}
	s = &LimitedSource[E]{
		source: source,
		from:   max(0, from),
		to:     max(0, to),
	}
	return
}

func MustNewLimitedSource[E any](source Source[E], from, to int) *LimitedSource[E] {
	s, err := NewLimitedSource[E](source, from, to)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *LimitedSource[E]) Count(context.Context) (int, error) {
	return s.to - s.from, nil
}

func (s *LimitedSource[E]) Select(ctx context.Context, idx int) (E, error) {
	return s.source.Select(ctx, s.from+idx)
}

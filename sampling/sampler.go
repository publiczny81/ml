package sampling

import (
	"context"
)

// Source defines contract for samples provider
type Source[E any] interface {
	Count(context.Context) (int, error)
	Select(context.Context, int) (E, error)
}

type Sample[E any] struct {
	Value E
	Error error
}

func ValueOf[E any](e E) Sample[E] {
	return Sample[E]{
		Value: e,
	}
}

func Error[E any](e error) Sample[E] {
	return Sample[E]{
		Error: e,
	}
}

// Strategy defines contract for selection samples from Source
type Strategy[E any] interface {
	// Samples returns channel with samples from Source
	Samples(context.Context, Source[E]) <-chan Sample[E]
}

// Sampler iterates over samples from the Source according to given Strategy
type Sampler[E any] struct {
	source   Source[E]
	strategy Strategy[E]
}

// New creates new sampler with source and sampling strategy
func New[E any](source Source[E], strategy Strategy[E]) (s *Sampler[E]) {
	s = &Sampler[E]{
		source:   source,
		strategy: strategy,
	}
	return
}

func (s *Sampler[E]) Samples(ctx context.Context) <-chan Sample[E] {
	return s.strategy.Samples(ctx, s.source)
}

package sampling

import "context"

// SystematicalStrategy implements Strategy interface. It generates consecutive integers in range from zero until limit of source is reached.
type SystematicalStrategy[E any] struct{}

func (s *SystematicalStrategy[E]) Samples(ctx context.Context, source Source[E]) <-chan Sample[E] {
	var (
		ch = make(chan Sample[E])
	)
	go func() {
		defer close(ch)
		var (
			current    = 0
			limit, err = source.Count(ctx)
			e          E
		)
		if err != nil {
			ch <- Error[E](err)
			return
		}
		for current < limit {
			if e, err = source.Select(ctx, current); err != nil {
				ch <- Error[E](err)
				return
			}
			current++
			select {
			case <-ctx.Done():
				ch <- Error[E](ctx.Err())
				return
			case ch <- ValueOf(e):
			}
		}
	}()
	return ch
}

// Rand defines contract for random number generator
type Rand interface {
	// IntN generates an integer number in range <0, n)
	IntN(n int) int
}

type RandomStrategy[E any] struct {
	rand Rand
}

func NewRandomStrategy[E any](rand Rand) *RandomStrategy[E] {
	return &RandomStrategy[E]{
		rand: rand,
	}
}

func (s *RandomStrategy[E]) Samples(ctx context.Context, source Source[E]) <-chan Sample[E] {
	var (
		ch = make(chan Sample[E])
	)
	go func() {
		defer close(ch)

		var (
			limit, err = source.Count(ctx)
			e          E
		)
		if err != nil {
			ch <- Error[E](err)
			return
		}
		if e, err = source.Select(ctx, s.rand.IntN(limit)); err != nil {
			ch <- Error[E](err)
			return
		}

		select {
		case <-ctx.Done():
			ch <- Error[E](ctx.Err())
			return
		case ch <- ValueOf(e):
		}

	}()
	return ch
}

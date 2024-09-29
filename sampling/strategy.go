package sampling

// SystematicalStrategy implements Strategy interface. It generates consecutive integers in range from zero until limit
type SystematicalStrategy[E any] struct {
	limit   int
	current int
}

// NewSystematicalStrategyWithLimit creates new SystematicalStrategy with given limit
func NewSystematicalStrategyWithLimit[E any](limit int) *SystematicalStrategy[E] {
	return &SystematicalStrategy[E]{
		limit: limit,
	}
}

// NewSystematicalStrategyFromSource creates new SystematicalStrategy from Source
func NewSystematicalStrategyFromSource[E any](source Source[E]) *SystematicalStrategy[E] {
	return &SystematicalStrategy[E]{
		limit: source.Count(),
	}
}

func (s *SystematicalStrategy[E]) Next(source Source[E]) (e E, found bool) {
	if s.current < s.limit {
		e = source.Select(s.current)
		found = true
		s.current++
	}
	return
}

// Rand defines contract for random number generator
type Rand interface {
	// IntN generates an integer number in range <0, n)
	IntN(n int) int
}

type RandomStrategy[E any] struct {
	rand     Rand
	executed bool
}

func NewRandomStrategy[E any](rand Rand) *RandomStrategy[E] {
	return &RandomStrategy[E]{
		rand: rand,
	}
}

func (s *RandomStrategy[E]) Next(source Source[E]) (e E, found bool) {
	if !s.executed {
		e = source.Select(s.rand.IntN(source.Count()))
		found = true
		s.executed = true
		return
	}
	return
}

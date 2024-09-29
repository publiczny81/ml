package sampling

// Source defines contract for samples provider
type Source[E any] interface {
	Count() int
	Select(int) E
}

// Strategy defines contract for selection samples from Source
type Strategy[E any] interface {
	// Next returns sample from Source and true if successful. Otherwise, it returns false
	Next(Source[E]) (E, bool)
}

// StrategyFactory creates new Strategy
type StrategyFactory[E any] func() Strategy[E]

// Sampler iterates over samples from the Source according to given Strategy
type Sampler[E any] struct {
	hasNext     bool
	next        E
	source      Source[E]
	newStrategy StrategyFactory[E]
	strategy    Strategy[E]
}

// New creates new sampler with source and given StrategyFactory
func New[E any](source Source[E], factory StrategyFactory[E]) (s *Sampler[E]) {
	s = &Sampler[E]{
		source:      source,
		newStrategy: factory,
	}
	return
}

// Next returns next sample
func (s *Sampler[E]) Next() (next E) {
	next = s.next
	s.next, s.hasNext = s.strategy.Next(s.source)
	return
}

// HasNext returns true as long as there are available more samples
func (s *Sampler[E]) HasNext() bool {
	return s.hasNext
}

// Reset resets Sampler allowing for new iteration over samples
func (s *Sampler[E]) Reset() {
	s.strategy = s.newStrategy()
	s.next, s.hasNext = s.strategy.Next(s.source)
}

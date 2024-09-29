package utils

type Seq[S ~[]E, E any] []E

func NewSeq[S ~[]E, E any](s S) Seq[S, E] {
	return Seq[S, E](s)
}

func (s Seq[S, E]) Iterate(f func(E) bool) {
	for _, e := range s {
		if f(e) {
			continue
		}
		return
	}
}

func (s Seq[S, E]) IterateWithIndex(f func(int, E) bool) {
	for idx, e := range s {
		if f(idx, e) {
			continue
		}
		return
	}
}

func (s Seq[S, E]) Apply(f func(E) E) {
	for idx, e := range s {
		s[idx] = f(e)
	}
}

func (s Seq[S, E]) ApplyWithIndex(f func(int, E) E) {
	for idx, e := range s {
		s[idx] = f(idx, e)
	}
}

func (s Seq[S, E]) Aggregate(initial E, f func(E, E) E) (value E) {
	value = initial
	for _, e := range s {
		value = f(value, e)
	}
	return
}

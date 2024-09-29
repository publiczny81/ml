package sampling

type SliceSource[S ~[]E, E any] struct {
	slice S
}

func NewSliceSource[S ~[]E, E any](s S) *SliceSource[S, E] {
	return &SliceSource[S, E]{
		slice: s,
	}
}

func (s *SliceSource[S, E]) Count() int {
	return len(s.slice)
}

func (s *SliceSource[S, E]) Select(idx int) (e E) {
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

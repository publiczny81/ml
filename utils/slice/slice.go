package slice

func Iterate[S ~[]E, E any](s S, f func(E) bool) {
	for i := range s {
		if !f(s[i]) {
			return
		}
	}
}

func IterateWithIndex[S ~[]E, E any](s S, f func(int, E) bool) {
	for i, e := range s {
		if !f(i, e) {
			return
		}
	}
}

func Aggregate[S ~[]E, E any](s S, initial E, f func(E, E) E) (value E) {
	value = initial
	for _, e := range s {
		value = f(value, e)
	}
	return
}

func Apply[S ~[]E, E any](s S, f func(E) E) {
	for idx, e := range s {
		s[idx] = f(e)
	}
}

func ApplyWithIndex[S ~[]E, E any](s S, f func(int, E) E) {
	for idx, e := range s {
		s[idx] = f(idx, e)
	}
}

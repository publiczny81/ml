package slices

func Initiate[S ~[]E, E any](initial S, initFunc func(idx int, e E) E) S {
	ApplyWithIndex(initial, initFunc)
	return initial
}

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

func Reverse[S ~[]E, E any](s S, f func(E) bool) {
	for i := len(s) - 1; i > -1; i-- {
		if !f(s[i]) {
			return
		}
	}
}

func ReverseWithIndex[S ~[]E, E any](s S, f func(int, E) bool) {
	for i := len(s) - 1; i > -1; i-- {
		if !f(i, s[i]) {
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

func AggregateWithIndex[S ~[]E, E any](s S, initial E, f func(E, int, E) E) (value E) {
	value = initial
	for i, e := range s {
		value = f(value, i, e)
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

func GroupBy[S ~[]E, E any, K comparable](s S, f func(E) K) map[K]S {
	var result = make(map[K]S)
	for _, e := range s {
		key := f(e)
		result[key] = append(result[key], e)
	}
	return result
}

func Filter[S ~[]E, E any](s S, f func(E) bool) (result S) {
	for _, e := range s {
		if f(e) {
			result = append(result, e)
		}
	}
	return
}

func Flatten[S ~[]R, R ~[]E, E any](s S) (result []E) {
	for _, e := range s {
		result = append(result, e...)
	}
	return
}

func FlattenMap[S ~[]R, R map[K]V, K comparable, V any](s S) (result map[K]V) {
	result = make(map[K]V)
	for _, e := range s {
		for k, v := range e {
			result[k] = v
		}
	}
	return
}

func Resize[S ~[]T, T any](v S, size int) (r S) {
	size = max(0, size)
	if len(v) >= size {
		return v[:size]
	}
	if size <= cap(v) {
		old := len(v)
		r = v[:size]
		clear(r[old:])
		return
	}
	r = make(S, size)
	copy(r, v)
	return
}

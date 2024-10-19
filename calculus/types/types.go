package types

type MetricsFunc[S ~[]T, T Number] func(S, S) T

type Float interface {
	float32 | float64
}

type Int interface {
	int8 | int16 | int32 | int64 | int
}

type Real interface {
	Int | Float
}

type Complex interface {
	complex64 | complex128
}

type Number interface {
	Real | Complex
}

type V[T Real] []T

func (v V[T]) Apply(operations ...VOperation[T]) {
	for _, o := range operations {
		o(v)
	}
}

func (v V[T]) Size() int {
	return len(v)
}

type VOperation[T Real] func(V[T])

type M[T Real] [][]T

type MOperation[T Real] func(M[T])

func (m M[T]) Apply(operation ...MOperation[T]) {
	for _, o := range operation {
		if o != nil {
			o(m)
		}
	}
	return
}

func (m M[T]) Shape() (int, int) {
	return len(m), len((m)[0])
}

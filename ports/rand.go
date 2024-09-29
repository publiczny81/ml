package types

type Rand interface {
	Int() int
	IntN(n int) int
	Uint64() uint64
	Float64() float64
}

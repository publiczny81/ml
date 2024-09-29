package utils

type RandSource struct{}

func (s RandSource) Uint64() uint64 {
	return uint64(Now().UnixNano())
}

func NewRandSource() RandSource {
	return RandSource{}
}

func Randomize[S ~[]E, E any](s S, generator func() E) {
	for i := range s {
		s[i] = generator()
	}
}

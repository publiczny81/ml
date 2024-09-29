package utils

func Pointer[PT *T, T any](value T) PT {
	return &value
}

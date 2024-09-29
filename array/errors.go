package array

import "github.com/pkg/errors"

var (
	InvalidDimensionError = errors.New("invalid index dimension")
	IndexOutOfRangeError  = errors.New("index out of range")
)

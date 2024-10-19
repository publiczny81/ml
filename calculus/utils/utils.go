package utils

import (
	"github.com/publiczny81/ml/calculus/types"
	"math"
)

var (
	float32Error = float32(1e-14)
	float64Error = 1e-29
)

func Abs[T types.Real](x T) T {
	return max(x, -x)
}

func IsZero[T types.Real](val T) bool {
	switch value := any(val).(type) {
	case float32:
		return Abs(value) < float32Error
	case float64:
		return Abs(value) < float64Error
	default:
		return value == 0
	}
}

func Dim[T types.Real](x, y T) T {
	v := x - y
	if v <= 0 {
		return 0
	}
	return v
}

func Round[T types.Real](value T, precision int) T {
	ratio := math.Pow(10, float64(precision))
	return T(math.Round(float64(value)*ratio) / ratio)
}

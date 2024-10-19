package functions

import "math"

func Sigmoid(value float64) float64 {
	return 1 / (1 + math.Exp(value))
}

func DerivativeSigmoid(value float64) (ret float64) {
	ret = Sigmoid(value)
	return ret * (1 - ret)
}

func Rectifier(value float64) float64 {
	return math.Max(0, value)
}

func ParametricRectifier(a float64) func(float64) float64 {
	return func(value float64) float64 {
		return math.Max(a*value, value)
	}
}

func DerivativeParametricRectifier(a float64) func(float64) float64 {
	return func(value float64) float64 {
		if value > 0 {
			return 1
		}
		return a
	}
}

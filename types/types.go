package types

type MetricsFunc func([]float64, []float64) float64

type Float interface {
	float32 | float64
}

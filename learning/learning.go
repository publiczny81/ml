package learning

type ConstantRate float64

func (r ConstantRate) Rate(int) float64 {
	return float64(r)
}

package learning

type Scheduler func(int) float64

func (f Scheduler) LearningRate(epoch int) float64 {
	return f(epoch)
}

type ConstantRate float64

func (r ConstantRate) LearningRate(int) float64 {
	return float64(r)
}

func LinearRateSchedule(epochs int) Scheduler {
	return func(epoch int) float64 {
		return 1.0 - float64(epoch)/float64(epochs)
	}
}

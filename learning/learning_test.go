package learning

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstantRate(t *testing.T) {
	var tests = []struct {
		Name     string
		Rate     ConstantRate
		Expected float64
	}{
		{
			Name:     "When ConstantRate is 0.6 then return 0.6",
			Rate:     ConstantRate(0.6),
			Expected: 0.6,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			for i := range 3 {
				assert.Equal(t, test.Expected, test.Rate.LearningRate(i))
			}
		})
	}
}

func TestLinearRate(t *testing.T) {
	var tests = []struct {
		Name     string
		Rate     Scheduler
		Epoch    int
		Expected float64
	}{
		{
			Name:     "When max epochs is 2 and current epoch is 1 then return 0.5",
			Rate:     LinearRateSchedule(2),
			Epoch:    1,
			Expected: 0.5,
		},
		{
			Name:     "When max epochs is 2 and current epoch is 2 then return 0",
			Rate:     LinearRateSchedule(2),
			Epoch:    2,
			Expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = test.Rate(test.Epoch)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

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
				assert.Equal(t, test.Expected, test.Rate.Rate(i))
			}
		})
	}
}

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Value    float64
			Expected float64
		}{
			{
				Name:     "When value is positive then return value",
				Value:    1,
				Expected: 1,
			},
			{
				Name:     "When value is negative then return positive value",
				Value:    -1,
				Expected: 1,
			},
			{
				Name:     "When value is zero then return zero",
				Value:    0,
				Expected: 0,
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = Abs(test.Value)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestIsZero(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Value    float64
			Expected bool
		}{
			{
				Name:     "When value is zero then return true",
				Value:    0,
				Expected: true,
			},
			{
				Name:     "When value is not zero then return false",
				Value:    1,
				Expected: false,
			},
			{
				Name:     "When value is close to zero then return true",
				Value:    1e-30,
				Expected: true,
			},
		}
	)
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = IsZero(test.Value)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestRound(t *testing.T) {
	var tests = []struct {
		Name      string
		Given     float64
		Precision int
		Expected  float64
	}{
		{
			Name:      "When precision is 2 then return rounded value",
			Given:     1.23456789,
			Precision: 2,
			Expected:  1.23,
		},
		{
			Name:      "When precision is 3 then return rounded value",
			Given:     1.23456789,
			Precision: 3,
			Expected:  1.235,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = Round(test.Given, test.Precision)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

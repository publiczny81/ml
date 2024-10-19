package losses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMeanSquareError(t *testing.T) {
	var tests = []struct {
		Name      string
		Actual    []float64
		Predicted []float64
		Partials  []float64
		Value     float64
		Error     string
	}{
		{
			Name:      "When actual and predicted are the same then value is 0",
			Actual:    []float64{1.0, 2.0, 3.0},
			Predicted: []float64{1.0, 2.0, 3.0},
			Partials:  []float64{0.0, 0.0, 0.0},
			Value:     0.0,
		},
		{
			Name:      "When actual and predicted are different then value is 0.6666666666666666",
			Actual:    []float64{1.0, 2.0, 3.0, 4.0},
			Predicted: []float64{2.0, 2.0, 2.0, 2.0},
			Partials:  []float64{-1.0, 0.0, 1.0, 2.0},
			Value:     1.5,
		},
		{
			Name:      "When actual and predicted are empty then value is 0",
			Actual:    []float64{},
			Predicted: []float64{},
			Partials:  nil,
			Value:     0.0,
		},
		{
			Name:      "When actual and predicted have different lengths then panic",
			Actual:    []float64{1.0, 2.0},
			Predicted: []float64{1.0},
			Error:     "unmatched size of vectors",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.Error != "" {
				assert.PanicsWithError(t, test.Error, func() {
					_, _ = MeanSquareError(test.Actual, test.Predicted)
				})
				return
			}
			partials, value := MeanSquareError(test.Actual, test.Predicted)
			assert.Equal(t, test.Value, value)
			assert.Equal(t, test.Partials, partials)
		})
	}
}

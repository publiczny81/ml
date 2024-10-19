package constructors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyOf(t *testing.T) {
	var tests = []struct {
		Name string
		S    []float64
	}{
		{
			Name: "When slice is nil then new vector is empty",
			S:    []float64{},
		},
		{
			Name: "When slice is empty then new vector is empty",
			S:    []float64{},
		},
		{
			Name: "When slice has 3 elements then new vector has 3 elements",
			S:    []float64{1.0, 2.0, 3.0},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var v = CopyOf[[]float64, float64](test.S)
			assert.Equal(t, test.S, v)
		})
	}
}

package activate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	var tests = []struct {
		Name     string
		Params   []any
		Expected string
		Found    bool
	}{
		{
			Name:     Sigmoid,
			Expected: Sigmoid,
			Found:    true,
		},
		{
			Name:     Rectifier,
			Params:   []any{0.1},
			Expected: "rectifier@0.1",
			Found:    true,
		},
		{
			Name:     "rectifier@0.1",
			Expected: "rectifier@0.1",
			Found:    true,
		},
		{
			Name:     Linear,
			Expected: "linear@1@0",
			Found:    true,
		},
		{
			Name:     "linear@2@1",
			Expected: "linear@2@1",
			Found:    true,
		},
		{
			Name:     "linear",
			Params:   []any{2.0, 1.0},
			Expected: "linear@2@1",
			Found:    true,
		},
		{
			Name:  "undefined",
			Found: false,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var a, found = Get(test.Name, test.Params...)
			if !test.Found {
				assert.False(t, found)
				return
			}
			assert.True(t, found)
			assert.Equal(t, test.Expected, a.Name)
			assert.NotNil(t, a.Function)
			assert.NotNil(t, a.Derivative)
		})
	}
}

func TestGetSigmoid(t *testing.T) {
	var a = GetSigmoid()
	assert.Equal(t, Sigmoid, a.Name)
	assert.NotNil(t, a.Function)
	assert.NotNil(t, a.Derivative)
	assert.Equal(t, 0.5, a.Function(0))
	assert.Equal(t, 0.25, a.Derivative(0.0))
}

func TestGetRectifier(t *testing.T) {
	var a = GetRectifier(0.1)
	assert.Equal(t, "rectifier@0.1", a.Name)
	assert.NotNil(t, a.Function)
	assert.NotNil(t, a.Derivative)
	assert.Equal(t, -0.1, a.Function(-1.0))
	assert.Equal(t, 0.0, a.Function(0.0))
	assert.Equal(t, 1.0, a.Function(1.0))
	assert.Equal(t, 0.1, a.Derivative(-1.0))
	assert.Equal(t, 1.0, a.Derivative(0.0))
	assert.Equal(t, 1.0, a.Derivative(2.0))
}

func TestGetLinear(t *testing.T) {
	var a = GetLinear(2.5, 1.0)
	assert.Equal(t, "linear@2.5@1", a.Name)
	assert.NotNil(t, a.Function)
	assert.NotNil(t, a.Derivative)
	assert.Equal(t, 0.0, a.Function(-0.4))
	assert.Equal(t, 2.5, a.Derivative(-0.4))
}

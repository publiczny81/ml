package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer(t *testing.T) {
	var (
		actual   = Pointer(1)
		expected = 1
	)
	assert.NotNil(t, actual)
	assert.Equal(t, &expected, actual)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunction(t *testing.T) {

	sum, err := MyFunction(2, 8)
	assert.Equal(t, 5, sum, "Sum of 2 and 3 is 5")
	assert.NoError(t, err, "There are no errors")

}

func TestMyFunction2(t *testing.T) {

	sum, err := MyFunction(-2, 3)
	assert.Equal(t, 0, sum, "Sum of -2 and 3 is 0")
	assert.Error(t, err, "Error is not nil for negative input")

}

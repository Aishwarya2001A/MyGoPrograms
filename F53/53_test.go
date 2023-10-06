package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateArea(t *testing.T) {

	mockObj := new(mockCalculateArea)

	mockObj.On("calculateArea", 5, 10).Return(50)

	actualArea := mockObj.calculateArea(5, 10)
	assert.Equal(t, 50, actualArea, "The calculated area is incorrect")

	mockObj.AssertExpectations(t)
}

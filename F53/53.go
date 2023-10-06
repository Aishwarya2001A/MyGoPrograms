package main

import (
	"github.com/stretchr/testify/mock"
)

type mockCalculateArea struct {
	mock.Mock
}

func (m *mockCalculateArea) calculateArea(width int, height int) int {
	args := m.Called(width, height)
	return args.Int(0)
}

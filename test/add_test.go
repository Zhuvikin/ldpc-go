package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	core "../matrix"
)

func TestAdd(t *testing.T) {
	matrix1 := core.CreateMatrix(3,3)
	matrix1.Set(0, 1)
	matrix1.Set(0, 0)
	matrix1.Set(2, 2)
	matrix1.Print("Matrix 1")

	matrix2 := core.CreateMatrix(3,3)
	matrix2.Set(0, 0)
	matrix2.Print("Matrix 2")

	sum := matrix1.Add(matrix2)
	sum.Print("Sum")

	expected := core.CreateMatrix(3,3)
	expected.Set(0, 1)
	expected.Set(2, 2)

	assert.True(t, sum.Equals(expected))
}

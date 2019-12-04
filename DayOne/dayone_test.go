package DayOne

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve_ShouldReturn2_WhenInputIs12(t *testing.T) {
	result := Solve(12)

	assert.Equal(t, 2, result)
}

func TestSolve_ShouldReturn2_WhenInputIs14(t *testing.T) {
	result := Solve(14)

	assert.Equal(t, 2, result)
}

func TestSolve_ShouldReturn654_WhenInputIs1969(t *testing.T) {
	result := Solve(1969)

	assert.Equal(t, 654, result)
}

func TestSolve_ShouldReturn33583_WhenInputIs100756(t *testing.T) {
	result := Solve(1969)

	assert.Equal(t, 654, result)
}

func TestSolveRecursive_ShouldReturn2_WhenInputIs14(t *testing.T) {
	result := SolveRecursive(14)

	assert.Equal(t, 2, result)
}

func TestSolveRecursive_ShouldReturn966_WhenInputIs1969(t *testing.T) {
	result := SolveRecursive(1969)

	assert.Equal(t, 966, result)
}

func TestSolveRecursive_ShouldReturn50346_WhenInputIs100756(t *testing.T) {
	result := SolveRecursive(100756)

	assert.Equal(t, 50346, result)
}
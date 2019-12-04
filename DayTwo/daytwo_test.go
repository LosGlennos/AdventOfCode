package DayTwo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve_ShouldReturnCorrectResult_GivenFirstScenario(t *testing.T) {
	input := []int{1,0,0,0,99}
	result := Solve(0, input)

	assert.Equal(t, []int{2, 0, 0, 0, 99}, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenSecondScenario(t *testing.T) {
	input := []int{2,3,0,3,99}
	result := Solve(0, input)

	assert.Equal(t, []int{2,3,0,6,99}, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenThirdScenario(t *testing.T) {
	input := []int{2,4,4,5,99,0}
	result := Solve(0, input)

	assert.Equal(t, []int{2,4,4,5,99,9801}, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenFourthScenario(t *testing.T) {
	input := []int{1,1,1,4,99,5,6,0,99}
	result := Solve(0, input)

	assert.Equal(t, []int{30,1,1,4,2,5,6,0,99}, result)
}
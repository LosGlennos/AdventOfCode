package DayThree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve_ShouldReturnCorrectResult_GivenFirstScenario(t *testing.T) {
	firstWire := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	secondWire := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	result := Solve(firstWire, secondWire)

	assert.Equal(t, 159.0, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenSecondScenario(t *testing.T) {
	firstWire := []string{"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"}
	secondWire := []string{"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"}
	result := Solve(firstWire, secondWire)

	assert.Equal(t, 135.0, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenThirdScenario(t *testing.T) {
	firstWire := []string{"U10", "R5"}
	secondWire := []string{"R1","U11"}
	result := Solve(firstWire, secondWire)

	assert.Equal(t, 11.0, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenFourthScenario(t *testing.T) {
	firstWire := []string{"D10", "R5"}
	secondWire := []string{"R1","D11"}
	result := Solve(firstWire, secondWire)

	assert.Equal(t, 11.0, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenFifthScenario(t *testing.T) {
	firstWire := []string{"D10", "L5"}
	secondWire := []string{"L1","D11"}
	result := Solve(firstWire, secondWire)

	assert.Equal(t, 11.0, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenSixthScenario(t *testing.T) {
	firstWire := []string{"U10", "L5", "D9", "R6"}
	secondWire := []string{"L6", "U11", "R5", "D2"}
	result := Solve(firstWire, secondWire)

	assert.Equal(t, 11.0, result)
}

func TestSolve_ShouldReturnCorrectResult_GivenSeventhScenario(t *testing.T) {
	firstWire := []string{"R8","U5","L5","D3"}
	secondWire := []string{"U7","R6","D4","L4"}
	result := Solve(firstWire, secondWire)

	assert.Equal(t, 6.0, result)
}
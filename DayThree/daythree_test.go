package DayThree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve_ShouldReturnCorrectDistanceAndSteps_GivenFirstScenario(t *testing.T) {
	firstWire := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	secondWire := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	distance, steps := Solve(firstWire, secondWire)

	assert.Equal(t, 159.0, distance)
	assert.Equal(t, 610, steps)
}

func TestSolve_ShouldReturnCorrectDistance_GivenSecondScenario(t *testing.T) {
	firstWire := []string{"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"}
	secondWire := []string{"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"}
	distance, steps := Solve(firstWire, secondWire)

	assert.Equal(t, 135.0, distance)
	assert.Equal(t, 410, steps)
}


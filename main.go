package main

import (
	"AdventOfCode/DayTwo"
	"log"
)
import "AdventOfCode/DayOne"

func main() {
	SolveDayOne()
	SolveDayTwo()
}

func SolveDayTwo() {
	input := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,19,5,23,2,6,23,27,1,6,27,31,2,31,9,35,1,35,6,39,1,10,39,43,2,9,43,47,1,5,47,51,2,51,6,55,1,5,55,59,2,13,59,63,1,63,5,67,2,67,13,71,1,71,9,75,1,75,6,79,2,79,6,83,1,83,5,87,2,87,9,91,2,9,91,95,1,5,95,99,2,99,13,103,1,103,5,107,1,2,107,111,1,111,5,0,99,2,14,0,0}
	partOneResult := SolveDayTwoPartOne(input)
	partTwoResult := SolveDayTwoPartTwo(input)

	log.Printf("Day two, part one answer: %d", partOneResult[0])
	log.Printf("Day two, part two answer: %d, %d", partTwoResult[1], partTwoResult[2])
}

func SolveDayTwoPartTwo(input []int) []int {
	for i := 0; i < 100; i++ {
		for y := 0; y < 100; y++ {
			copiedArray := make([]int, len(input))
			copy(copiedArray, input)
			copiedArray[1] = i
			copiedArray[2] = y
			result := DayTwo.Solve(0, copiedArray)
			if result[0] == 19690720 {
				return copiedArray
			}
		}
	}

	panic("found nothing")
}

func SolveDayTwoPartOne(input []int) []int {
	copiedArray := make([]int, len(input))
	copy(copiedArray, input)
	copiedArray[1] = 12
	copiedArray[2] = 2
	return DayTwo.Solve(0, copiedArray)
}

func SolveDayOne() {
	input := []int{144365, 124674, 99039, 132924, 126960, 103950, 78451, 123596, 119950, 116772, 134137, 50247, 99543, 147151, 103063, 59247, 59281, 141531, 104417, 75105, 57868, 149148, 76973, 87424, 135220, 141885, 106241, 128482, 54020, 67575, 97719, 110237, 137361, 70772, 103397, 117471, 99611, 142905, 135345, 122338, 62708, 103663, 146189, 81657, 126628, 133113, 135399, 52731, 116597, 61749, 61519, 56234, 64306, 127237, 133320, 79782, 132431, 142449, 91926, 146277, 55314, 111507, 126347, 124086, 120868, 127433, 126838, 77814, 144388, 86786, 134780, 109082, 101772, 140013, 100282, 115632, 73057, 139318, 85633, 67693, 55545, 53545, 125871, 115201, 105202, 148104, 68677, 64761, 54368, 110380, 102082, 106684, 89933, 71703, 147332, 99699, 98447, 96963, 148686, 92651}
	partOneResult := SolveDayOnePartOne(input)
	partTwoResult := SolveDayOnePartTwo(input)

	log.Printf("Day one, part one answer: %d", partOneResult)
	log.Printf("Day one, part two answer: %d", partTwoResult)
}

func SolveDayOnePartTwo(input []int) int {
	partTwoSum := 0
	for _, e := range input {
		partTwoSum += DayOne.SolveRecursive(e)
	}

	return partTwoSum
}

func SolveDayOnePartOne(input []int) int {
	partOneSum := 0
	for _, e := range input {
		partOneSum += DayOne.Solve(e)
	}

	return partOneSum
}

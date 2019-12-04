package main

import "log"
import "AdventOfCode/DayOne"

func main() {
	SolveDayOne()
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
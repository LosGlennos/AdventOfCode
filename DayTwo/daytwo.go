package DayTwo

func Solve(startIndex int, input []int) []int {
		operator := input[startIndex]

		if operator == 99 {
			return input
		}

		firstValueIndex := input[startIndex + 1]
		secondValueIndex := input[startIndex + 2]
		replacementIndex := input[startIndex + 3]

		var returnArray []int
		switch operator {
			case 1:
				additionResult := addValues(input[firstValueIndex], input[secondValueIndex])
				newArray := replaceValueAtIndex(replacementIndex, additionResult, input)
				returnArray = Solve(startIndex + 4, newArray)
				break
			case 2:
				multiplicationResult := multiplyValues(input[firstValueIndex], input[secondValueIndex])
				newArray := replaceValueAtIndex(replacementIndex, multiplicationResult, input)
				returnArray = Solve(startIndex + 4, newArray)
				break
		}

		return returnArray
}

func addValues(first int, second int) int {
	return first + second
}

func multiplyValues(first int, second int) int {
	return first * second
}

func replaceValueAtIndex(index int, value int, array []int) []int {
	array[index] = value

	return array
}
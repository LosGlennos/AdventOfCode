package DayOne

import "math"

func Solve(mass int) int {
	part := float64(mass/3)
	rounded := int(math.Floor(part))
	fuelAmount := rounded - 2
	return fuelAmount
}

func SolveRecursive(mass int) int {
	total := 0
	fuel := Solve(mass)
	if fuel <= 0 {
		return total
	}
	total += fuel
	total += SolveRecursive(fuel)
	return total
}
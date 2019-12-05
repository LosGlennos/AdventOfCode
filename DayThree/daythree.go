package DayThree

import (
	"math"
	"strconv"
)

type Walker struct {
	StepsTaken      map[string]int
	Intersects      map[Coordinate]int
	Path            map[Coordinate]Metadata
	CurrentPosition Coordinate
}

type Coordinate struct {
	X float64
	Y float64
}

type Metadata struct {
	Label string
	Steps int
}

func Solve(firstWire []string, secondWire []string) (float64, int) {
	walker := NewWalker()
	walker.WalkWire(firstWire, "first")

	walker.Intersects = make(map[Coordinate]int)
	walker.CurrentPosition = Coordinate{X: 0, Y: 0}
	walker.WalkWire(secondWire, "second")

	shortestDistance := walker.GetShortestDistance()
	steps := walker.GetLeastAmountOfSteps()

	return shortestDistance, steps
}

func (w *Walker) GetShortestDistance() float64 {
	var shortestDistance float64

	for key := range w.Intersects {
		distance := distance(key)
		if shortestDistance == 0 {
			shortestDistance = distance
		}

		if distance < shortestDistance {
			shortestDistance = distance
		}
	}
	return shortestDistance
}

func (w *Walker) GetLeastAmountOfSteps() int {
	var steps int
	for k, v := range w.Intersects {
		if steps == 0 {
			steps = v + w.Path[k].Steps
		}

		if (v + w.Path[k].Steps) < steps {
			steps = v + w.Path[k].Steps
		}
	}

	return steps
}

func distance(coordinate Coordinate) float64 {
	return math.Abs(coordinate.Y) + math.Abs(coordinate.X)
}

func (w *Walker) WalkWire(wire []string, label string) {
	for _, instruction := range wire {
		direction := instruction[:1]
		steps, _ := strconv.Atoi(instruction[1:])
		w.Walk(direction, steps, label)
	}
}

func (w *Walker) Walk(direction string, steps int, label string) {
	if direction == "U" {
		w.WalkDirection(0, 1, steps, label)
	} else if direction == "D" {
		w.WalkDirection(0, -1, steps, label)
	} else if direction == "R" {
		w.WalkDirection(1, 0, steps, label)
	} else if direction == "L" {
		w.WalkDirection(-1, 0, steps, label)
	}
}

func (w *Walker) WalkDirection(xOffset float64, yOffset float64, steps int, label string) {
	for i := 0; i < steps; i++ {
		w.CurrentPosition.Y += yOffset
		w.CurrentPosition.X += xOffset
		w.StepsTaken[label] += 1
		metaData, exists := w.Path[w.CurrentPosition]
		if exists && label != metaData.Label {
			w.Intersects[w.CurrentPosition] = w.StepsTaken[label]
		} else {
			w.Path[w.CurrentPosition] = Metadata{Label: label, Steps: w.StepsTaken[label]}
		}
	}
}

func NewWalker() Walker {
	return Walker{
		StepsTaken:      make(map[string]int),
		Path:            make(map[Coordinate]Metadata),
		CurrentPosition: Coordinate{X: 0, Y: 0},
	}
}

package DayThree

import (
	"math"
	"strconv"
)

type Walker struct {
	Intersects      []Coordinate
	Path            map[Coordinate]string
	CurrentPosition Coordinate
}

type Coordinate struct {
	X float64
	Y float64
}

func Solve(firstWire []string, secondWire []string) float64 {
	walker := NewWalker()
	walker.WalkWire(firstWire, "first")

	walker.Intersects = []Coordinate{}
	walker.CurrentPosition = Coordinate{X: 0, Y: 0}
	walker.WalkWire(secondWire, "second")

	return GetShortestDistance(walker.GetIntersects())
}

func (w *Walker) GetIntersects() []Coordinate {
	return w.Intersects
}

func GetShortestDistance(intersects []Coordinate) float64 {
	shortestDistance := distance(intersects[0])

	for _, coordinate := range intersects {
		distance := distance(coordinate)
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}
	return shortestDistance
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
		storedLabel, exists := w.Path[w.CurrentPosition]
		if  exists && label != storedLabel {
			w.Intersects = append(w.Intersects, w.CurrentPosition)
		} else {
			w.Path[w.CurrentPosition] = label
		}
	}
}

func NewWalker() Walker {
	return Walker{
		Path:            make(map[Coordinate]string),
		CurrentPosition: Coordinate{X: 0, Y: 0},
	}
}

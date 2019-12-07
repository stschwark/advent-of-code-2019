package day03

import (
	"math"
	"strconv"
)

type Step struct {
	x int
	y int
}

func CreatePath(instructions []string) []Step {
	current := Step{0, 0}
	path := []Step{}
	for _, instruction := range instructions {
		var next Step
		direction := string(instruction[0])
		distance, err := strconv.Atoi(instruction[1:])
		if err == nil {
			for i := 0; i < distance; i++ {
				switch direction {
				case "R":
					next = Step{current.x + 1, current.y}
				case "D":
					next = Step{current.x, current.y - 1}
				case "L":
					next = Step{current.x - 1, current.y}
				case "U":
					next = Step{current.x, current.y + 1}
				}
				path = append(path, next)
				current = next
			}
		}
	}
	return path
}

func findIntersections(path1 []Step, path2 []Step) []Step {
	intersections := []Step{}

	path1Map := map[Step]bool{}
	for _, step := range path1 {
		path1Map[step] = true
	}

	for _, step := range path2 {
		if _, ok := path1Map[step]; ok {
			intersections = append(intersections, step)
		}
	}

	return intersections
}

func calculatManhattenDistances(steps []Step) []int {
	distances := []int{}
	for _, step := range steps {
		distance := abs(step.x) + abs(step.y)
		distances = append(distances, distance)
	}
	return distances
}

func min(numbers []int) int {
	var min int = math.MaxInt64
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CalculateManhattanDistanceForNearestIntersection(path1 []Step, path2 []Step) int {
	intersections := findIntersections(path1, path2)
	return min(calculatManhattenDistances(intersections))
}

func calculateNumberOfStepsOnPathToTarget(path []Step, target Step) int {
	for count, current := range path {
		if current == target {
			return count + 1
		}
	}
	panic("Target not on path")
}

func CalculateMinNumberOfStepsToReachAnIntersection(path1 []Step, path2 []Step) int {
	intersections := findIntersections(path1, path2)
	totalNumberOfSteps := []int{}
	for _, intersection := range intersections {
		stepsOnPath1 := calculateNumberOfStepsOnPathToTarget(path1, intersection)
		stepsOnPath2 := calculateNumberOfStepsOnPathToTarget(path2, intersection)
		totalNumberOfSteps = append(totalNumberOfSteps, stepsOnPath1+stepsOnPath2)
	}
	return min(totalNumberOfSteps)
}

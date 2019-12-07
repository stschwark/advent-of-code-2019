package day03

import (
	"advent-of-code-2019/utils"
	"strings"
	"testing"
)

func TestCalculateManhattanDistanceForNearestIntersection(t *testing.T) {
	tests := []struct {
		wire1    string
		wire2    string
		distance int
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for _, test := range tests {
		path1 := CreatePath(strings.Split(test.wire1, ","))
		path2 := CreatePath(strings.Split(test.wire2, ","))

		if got := CalculateManhattanDistanceForNearestIntersection(path1, path2); got != test.distance {
			t.Errorf("CalculateManhattanDistanceForNearestIntersection() = %v, want %v", got, test.distance)
		}
	}
}

func TestSolvePart1(t *testing.T) {
	paths := utils.ReadFromFileAndSplit("day03_input.txt", "\n")
	path1 := CreatePath(strings.Split(paths[0], ","))
	path2 := CreatePath(strings.Split(paths[1], ","))

	want := 446
	if got := CalculateManhattanDistanceForNearestIntersection(path1, path2); got != want {
		t.Errorf("CalculateManhattanDistanceForNearestIntersection() = %v, want %v", got, want)
	}
}

func TestCalculateMinNumberOfStepsToReachIntersection(t *testing.T) {
	tests := []struct {
		wire1 string
		wire2 string
		steps int
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 30},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}
	for _, test := range tests {
		path1 := CreatePath(strings.Split(test.wire1, ","))
		path2 := CreatePath(strings.Split(test.wire2, ","))

		if got := CalculateMinNumberOfStepsToReachAnIntersection(path1, path2); got != test.steps {
			t.Errorf("CalculateMinManhattanDistanceForNearestIntersection() = %v, want %v", got, test.steps)
		}
	}
}

func TestSolvePart2(t *testing.T) {
	paths := utils.ReadFromFileAndSplit("day03_input.txt", "\n")
	path1 := CreatePath(strings.Split(paths[0], ","))
	path2 := CreatePath(strings.Split(paths[1], ","))

	want := 9006
	if got := CalculateMinNumberOfStepsToReachAnIntersection(path1, path2); got != want {
		t.Errorf("CalculateManhattanDistanceForNearestIntersection() = %v, want %v", got, want)
	}
}

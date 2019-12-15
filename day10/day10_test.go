package day10

import (
	"advent-of-code-2019/utils"
	"reflect"
	"testing"
)

func TestMaxAsteroidsThatCanBeDetectedFromAStation(t *testing.T) {
	asteroids := []string{".#..#", ".....", "#####", "....#", "...##"}

	want := 8

	if got := MaxAsteroidsThatCanBeDetectedFromAStation(asteroids); got != want {
		t.Errorf("MaxAsteroidsThatCanBeDetectedFromAStation() = %v, want %v", got, want)
	}
}

func TestSolvePart1(t *testing.T) {
	asteroids := utils.ReadFromFileAndSplit("day10_input.txt", "\n")

	want := 280

	if got := MaxAsteroidsThatCanBeDetectedFromAStation(asteroids); got != want {
		t.Errorf("MaxAsteroidsThatCanBeDetectedFromAStation() = %v, want %v", got, want)
	}
}

func TestVaporizedAsteroids(t *testing.T) {
	asteroids := []string{".#....#####...#..", "##...##.#####..##", "##...#...#.#####.", "..#.....X...###..", "..#.#.....#....##"}

	want := []Position{Position{8, 1}, Position{9, 0}, Position{9, 1}, Position{10, 0}, Position{9, 2}, Position{11, 1}, Position{12, 1}, Position{11, 2}, Position{15, 1}, Position{12, 2}, Position{13, 2}, Position{14, 2}, Position{15, 2}, Position{12, 3}, Position{16, 4}, Position{15, 4}, Position{10, 4}, Position{4, 4}, Position{2, 4}, Position{2, 3}, Position{0, 2}, Position{1, 2}, Position{0, 1}, Position{1, 1}, Position{5, 2}, Position{1, 0}, Position{5, 1}, Position{6, 1}, Position{6, 0}, Position{7, 0}, Position{8, 0}, Position{10, 1}, Position{14, 0}, Position{16, 1}, Position{13, 3}, Position{14, 3}}

	if got := VaporizedAsteroids(asteroids, Position{8, 3}); !reflect.DeepEqual(got, want) {
		t.Errorf("VaporizedAsteroids() = %v, want %v", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	asteroids := utils.ReadFromFileAndSplit("day10_input.txt", "\n")
	vaporized := VaporizedAsteroids(asteroids, Position{20, 18})

	want := Position{7, 6}
	if got := vaporized[199]; got != want {
		t.Errorf("VaporizedAsteroids() = %v, want %v", got, want)
	}
}

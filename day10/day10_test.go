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

	want := []Location{Location{8, 1}, Location{9, 0}, Location{9, 1}, Location{10, 0}, Location{9, 2}, Location{11, 1}, Location{12, 1}, Location{11, 2}, Location{15, 1}, Location{12, 2}, Location{13, 2}, Location{14, 2}, Location{15, 2}, Location{12, 3}, Location{16, 4}, Location{15, 4}, Location{10, 4}, Location{4, 4}, Location{2, 4}, Location{2, 3}, Location{0, 2}, Location{1, 2}, Location{0, 1}, Location{1, 1}, Location{5, 2}, Location{1, 0}, Location{5, 1}, Location{6, 1}, Location{6, 0}, Location{7, 0}, Location{8, 0}, Location{10, 1}, Location{14, 0}, Location{16, 1}, Location{13, 3}, Location{14, 3}}

	if got := VaporizedAsteroids(asteroids, Location{8, 3}); !reflect.DeepEqual(got, want) {
		t.Errorf("VaporizedAsteroids() = %v, want %v", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	asteroids := utils.ReadFromFileAndSplit("day10_input.txt", "\n")
	vaporized := VaporizedAsteroids(asteroids, Location{20, 18})

	want := Location{7, 6}
	if got := vaporized[199]; got != want {
		t.Errorf("VaporizedAsteroids() = %v, want %v", got, want)
	}
}

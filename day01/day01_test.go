package day01

import (
	"advent-of-code-2019/utils"
	"strconv"
	"testing"
)

func TestCalculateFuelForMass(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, test := range tests {
		if fuel := CalculateFuelForMass(test.mass); fuel != test.fuel {
			t.Errorf("CalculateFuelForMass() = %v, want %v", fuel, test.fuel)
		}
	}
}

func TestSolvePart1(t *testing.T) {
	lines := utils.ReadFromFileAndSplit("day01_input.txt", "\n")

	masses := make([]int, len(lines))

	for i, line := range lines {
		mass, err := strconv.Atoi(line)
		if err == nil {
			masses[i] = mass
		}
	}

	want := 3289802
	if got := CalculateBaseFuel(masses); got != want {
		t.Errorf("CalculateBaseFuel() = %v, want %v", got, want)
	}
}

func TestCalculaCalculateFuelPerModule(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, test := range tests {
		if fuel := CalculateFuelPerModule(test.mass); fuel != test.fuel {
			t.Errorf("CalculateFuelPerModule() = %v, want %v", fuel, test.fuel)
		}
	}
}

func TestSolvePart2(t *testing.T) {
	lines := utils.ReadFromFileAndSplit("day01_input.txt", "\n")

	masses := make([]int, len(lines))

	for i, line := range lines {
		mass, err := strconv.Atoi(line)
		if err == nil {
			masses[i] = mass
		}
	}

	want := 4931831
	if got := CalculateTotalFuelRequirements(masses); got != want {
		t.Errorf("CalculateTotalFuelRequirements() = %v, want %v", got, want)
	}
}

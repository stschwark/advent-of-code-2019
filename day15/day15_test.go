package day15

import (
	"advent-of-code-2019/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day15_input.txt", ",")

	want := 220

	if program, err := utils.StringsToInts(input); err == nil {
		if _, got := ExploreShip(program); got != want {
			t.Errorf("ExploreShip() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day15_input.txt", ",")

	want := 334

	if program, err := utils.StringsToInts(input); err == nil {
		grid, _ := ExploreShip(program)

		if got := MinutesUntilFilledWithOxygen(grid); got != want {
			t.Errorf("MinutesUntilFilledWithOxygen() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

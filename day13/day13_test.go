package day13

import (
	"advent-of-code-2019/utils"
	"fmt"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day13_input.txt", ",")

	want := 255

	if program, err := utils.StringsToInts(input); err == nil {
		grid := draw(program)

		if got := CountType(grid, 2); got != want {
			t.Errorf("PanelsPaintedAtLeastOnce() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day13_input.txt", ",")
	if program, err := utils.StringsToInts(input); err == nil {
		want := 12338

		if got := Play(program); got != want {
			t.Errorf("Play() = %v, want %v", got, want)
		}
	} else {
		fmt.Println("Cannot read program")
	}
}

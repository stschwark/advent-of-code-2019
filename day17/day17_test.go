package day17

import (
	"advent-of-code-2019/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day17_input.txt", ",")

	if program, err := utils.StringsToInts(input); err == nil {
		image := TakeImage(program)

		//PrintImage(image)

		want := 7720
		if got := CalibrateCamera(image); got != want {
			t.Errorf("Calibrate() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day17_input.txt", ",")

	want := 0

	if program, err := utils.StringsToInts(input); err == nil {
		if got := NotifyOtherRobots(program); got != want {
			t.Errorf("Calibrate() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

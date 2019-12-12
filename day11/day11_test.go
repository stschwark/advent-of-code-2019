package day11

import (
	"advent-of-code-2019/utils"
	"reflect"
	"testing"
)

func TestPanelsPaintedAtLeastOnce(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day11_input.txt", ",")

	want := 0

	if program, err := utils.StringsToInts(input); err == nil {
		if got := PanelsPaintedAtLeastOnce(program); got != want {
			t.Errorf("PanelsPaintedAtLeastOnce() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

func TestPaintedRegistration(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day11_input.txt", ",")

	want := []string{
		"  ##  ###  #### #  # ####  ##  ####  ##    ",
		" #  # #  # #    # #     # #  # #    #  #   ",
		" #  # ###  ###  ##     #  #    ###  #      ",
		" #### #  # #    # #   #   # ## #    # ##   ",
		" #  # #  # #    # #  #    #  # #    #  #   ",
		" #  # ###  #### #  # ####  ### #     ###   ",
	}

	if program, err := utils.StringsToInts(input); err == nil {
		if got := PaintedRegistration(program); !reflect.DeepEqual(got, want) {
			t.Errorf("PaintedRegistration() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

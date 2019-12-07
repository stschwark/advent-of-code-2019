package day02

import (
	"advent-of-code-2019/utils"
	"reflect"
	"testing"
)

func TestRunPart1(t *testing.T) {
	tests := []struct {
		name   string
		start  []int
		result []int
	}{
		{"(30 + 40) * 50 = 3500", []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{"1 + 1 = 2", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"3 * 2 = 6", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"99 * 99 = 9801", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"5 * 6 = 30", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := Run(test.start); !reflect.DeepEqual(result, test.result) {
				t.Errorf("Run() = %v, want %v", result, test.result)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	codes := utils.ReadFromFileAndSplit("day02_input.txt", ",")
	want := 6730673

	if memory, err := utils.StringsToInts(codes); err == nil {
		if result := RunWithNounAndVerb(memory, 12, 2); result[0] != want {
			t.Errorf("Run() = %v, want %v", result[0], want)
		}
	} else {
		t.FailNow()
	}
}

func TestSolvePart2(t *testing.T) {
	codes := utils.ReadFromFileAndSplit("day02_input.txt", ",")
	want := 19690720

	if memory, err := utils.StringsToInts(codes); err == nil {
		if result := RunWithNounAndVerb(memory, 37, 49); result[0] != want {
			t.Errorf("RunWithNounAndVerb() = %v, want %v", result, want)
		}
	}
}

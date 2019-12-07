package day05

import (
	"advent-of-code-2019/utils"
	"reflect"
	"testing"
)

func TestRunPart1(t *testing.T) {
	tests := []struct {
		name   string
		start  []int
		input  int
		result []int
		output int
	}{
		{"(30 + 40) * 50 = 3500", []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, 0, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, 0},
		{"1 + 1 = 2", []int{1, 0, 0, 0, 99}, 0, []int{2, 0, 0, 0, 99}, 0},
		{"3 * 2 = 6", []int{2, 3, 0, 3, 99}, 0, []int{2, 3, 0, 6, 99}, 0},
		{"99 * 99 = 9801", []int{2, 4, 4, 5, 99, 0}, 0, []int{2, 4, 4, 5, 99, 9801}, 0},
		{"5 * 6 = 30", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 0, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, output := Run(test.start, test.input)

			if !reflect.DeepEqual(result, test.result) || output != test.output {
				t.Errorf("Run() = (%v,%v), want (%v,%v)", result, output, test.result, test.output)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day05_input.txt", ",")

	want := 15508323
	if program, err := utils.StringsToInts(input); err == nil {
		testSystemID := 1

		if _, output := Run(program, testSystemID); output != want {
			t.Errorf("Run() = %v, want %v", output, want)
		}
	} else {
		t.Fail()
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day05_input.txt", ",")

	want := 9006327
	if program, err := utils.StringsToInts(input); err == nil {
		testSystemID := 5

		if _, output := Run(program, testSystemID); output != want {
			t.Errorf("Run() = %v, want %v", output, want)
		}
	} else {
		t.Fail()
	}
}

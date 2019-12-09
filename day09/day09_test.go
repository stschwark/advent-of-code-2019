package day09

import (
	"advent-of-code-2019/utils"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		program []int
		want    []int
	}{
		{"Example 1", []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
		{"Example 2", []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, []int{1219070632396864}},
		{"Example 2", []int{104, 1125899906842624, 99}, []int{1125899906842624}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Run(test.program, nil); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Run() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day09_input.txt", ",")

	if program, err := utils.StringsToInts(input); err == nil {
		want := []int{2457252183}
		if got := Run(program, []int{1}); !reflect.DeepEqual(got, want) {
			t.Errorf("Run() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day09_input.txt", ",")

	if program, err := utils.StringsToInts(input); err == nil {
		want := []int{70634}
		if got := Run(program, []int{2}); !reflect.DeepEqual(got, want) {
			t.Errorf("Run() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

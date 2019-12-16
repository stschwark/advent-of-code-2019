package day16

import (
	"advent-of-code-2019/utils"
	"strings"
	"testing"
)

var basePattern = []int{0, 1, 0, -1}

func TestFFT(t *testing.T) {
	type args struct {
		signal string
		phases int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1 - 1 phase", args{"12345678", 1}, "48226158"},
		{"Example 1 - 2 phases", args{"12345678", 2}, "34040438"},
		{"Example 1 - 3 phases", args{"12345678", 3}, "03415518"},
		{"Example 1 - 4 phases", args{"12345678", 4}, "01029498"},
		{"Example 2 - 100 phases", args{"80871224585914546619083218645595", 100}, "24176176"},
		{"Example 3 - 100 phases", args{"19617804207202209144916044189917", 100}, "73745418"},
		{"Example 4 - 100 phases", args{"69317163492948606335995924319873", 100}, "52432133"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := FFT(test.args.signal, test.args.phases, basePattern)[:8]; got != test.want {
				t.Errorf("FFT() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	input := utils.ReadFromFile("day16_input.txt")

	want := "30369587"

	if got := FFT(input, 100, basePattern)[:8]; got != want {
		t.Errorf("FFT() = %v, want %v", got, want)
	}
}

func xTestSolvePart2(t *testing.T) {
	input := strings.Repeat(utils.ReadFromFile("day16_input.txt"), 10000)

	want := "30369587"

	if got := FFT(input, 100, basePattern)[:8]; got != want {
		t.Errorf("FFT() = %v, want %v", got, want)
	}
}

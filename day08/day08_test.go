package day08

import (
	"advent-of-code-2019/utils"
	"reflect"
	"testing"
)

func TestCheckSum(t *testing.T) {
	width, height := 3, 2
	image := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}

	want := 1

	if got := CheckSum(image, width, height); got != want {
		t.Errorf("CheckSum() = %v, want %v", got, want)
	}
}

func TestSolvePart1(t *testing.T) {
	width, height := 25, 6

	input := utils.ReadFromFileAndSplit("day08_input.txt", "")
	if image, err := utils.StringsToInts(input); err == nil {
		want := 2806

		if got := CheckSum(image, width, height); got != want {
			t.Errorf("CheckSum() = %v, want %v", got, want)
		}
	} else {
		t.Fail()
	}
}

func TestMergeLayers(t *testing.T) {
	width, height := 2, 2
	image := []int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}

	want := []int{0, 1, 1, 0}

	if got := MergeLayers(image, width, height); !reflect.DeepEqual(got, want) {
		t.Errorf("MergeLayers() = %v, want %v", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	width, height := 25, 6

	input := utils.ReadFromFileAndSplit("day08_input.txt", "")
	if image, err := utils.StringsToInts(input); err == nil {
		want := []int{1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0}

		mergedImage := MergeLayers(image, width, height)
		//PrintImage(mergedImage, width, height)

		if !reflect.DeepEqual(mergedImage, want) {
			t.Errorf("MergeLayers() = %v, want %v", mergedImage, want)
		}
	} else {
		t.Fail()
	}
}

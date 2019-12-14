package day14

import (
	"advent-of-code-2019/utils"
	"testing"
)

func TestExample1Part1(t *testing.T) {
	input := []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
		"7 A, 1 C => 1 D",
		"7 A, 1 D => 1 E",
		"7 A, 1 E => 1 FUEL",
	}

	receipe := ParseRecipe(input)

	want := 31

	if got := CalculateOreRequiredFor1Fuel(receipe); got != want {
		t.Errorf("CalculateOreRequiredFor1Fuel() = %v, want %v", got, want)
	}
}

func TestExample2Part1(t *testing.T) {
	input := []string{
		"157 ORE => 5 NZVS",
		"165 ORE => 6 DCFZ",
		"44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL",
		"12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ", "179 ORE => 7 PSHF",
		"177 ORE => 5 HKGWZ",
		"7 DCFZ, 7 PSHF => 2 XJWVT",
		"165 ORE => 2 GPVTF",
		"3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT",
	}

	receipe := ParseRecipe(input)

	want := 13312

	if got := CalculateOreRequiredFor1Fuel(receipe); got != want {
		t.Errorf("CalculateOreRequiredFor1Fuel() = %v, want %v", got, want)
	}
}

func TestSolvePart1(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day14_input.txt", "\n")

	receipe := ParseRecipe(input)

	want := 216477

	if got := CalculateOreRequiredFor1Fuel(receipe); got != want {
		t.Errorf("CalculateOreRequiredFor1Fuel() = %v, want %v", got, want)
	}
}

func TestExample2Part2(t *testing.T) {
	input := []string{
		"157 ORE => 5 NZVS",
		"165 ORE => 6 DCFZ",
		"44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL",
		"12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ", "179 ORE => 7 PSHF",
		"177 ORE => 5 HKGWZ",
		"7 DCFZ, 7 PSHF => 2 XJWVT",
		"165 ORE => 2 GPVTF",
		"3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT",
	}

	receipe := ParseRecipe(input)

	want := 82892753

	if got := CalculateFuelCreatedFromOre(receipe, 13312, 1000000000000); got != want {
		t.Errorf("CalculateFuelCreatedFromOre() = %v, want %v", got, want)
	}
}

func TestExample3Part2(t *testing.T) {
	input := []string{
		"2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG",
		"17 NVRVD, 3 JNWZP => 8 VPVL",
		"53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL",
		"22 VJHF, 37 MNCFX => 5 FWMGM",
		"139 ORE => 4 NVRVD",
		"144 ORE => 7 JNWZP",
		"5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC",
		"5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV",
		"145 ORE => 6 MNCFX",
		"1 NVRVD => 8 CXFTF",
		"1 VJHF, 6 MNCFX => 4 RFSQX",
		"176 ORE => 6 VJHF",
	}

	receipe := ParseRecipe(input)

	want := 5586022

	if got := CalculateFuelCreatedFromOre(receipe, 180697, 1000000000000); got != want {
		t.Errorf("CalculateFuelCreatedFromOre() = %v, want %v", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadFromFileAndSplit("day14_input.txt", "\n")

	receipe := ParseRecipe(input)

	want := 11788286

	if got := CalculateFuelCreatedFromOre(receipe, 216477, 1000000000000); got != want {
		t.Errorf("CalculateFuelCreatedFromOre() = %v, want %v", got, want)
	}
}

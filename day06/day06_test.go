package day06

import (
	"advent-of-code-2019/utils"
	"testing"
)

func TestCalculateNumberOfDirectAndIndirectOrbits(t *testing.T) {
	orbits := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}

	want := 42

	if got := CalculateNumberOfDirectAndIndirectOrbits(orbits); got != want {
		t.Errorf("CalculateNumberOfDirectAndIndirectOrbits() = %v, want %v", got, want)
	}
}

func TestSolvePar1(t *testing.T) {
	orbits := utils.ReadFromFileAndSplit("day06_input.txt", "\n")

	want := 145250

	if got := CalculateNumberOfDirectAndIndirectOrbits(orbits); got != want {
		t.Errorf("CalculateNumberOfDirectAndIndirectOrbits() = %v, want %v", got, want)
	}
}

func TestCalculateMinimumNumberOfOrbitalTransfers(t *testing.T) {
	orbits := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}

	want := 4

	if got := CalculateMinimumNumberOfOrbitalTransfers(orbits, "SAN", "YOU"); got != want {
		t.Errorf("CalculateNumberOfDirectAndIndirectOrbits() = %v, want %v", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	orbits := utils.ReadFromFileAndSplit("day06_input.txt", "\n")

	want := 274

	if got := CalculateMinimumNumberOfOrbitalTransfers(orbits, "SAN", "YOU"); got != want {
		t.Errorf("CalculateNumberOfDirectAndIndirectOrbits() = %v, want %v", got, want)
	}
}

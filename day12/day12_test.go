package day12

import "testing"

func moon(px, py, pz int) []Vector {
	return []Vector{Vector{px, 0}, Vector{py, 0}, Vector{pz, 0}}
}

func TestCalculateTotalEnergy(t *testing.T) {
	moons := [][]Vector{
		moon(-1, 0, 2),
		moon(2, -10, -7),
		moon(4, -8, 8),
		moon(3, 5, -1),
	}

	want := 179

	after := Simulate(moons, 10)

	if got := CalculateTotalEnergy(after); got != want {
		t.Errorf("CalculateTotalEnergy() = %v, want %v", got, want)
	}
}

func TestSolvePart1(t *testing.T) {
	moons := [][]Vector{
		moon(4, 12, 13),
		moon(-9, 14, -3),
		moon(-7, -1, 2),
		moon(-11, 17, -1),
	}

	want := 5350

	after := Simulate(moons, 1000)

	if got := CalculateTotalEnergy(after); got != want {
		t.Errorf("CalculateTotalEnergy() = %v, want %v", got, want)
	}
}

func TestCalculateCycleCount(t *testing.T) {
	moons := [][]Vector{
		moon(-1, 0, 2),
		moon(2, -10, -7),
		moon(4, -8, 8),
		moon(3, 5, -1),
	}

	want := 2772

	if got := CalculateCycleCount(moons); got != want {
		t.Errorf("TestCalculateCycleCount() = %v, want %v", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	moons := [][]Vector{
		moon(4, 12, 13),
		moon(-9, 14, -3),
		moon(-7, -1, 2),
		moon(-11, 17, -1),
	}

	want := 467034091553512

	if got := CalculateCycleCount(moons); got != want {
		t.Errorf("TestCalculateCycleCount() = %v, want %v", got, want)
	}
}

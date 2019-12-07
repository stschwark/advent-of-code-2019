package day04

import "testing"

func TestIsValidPasswordForPart1(t *testing.T) {
	tests := []struct {
		name     string
		password int
		isValid  bool
	}{
		{"double digit, never decreases", 111111, true},
		{"decreasing pair of digits", 223450, false},
		{"no double", 123789, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsValidPasswordForPart1(test.password); got != test.isValid {
				t.Errorf("IsValidPassword() = %v, want %v", got, test.isValid)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	want := 979
	if got := CoundValidPasswordsInRangeForPart1(256310, 732736); got != want {
		t.Errorf("CoundValidPasswordsInRangeForPart1() = %v, want %v", got, want)
	}
}

func TestIsValidPasswordForPart2(t *testing.T) {
	tests := []struct {
		name     string
		password int
		isValid  bool
	}{
		{"double digit, never decreases", 112233, true},
		{"no double", 123444, false},
		{"one double, never decreases", 111122, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsValidPasswordForPart2(test.password); got != test.isValid {
				t.Errorf("IsValidPassword() = %v, want %v", got, test.isValid)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	want := 635
	if got := CoundValidPasswordsInRangeForPart2(256310, 732736); got != want {
		t.Errorf("CoundValidPasswordsInRangeForPart2() = %v, want %v", got, want)
	}
}

package day04

import "sort"

func intToDigits(n int) []int {
	digits := []int{}
	for ; n != 0; n = n / 10 {
		digits = append([]int{n % 10}, digits...)
	}
	return digits
}

func hasGroupsOfTwoOrMore(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] == digits[i-1] {
			return true
		}
	}
	return false
}

func IsValidPasswordForPart1(password int) bool {
	digits := intToDigits(password)

	digitsAreIncreasing := sort.IntsAreSorted(digits)
	hasGroupsOfTwoOrMoreDigits := hasGroupsOfTwoOrMore(digits)

	return digitsAreIncreasing && hasGroupsOfTwoOrMoreDigits
}

func CoundValidPasswordsInRangeForPart1(start int, end int) int {
	validPasswords := 0
	for password := start; password <= end; password++ {
		if IsValidPasswordForPart1(password) {
			validPasswords++
		}
	}
	return validPasswords
}

func hasDouble(digits []int) bool {
	seen := 1

	for i := 1; i < len(digits); i++ {
		if digits[i] == digits[i-1] {
			seen++
		} else {
			if seen == 2 {
				return true
			}
			seen = 1
		}
	}
	return seen == 2
}

func IsValidPasswordForPart2(password int) bool {
	digits := intToDigits(password)

	digitsAreIncreasing := sort.IntsAreSorted(digits)
	hasDoubleDigits := hasDouble(digits)

	return digitsAreIncreasing && hasDoubleDigits
}

func CoundValidPasswordsInRangeForPart2(start int, end int) int {
	validPasswords := 0
	for password := start; password <= end; password++ {
		if IsValidPasswordForPart2(password) {
			validPasswords++
		}
	}
	return validPasswords
}

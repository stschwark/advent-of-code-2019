package utils

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FilterInt(values []int, f func(int) bool) []int {
	filtered := make([]int, 0)
	for _, value := range values {
		if f(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

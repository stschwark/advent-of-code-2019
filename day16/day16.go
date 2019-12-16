package day16

import (
	"advent-of-code-2019/utils"
	"strconv"
	"strings"
)

func repeat(val int, n int) (result []int) {
	for i := 0; i < n; i++ {
		result = append(result, val)
	}
	return result
}

func generatePattern(basePattern []int, repetitions int, length int) []int {
	pattern := make([]int, 0, length)
	for i := 0; i < length; i++ {
		pattern = append(pattern, repeat(basePattern[i%len(basePattern)], repetitions)...)
	}
	return pattern
}

func join(values []int) (joined string) {
	for _, value := range values {
		joined += strconv.Itoa(value)
	}
	return joined
}

func FFT(signal string, phases int, basePattern []int) string {
	list, err := utils.StringsToInts(strings.Split(signal, ""))
	if err != nil {
		panic("Invalid signal")
	}

	patterns := make([][]int, len(list))
	for i := range list {
		patterns[i] = generatePattern(basePattern, i+1, len(list)+1)[1:]
	}

	for phase := 0; phase < phases; phase++ {
		for i := range list {
			sum := 0
			for j, element := range list {
				factor := patterns[i][j]
				sum += element * factor
			}
			list[i] = utils.AbsInt(sum % 10)
		}
	}

	return join(list)
}

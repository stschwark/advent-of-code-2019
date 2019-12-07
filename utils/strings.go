package utils

import "strconv"

func StringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))

	for i, str := range strings {
		if number, err := strconv.Atoi(str); err != nil {
			return nil, err
		} else {
			ints[i] = number
		}
	}

	return ints, nil
}

package day17

import (
	"advent-of-code-2019/intcode"
	"fmt"
)

var scaffold = 35

func isScaffold(image [][]int, x int, y int) bool {
	if y >= 0 && y < len(image) && x >= 0 && x < len(image[0]) {
		return image[y][x] == scaffold
	}
	return false
}

func isJunction(image [][]int, x int, y int) bool {
	return isScaffold(image, y-1, x) &&
		isScaffold(image, y, x+1) &&
		isScaffold(image, y+1, x) &&
		isScaffold(image, y, x-1)
}

func PrintImage(image [][]int) {
	for _, row := range image {
		for _, pixel := range row {
			fmt.Printf("%c", pixel)
		}
		fmt.Println()
	}
}

func TakeImage(program []int) [][]int {
	image := [][]int{}

	row := 0

	output := func(val int) {
		if val == 10 {
			row++
		} else {
			if row > len(image)-1 {
				image = append(image, []int{})
			}
			image[row] = append(image[row], val)
		}
	}

	intcode.Run(program, func() int { return 0 }, output, func() bool { return false })

	return image
}

func CalibrateCamera(image [][]int) int {
	calibration := 0
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[0]); x++ {
			if isJunction(image, x, y) {
				calibration += x * y
			}
		}
	}

	return calibration
}

func NotifyOtherRobots(program []int) (dust int) {
	program[0] = 2

	movementInstructions := []byte("A,B,A,B,C,C,B,A,B,C\n" +
		"L,10,R,10,L,10,L,10\n" +
		"R,10,R,12,L,12\n" +
		"R,12,L,12,R,6\n" +
		"n\n")

	input := func() int {
		next := movementInstructions[0]
		movementInstructions = movementInstructions[1:]
		return int(next)
	}

	output := func(val int) {
		dust = val
	}

	intcode.Run(program, input, output, func() bool { return false })

	return dust
}

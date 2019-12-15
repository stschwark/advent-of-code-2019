package day13

import (
	"advent-of-code-2019/intcode"
	"advent-of-code-2019/utils"
)

func draw(program []int) utils.Grid {
	grid := utils.Grid{}

	input := func() int {
		panic("Unexpected input request")
	}

	sequence := []int{}

	output := func(val int) {
		if len(sequence) == 2 {
			grid[utils.Position{sequence[0], sequence[1]}] = val
			sequence = sequence[:0]
		} else {
			sequence = append(sequence, val)
		}
	}

	intcode.Run(program, input, output, nil)

	return grid
}

func CountType(m utils.Grid, t int) (count int) {
	for _, v := range m.Values() {
		if v == t {
			count++
		}
	}
	return count
}

func tileIdToDisplay(tileId int) string {
	switch tileId {
	case 1:
		return "#"
	case 2:
		return "x"
	case 3:
		return "-"
	case 4:
		return "o"
	default:
		return " "
	}
}

func Play(program []int) (score int) {
	program[0] = 2

	grid := utils.Grid{}

	ballX := 0
	paddleX := 0

	input := func() int {
		// utils.Clear()
		// grid.Print(tileIdToDisplay)

		switch {
		case paddleX < ballX:
			return 1
		case paddleX > ballX:
			return -1
		default:
			return 0
		}
	}

	sequence := []int{}

	output := func(val int) {
		if len(sequence) == 2 {
			x := sequence[0]
			y := sequence[1]
			if x == -1 && y == 0 {
				score = val
			} else {
				grid[utils.Position{x, y}] = val

				switch val {
				case 3:
					paddleX = x
				case 4:
					ballX = x
				}
			}
			sequence = sequence[:0]
		} else {
			sequence = append(sequence, val)
		}
	}

	intcode.Run(program, input, output, nil)

	return score
}

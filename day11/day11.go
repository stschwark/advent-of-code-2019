package day11

import (
	"advent-of-code-2019/intcode"
	"advent-of-code-2019/utils"
	"math"
)

type Position struct {
	x, y int
}

type Robot struct {
	position  Position
	direction int
}

func (r *Robot) turnLeft() {
	r.direction = (r.direction + 3) % 4
}

func (r *Robot) turnRight() {
	r.direction = (r.direction + 1) % 4
}

func (r *Robot) moveForward() {
	switch r.direction {
	case 0:
		r.position.y--
	case 1:
		r.position.x++
	case 2:
		r.position.y++
	case 3:
		r.position.x--
	}
}

func (r *Robot) Paint(program []int, startPanelColour int) map[Position]int {
	grid := make(map[Position]int)
	grid[Position{0, 0}] = startPanelColour

	input := func() int {
		colour, ok := grid[r.position]
		if ok {
			return colour
		}
		return 0
	}

	first := true
	output := func(val int) {
		if first {
			grid[r.position] = val
		} else {
			if val == 0 {
				r.turnLeft()
			} else {
				r.turnRight()
			}
			r.moveForward()
		}
		first = !first
	}

	intcode.Run(program, input, output, nil)

	return grid
}

func PanelsPaintedAtLeastOnce(program []int) int {
	robot := Robot{}
	grid := robot.Paint(program, 0)
	return len(grid)
}

func snapshot(grid map[Position]int) []string {
	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, -math.MaxInt64, -math.MaxInt64
	for p := range grid {
		minX = utils.MinInt(minX, p.x)
		minY = utils.MinInt(minY, p.y)
		maxX = utils.MaxInt(maxX, p.x)
		maxY = utils.MaxInt(maxY, p.y)
	}

	result := []string{}
	for y := minY; y <= maxY; y++ {
		row := ""
		for x := minX; x <= maxX; x++ {
			if colour, ok := grid[Position{x, y}]; ok && colour == 1 {
				row += "#"
			} else {
				row += " "
			}
		}
		result = append(result, row)
	}
	return result
}

func PaintedRegistration(program []int) []string {
	robot := Robot{}
	grid := robot.Paint(program, 1)
	return snapshot(grid)
}

package day15

import (
	"advent-of-code-2019/intcode"
	"advent-of-code-2019/utils"
	"math/rand"
	"reflect"
	"time"
)

const (
	unexplored = 0
	space      = 1
	wall       = 2
	droid      = 3
	oxygen     = 4
)

const (
	north = 1
	east  = 4
	south = 2
	west  = 3
)

var start = utils.Position{}

var directions = []int{north, east, south, west}

func ahead(position utils.Position, direction int) utils.Position {
	switch direction {
	case north:
		return utils.Position{position.X, position.Y - 1}
	case east:
		return utils.Position{position.X + 1, position.Y}
	case south:
		return utils.Position{position.X, position.Y + 1}
	case west:
		return utils.Position{position.X - 1, position.Y}
	default:
		panic("Invalid direction")
	}
}

func leftOf(direction int) int {
	switch direction {
	case north:
		return west
	case east:
		return north
	case south:
		return east
	default:
		return south
	}
}

func rightOf(direction int) int {
	switch direction {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	default:
		return north
	}
}

func paint(p utils.Position, val int) string {
	if p == start {
		return "x"
	}

	switch val {
	case space:
		return "."
	case wall:
		return "#"
	case oxygen:
		return "O"
	default:
		return " "
	}
}

func ExploreShip(program []int) (utils.Grid, int) {
	rand.Seed(time.Now().UnixNano())

	stepsToPosition := make(map[utils.Position]int)
	stepsToOxygen := 0

	backAtStart := false
	oxygenDiscovered := false

	grid := utils.Grid{}
	position := utils.Position{}

	direction := north

	input := func() int {
		if grid[ahead(position, leftOf(direction))] != wall {
			direction = leftOf(direction)
		}
		return direction
	}

	output := func(val int) {
		if oxygenDiscovered && reflect.DeepEqual(position, start) {
			backAtStart = true
		}

		next := ahead(position, direction)

		switch val {
		case 0: //robot hit wall
			grid[next] = wall
			direction = rightOf(direction)
		case 1: //robot moved forward
			grid[next] = space
			position = next

			if !oxygenDiscovered {
				if steps, ok := stepsToPosition[position]; ok {
					stepsToOxygen = steps
				} else {
					stepsToOxygen++
					stepsToPosition[position] = stepsToOxygen
				}
			}
		case 2: //robot found oxygen system
			oxygenDiscovered = true
			stepsToOxygen++
			grid[next] = oxygen
			position = next
		default:
			panic("Invalid output")
		}
	}

	interrupt := func() bool {
		return oxygenDiscovered && backAtStart
	}

	intcode.Run(program, input, output, interrupt)

	//grid.Print(paint)

	return grid, stepsToOxygen
}

func MinutesUntilFilledWithOxygen(grid utils.Grid) int {
	minutes := 0
	for {
		unfilled := len(utils.FilterInt(grid.Values(), func(value int) bool { return value == space }))
		if unfilled == 0 {
			break
		}

		nextGrid := utils.Grid{}
		for k, v := range grid {
			nextGrid[k] = v
		}

		for position := range grid {
			if grid[position] == oxygen {
				for _, direction := range directions {
					if grid[ahead(position, direction)] == space {
						nextGrid[ahead(position, direction)] = oxygen
					}
				}
			}
		}

		grid = nextGrid

		minutes++
	}

	return minutes
}

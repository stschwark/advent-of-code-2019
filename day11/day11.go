package day11

import (
	"fmt"
	"math"
)

const positionMode = 0
const immediateMode = 1
const relativeMode = 2

type shipComputer struct {
	input          func() int
	memory         []int
	extendedMemory map[int]int
	pos            int
	output         func(int)
	relativeBase   int
}

func (sc *shipComputer) instruction() int {
	return sc.memory[sc.pos]
}

func (sc *shipComputer) opcode() int {
	return sc.instruction() % 100
}

func (sc *shipComputer) parameterMode(offset int) int {
	return (sc.instruction() / int(math.Pow10(offset+1))) % 10
}

func (sc *shipComputer) read(param int) int {
	switch sc.parameterMode(param) {
	case positionMode:
		return sc.readAbsolute(sc.readRelative(param))
	case immediateMode:
		return sc.readRelative(param)
	case relativeMode:
		return sc.readAbsolute(sc.relativeBase + sc.readRelative(param))
	default:
		panic("Unknown read parameterMode")
	}
}

func (sc *shipComputer) readRelative(offset int) int {
	return sc.readAbsolute(sc.pos + offset)
}

func (sc *shipComputer) readAbsolute(pos int) int {
	if pos < len(sc.memory) {
		return sc.memory[pos]
	}

	return sc.extendedMemory[pos]
}

func (sc *shipComputer) write(param int, value int) {
	switch sc.parameterMode(param) {
	case positionMode:
		sc.writeAbsolute(sc.readRelative(param), value)
	case relativeMode:
		sc.writeAbsolute(sc.relativeBase+sc.readRelative(param), value)
	default:
		panic("Unknown write parameterMode")
	}
}

func (sc *shipComputer) writeAbsolute(pos int, value int) {
	if pos < len(sc.memory) {
		sc.memory[pos] = value
	} else {
		sc.extendedMemory[pos] = value
	}
}

func (sc *shipComputer) step() int {
	opcode := sc.opcode()

	switch opcode {
	case 1:
		value := sc.read(1) + sc.read(2)
		sc.write(3, value)
		sc.pos += 4
	case 2:
		v1, v2 := sc.read(1), sc.read(2)
		sc.write(3, v1*v2)
		sc.pos += 4
	case 3:
		val := sc.input()
		sc.write(1, val)
		sc.pos += 2
	case 4:
		val := sc.read(1)
		sc.output(val)
		sc.pos += 2
	case 5:
		if sc.read(1) != 0 {
			sc.pos = sc.read(2)
		} else {
			sc.pos += 3
		}
	case 6:
		if sc.read(1) == 0 {
			sc.pos = sc.read(2)
		} else {
			sc.pos += 3
		}
	case 7:
		if sc.read(1) < sc.read(2) {
			sc.write(3, 1)
		} else {
			sc.write(3, 0)
		}
		sc.pos += 4
	case 8:
		if sc.read(1) == sc.read(2) {
			sc.write(3, 1)
		} else {
			sc.write(3, 0)
		}
		sc.pos += 4
	case 9:
		sc.relativeBase += sc.read(1)
		sc.pos += 2
	case 99:
	default:
		panic(fmt.Sprintf("unknown opcode %v", opcode))
	}

	return opcode
}

func run(program []int, input func() int, output func(int)) {
	sc := shipComputer{}
	sc.memory = make([]int, len(program))
	copy(sc.memory, program)
	sc.extendedMemory = make(map[int]int)
	sc.input = input
	sc.output = output

	for {
		opcode := sc.step()

		if opcode == 99 {
			break
		}
	}
}

type Location struct {
	x, y int
}

type Robot struct {
	location  Location
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
		r.location.y--
	case 1:
		r.location.x++
	case 2:
		r.location.y++
	case 3:
		r.location.x--
	}
}

func (r *Robot) Paint(program []int, startPanelColour int) map[Location]int {
	grid := make(map[Location]int)
	grid[Location{0, 0}] = startPanelColour

	input := func() int {
		colour, ok := grid[r.location]
		if ok {
			return colour
		}
		return 0
	}

	first := true
	output := func(val int) {
		if first {
			grid[r.location] = val
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

	run(program, input, output)

	return grid
}

func PanelsPaintedAtLeastOnce(program []int) int {
	robot := Robot{}
	grid := robot.Paint(program, 0)
	return len(grid)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func snapshot(grid map[Location]int) []string {
	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, -math.MaxInt64, -math.MaxInt64
	for l := range grid {
		minX = min(minX, l.x)
		minY = min(minY, l.y)
		maxX = max(maxX, l.x)
		maxY = max(maxY, l.y)
	}

	result := []string{}
	for y := minY; y <= maxY; y++ {
		row := ""
		for x := minX; x <= maxX; x++ {
			if colour, ok := grid[Location{x, y}]; ok && colour == 1 {
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

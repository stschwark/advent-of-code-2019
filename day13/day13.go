package day13

import (
	"advent-of-code-2019/utils"
	"fmt"
	"math"
	"os"
	"os/exec"
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

type Position struct {
	x, y int
}

func draw(program []int) map[Position]int {
	grid := make(map[Position]int)

	input := func() int {
		panic("Unexpected input request")
	}

	sequence := []int{}

	output := func(val int) {
		if len(sequence) == 2 {
			grid[Position{sequence[0], sequence[1]}] = val
			sequence = sequence[:0]
		} else {
			sequence = append(sequence, val)
		}
	}

	run(program, input, output)

	return grid
}

func CountType(m map[Position]int, t int) (count int) {
	for _, v := range valuesOf(m) {
		if v == t {
			count++
		}
	}
	return count
}

func valuesOf(m map[Position]int) (values []int) {
	for key := range m {
		values = append(values, m[key])
	}
	return values
}

func print(grid map[Position]int) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, -math.MaxInt64, -math.MaxInt64
	for l := range grid {
		minX = utils.MinInt(minX, l.x)
		minY = utils.MinInt(minY, l.y)
		maxX = utils.MaxInt(maxX, l.x)
		maxY = utils.MaxInt(maxY, l.y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			switch grid[Position{x, y}] {
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("x")
			case 3:
				fmt.Print("-")
			case 4:
				fmt.Print("o")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Play(program []int) (score int) {
	program[0] = 2

	grid := make(map[Position]int)

	ballX := 0
	paddleX := 0

	input := func() int {
		//print(grid)

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
				grid[Position{x, y}] = val

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

	run(program, input, output)

	return score
}

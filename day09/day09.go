package day09

import (
	"fmt"
	"math"
)

type stack []int

func (s *stack) Pop() int {
	res := (*s)[0]
	*s = (*s)[1:len(*s)]
	return res
}

const positionMode = 0
const immediateMode = 1
const relativeMode = 2

type shipComputer struct {
	input          stack
	memory         []int
	extendedMemory map[int]int
	pos            int
	output         []int
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
		val := sc.input.Pop()
		sc.write(1, val)
		sc.pos += 2
	case 4:
		val := sc.read(1)
		sc.output = append(sc.output, val)
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

func Run(program []int, input []int) (output []int) {
	sc := shipComputer{}
	sc.memory = make([]int, len(program))
	copy(sc.memory, program)
	sc.extendedMemory = make(map[int]int)
	sc.input = input

	for {
		opcode := sc.step()

		if opcode == 99 {
			return sc.output
		}
	}
}

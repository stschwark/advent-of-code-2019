package day05

import (
	"fmt"
	"math"
)

type shipComputer struct {
	input  int
	memory []int
	pos    int
	output int
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

func (sc *shipComputer) read(offset int) int {
	switch sc.parameterMode(offset) {
	case 0:
		return sc.memory[sc.readImmediate(offset)]
	case 1:
		return sc.readImmediate(offset)
	}
	panic("Unknown parameterMode")
}

func (sc *shipComputer) readImmediate(offset int) int {
	return sc.memory[sc.pos+offset]
}

func (sc *shipComputer) write(offset int, value int) {
	sc.memory[sc.readImmediate(offset)] = value
}

func (sc *shipComputer) step() int {
	opcode := sc.opcode()

	switch opcode {
	case 1:
		value := sc.read(1) + sc.read(2)
		sc.write(3, value)
		sc.pos += 4
	case 2:
		value := sc.read(1) * sc.read(2)
		sc.write(3, value)
		sc.pos += 4
	case 3:
		sc.write(1, sc.input)
		sc.pos += 2
	case 4:
		sc.output = sc.read(1)
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
	case 99:
	default:
		panic(fmt.Sprintf("unknown opcode %v", opcode))
	}

	return opcode
}

func Run(code []int, input int) ([]int, int) {
	sc := shipComputer{}
	sc.memory = make([]int, len(code))
	copy(sc.memory, code)
	sc.input = input

	for {
		opcode := sc.step()

		if opcode == 99 {
			return sc.memory, sc.output
		}
	}
}

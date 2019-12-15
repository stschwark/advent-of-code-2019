package intcode

import (
	"fmt"
	"math"
)

const positionMode = 0
const immediateMode = 1
const relativeMode = 2

type computer struct {
	input          func() int
	memory         []int
	extendedMemory map[int]int
	pos            int
	output         func(int)
	relativeBase   int
}

func (c *computer) instruction() int {
	return c.memory[c.pos]
}

func (c *computer) opcode() int {
	return c.instruction() % 100
}

func (c *computer) parameterMode(offset int) int {
	return (c.instruction() / int(math.Pow10(offset+1))) % 10
}

func (c *computer) read(param int) int {
	switch c.parameterMode(param) {
	case positionMode:
		return c.readAbsolute(c.readRelative(param))
	case immediateMode:
		return c.readRelative(param)
	case relativeMode:
		return c.readAbsolute(c.relativeBase + c.readRelative(param))
	default:
		panic("Unknown read parameterMode")
	}
}

func (c *computer) readRelative(offset int) int {
	return c.readAbsolute(c.pos + offset)
}

func (c *computer) readAbsolute(pos int) int {
	if pos < len(c.memory) {
		return c.memory[pos]
	}

	return c.extendedMemory[pos]
}

func (c *computer) write(param int, value int) {
	switch c.parameterMode(param) {
	case positionMode:
		c.writeAbsolute(c.readRelative(param), value)
	case relativeMode:
		c.writeAbsolute(c.relativeBase+c.readRelative(param), value)
	default:
		panic("Unknown write parameterMode")
	}
}

func (c *computer) writeAbsolute(pos int, value int) {
	if pos < len(c.memory) {
		c.memory[pos] = value
	} else {
		c.extendedMemory[pos] = value
	}
}

func (c *computer) step() int {
	opcode := c.opcode()

	switch opcode {
	case 1:
		value := c.read(1) + c.read(2)
		c.write(3, value)
		c.pos += 4
	case 2:
		v1, v2 := c.read(1), c.read(2)
		c.write(3, v1*v2)
		c.pos += 4
	case 3:
		val := c.input()
		c.write(1, val)
		c.pos += 2
	case 4:
		val := c.read(1)
		c.output(val)
		c.pos += 2
	case 5:
		if c.read(1) != 0 {
			c.pos = c.read(2)
		} else {
			c.pos += 3
		}
	case 6:
		if c.read(1) == 0 {
			c.pos = c.read(2)
		} else {
			c.pos += 3
		}
	case 7:
		if c.read(1) < c.read(2) {
			c.write(3, 1)
		} else {
			c.write(3, 0)
		}
		c.pos += 4
	case 8:
		if c.read(1) == c.read(2) {
			c.write(3, 1)
		} else {
			c.write(3, 0)
		}
		c.pos += 4
	case 9:
		c.relativeBase += c.read(1)
		c.pos += 2
	case 99:
	default:
		panic(fmt.Sprintf("unknown opcode %v", opcode))
	}

	return opcode
}

func Run(intcode []int, input func() int, output func(int)) {
	c := computer{}
	c.memory = make([]int, len(intcode))
	copy(c.memory, intcode)
	c.extendedMemory = make(map[int]int)
	c.input = input
	c.output = output

	for {
		opcode := c.step()

		if opcode == 99 {
			break
		}
	}
}

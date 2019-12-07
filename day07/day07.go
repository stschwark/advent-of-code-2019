package day07

import (
	"fmt"
	"math"
)

type shipComputer struct {
	id     int
	input  chan int
	memory []int
	pos    int
	output chan int
	onExit chan bool
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
		val := <-sc.input
		sc.write(1, val)
		sc.pos += 2
	case 4:
		val := sc.read(1)
		sc.output <- val
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
		if sc.onExit != nil {
			sc.onExit <- true
		}
	default:
		panic(fmt.Sprintf("unknown opcode %v", opcode))
	}

	return opcode
}

func Run(id int, code []int, input chan int, output chan int, onExit chan bool) {
	sc := shipComputer{}
	sc.id = id
	sc.memory = make([]int, len(code))
	copy(sc.memory, code)
	sc.input = input
	sc.output = output
	sc.onExit = onExit

	for {
		opcode := sc.step()

		if opcode == 99 {
			return
		}
	}
}

func Permutations(of []int) (permutations [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permutations = append(permutations, append([]int{}, a...))
		} else {
			for i := k; i < len(of); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(of, 0)

	return permutations
}

func AmplifyChained(program []int, phases []int, input chan int) (output chan int) {
	for id, phase := range phases {
		output = make(chan int)

		go Run(id, program, input, output, nil)

		input <- phase

		input = output
	}

	return output
}

func MaxSignalChained(program []int) (maxSignal int) {
	candidates := Permutations([]int{0, 1, 2, 3, 4})
	for _, candidate := range candidates {
		input := make(chan int)
		output := AmplifyChained(program, candidate, input)
		input <- 0
		close(input)

		if signal := <-output; signal > maxSignal {
			maxSignal = signal
		}
	}
	return maxSignal
}

func AmplifyLooped(program []int, phases []int, input chan int) chan int {
	c1 := make(chan int)
	go Run(0, program, input, c1, nil)
	input <- phases[0]

	onExit := make(chan bool)

	c2 := make(chan int)
	go Run(1, program, c1, c2, onExit)
	c1 <- phases[1]

	c3 := make(chan int)
	go Run(2, program, c2, c3, nil)
	c2 <- phases[2]

	c4 := make(chan int)
	go Run(3, program, c3, c4, nil)
	c3 <- phases[3]

	go Run(4, program, c4, input, nil)
	c4 <- phases[4]

	input <- 0

	<-onExit

	return input
}

func MaxSignalLooped(program []int) (maxSignal int) {
	candidates := Permutations([]int{5, 6, 7, 8, 9})
	for _, candidate := range candidates {
		input := make(chan int)
		if signal := <-AmplifyLooped(program, candidate, input); signal > maxSignal {
			maxSignal = signal
		}
	}
	return maxSignal
}

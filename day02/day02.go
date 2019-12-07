package day02

func isInstructionValid(memory []int, address int) bool {
	return address+3 < len(memory) && memory[address+1] < len(memory) && memory[address+2] < len(memory) && memory[address+3] < len(memory)
}

func Run(input []int) []int {
	memory := make([]int, len(input))
	copy(memory, input)

	pos := 0
	for {
		switch memory[pos] {
		case 1:
			if !isInstructionValid(memory, pos) {
				return input
			}
			memory[memory[pos+3]] = memory[memory[pos+1]] + memory[memory[pos+2]]
		case 2:
			if !isInstructionValid(memory, pos) {
				return input
			}
			memory[memory[pos+3]] = memory[memory[pos+1]] * memory[memory[pos+2]]
		case 99:
			return memory
		}
		pos += 4
	}
}

func RunWithNounAndVerb(input []int, noun int, verb int) []int {
	input[1] = noun
	input[2] = verb

	return Run(input)
}

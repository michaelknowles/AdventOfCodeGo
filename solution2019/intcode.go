package solution2019

import "AdventOfCodeGo/util"

func instructionArity(pointer int) (count int) {
	switch pointer {
	case 1, 2:
		count = 3
	case 99:
		count = 0
	}
	return
}

// Intcode is the computer that runs a given Intcode program
func Intcode(input []int) (output int, state []int) {
	program := util.CloneIntSlice(input)
	pointer := 0
	instruction := program[pointer]
	arity := instructionArity(instruction)
	for {
		switch instruction {
		case 1:
			a1 := program[pointer+1]
			a2 := program[pointer+2]
			r := program[pointer+3]
			sum := program[a1] + program[a2]
			program[r] = sum
		case 2:
			a1 := program[pointer+1]
			a2 := program[pointer+2]
			r := program[pointer+3]
			product := program[a1] * program[a2]
			program[r] = product
		}
		pointer += arity+1
		instruction = program[pointer]
		if instruction == 99 {
			break
		}
	}
	output = program[0]
	state = program
	return
}
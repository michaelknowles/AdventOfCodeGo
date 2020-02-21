package solution2019

import (
	"AdventOfCodeGo/util"
	"fmt"
	"strconv"
	"strings"
)

func day2part1(input []int) (output int) {
	program := util.CloneIntSlice(input)
	program[1] = 12
	program[2] = 2
	output, _ = Intcode(program)
	return
}

func day2part2(input []int) (total int) {
	program := util.CloneIntSlice(input)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program[1] = noun
			program[2] = verb
			output, _ := Intcode(program)
			if output == 19690720 {
				total = (100 * noun) + verb
				return
			}
		}
	}
	fmt.Println("Answer not found")
	return
}

// Day2 returns a map containing solutions
func Day2() map[string]int {
	lines, err := util.ReadFile(2019, 2)
	if err != nil {
		panic(err)
	}
	input := make([]int, 0)
	for _, v := range strings.Split(lines[0], ",") {
		x, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		input = append(input, x)
	}

	solutions := make(map[string]int)
	solutions["Part 1"] = day2part1(input)
	solutions["Part 2"] = day2part2(input)
	return solutions
}
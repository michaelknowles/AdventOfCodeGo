package solution2019

import (
	"AdventOfCodeGo/util"
	"strconv"
)

// CalcFuel takes a mass and returns the fuel requirement.
// It divides the mass by 3, rounds down, then subtracts 2
func calcFuel(mass int) (fuel int) {
	fuel = (mass / 3) - 2
	return
}

// CalcFuelRec takes a mass and returns the fuel requirement recursively.
// It works like CalcFuel but calculates recursively.
// A negative fuel will be 0 instead
func calcFuelRec(mass int) (fuel int) {
	for mass > 8 {
		mass = calcFuel(mass)
		fuel += mass
	}
	return
}

func day1part1(input []int) (total int) {
	for _, mass := range input {
		total += calcFuel(mass)
	}
	return
}

func day1part2(input []int) (total int) {
	for _, mass := range input {
		total += calcFuelRec(mass)
	}
	return
}

// Day1 returns a map containing solutions
func Day1() map[string]int {
	lines, err := util.ReadFile(2019, 1)
	if err != nil {
		panic(err)
	}

	input := make([]int, 0)
	for _, v := range lines {
		mass, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		input = append(input, mass)
	}

	solutions := make(map[string]int)
	solutions["Part 1"] = day1part1(input)
	solutions["Part 2"] = day1part2(input)
	return solutions
}
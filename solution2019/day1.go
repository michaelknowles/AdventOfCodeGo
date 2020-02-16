package solution2019

import (
	"AdventOfCodeGo/util"
	"strconv"
)

// CalcFuel takes a mass and returns the fuel requirement.
// It divides the mass by 3, rounds down, then subtracts 2
func CalcFuel(mass int) (fuel int) {
	fuel = (mass / 3) - 2
	return
}

// CalcFuelRec takes a mass and returns the fuel requirement recursively.
// It works like CalcFuel but calculates recursively.
// A negative fuel will be 0 instead
func CalcFuelRec(mass int) (fuel int) {
	for mass > 8 {
		mass = CalcFuel(mass)
		fuel += mass
	}
	return
}

func setup() (input []int) {
	lines, err := util.ReadFile(2019, 1)
	if err != nil {
		panic(err)
	}

	for _, v := range lines {
		mass, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		input = append(input, mass)
	}
	return
}

func part1(input []int) (total int) {
	for _, mass := range input {
		total += CalcFuel(mass)
	}
	return
}

func part2(input []int) (total int) {
	for _, mass := range input {
		total += CalcFuelRec(mass)
	}
	return
}

// Day1 returns a map containing solutions
func Day1() map[string]int {
	input := setup()
	solutions := make(map[string]int)
	solutions["Part 1"] = part1(input)
	solutions["Part 2"] = part2(input)
	return solutions
}
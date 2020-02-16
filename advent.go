package main

import (
	"AdventOfCodeGo/solution2019"
	"fmt"
)

func main() {
	fmt.Println(("Hi!"))
	day1 := solution2019.Day1()
	for k, v := range day1 {
		fmt.Println(k, v)
	}
}

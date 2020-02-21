package util

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile reads the file for the given year/day and returns a slice of strings
func ReadFile(year int, day int) (lines []string, err error) {
	fileName := fmt.Sprintf("../resources/%d-day%d.txt", year, day)
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

// CloneIntSlice copies a slice into another
func CloneIntSlice(input []int) []int {
	if input == nil {
		return nil
	} 
	clone := make([]int, len(input))
	copy(clone, input)
	return clone
}

// CloneStrSlice copies a slice into another
func CloneStrSlice(input []string) []string {
	if input == nil {
		return nil
	} 
	clone := make([]string, len(input))
	copy(clone, input)
	return clone
}
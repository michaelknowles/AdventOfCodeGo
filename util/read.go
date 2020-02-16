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
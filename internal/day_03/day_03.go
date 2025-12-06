package day_03

import (
	"github.com/lorentzforces/advent-2025/internal/puzzle_tools"
)

func PartOne(input string) (int, error) {
	lines := puzzle_tools.AsLines(input)

	runningTotal := 0
	for _, line := range lines {
		first, index := findLargestInRange(line, 0, len(line) - 1)
		second, _ := findLargestInRange(line, index + 1, len(line))

		lineVal := first * 10 + second
		runningTotal += lineVal
	}
	return runningTotal, nil
}

// find the largest digit in the specified range
// start is inclusive, end is exclusive
func findLargestInRange(s string, start int, end int) (val int, index int) {
	largestVal := int(s[start] - '0')
	index = start

	for i := start; i < end; i++ {
		val := int(s[i] - '0')
		if val > largestVal {
			largestVal = val
			index = i
		}
	}

	return largestVal, index
}

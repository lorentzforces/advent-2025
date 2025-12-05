package day_01

import (
	"fmt"
	"strconv"

	"github.com/lorentzforces/advent-2025/internal/puzzle_tools"
)

func PartOne(input string) (int, error) {
	moves, err := parseMoves(input)
	if err != nil { return 0, err }

	stopsAtZero := 0
	dial := 50
	for _, move := range moves {
		dial += move
		if dial % 100 == 0 {
			stopsAtZero++
		}
	}

	return stopsAtZero, nil
}

func PartTwo(input string) (int, error) {
	moves, err := parseMoves(input)
	if err != nil { return 0, err }

	pointsAtZero := 0
	dial := 50
	for _, move := range moves {
		startingValue := dial
		fullRotations := move / 100
		pointsAtZero += max(fullRotations, -fullRotations)

		remainder := move - (fullRotations * 100)
		dial += remainder
		if dial == 0 {
			pointsAtZero++
		} else if dial < 0 {
			if startingValue != 0 {
				pointsAtZero++
			}
			dial += 100
		} else if dial > 99 {
			pointsAtZero++
			dial -= 100
		}
	}

	return pointsAtZero, nil
}

func parseMoves(s string) ([]int, error) {
	lines := puzzle_tools.AsLines(s)
	moves := make([]int, 0, len(lines))

	for _, line := range lines {
		runes := []rune(line)
		var direction int
		switch runes[0] {
		case 'L': direction = -1
		case 'R': direction = 1
		default: return nil, fmt.Errorf("Could not parse direction for line \"%s\"", line)
		}

		distance, err := strconv.Atoi(string(runes[1:]))
		if err != nil {
			return nil, fmt.Errorf("Error parsing numeric value for line \"%s\": %w", line, err)
		}
		if distance < 0 {
			return nil, fmt.Errorf("Found negative value %d for line \"%s\"", distance, line)
		}

		moves = append(moves, distance * direction)
	}
	return moves, nil
}

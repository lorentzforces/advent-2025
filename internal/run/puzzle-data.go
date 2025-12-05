package run

import "github.com/lorentzforces/advent-2025/internal/day_01"

type PuzzleData struct {
	Day int
	Part int
	InputFile string
	Fn PuzzleRunFunc
}

var listOfPuzzles = []PuzzleData{
	{
		Day: 1,
		Part: 1,
		InputFile: "day1-input.txt",
		Fn: func(s string) (any, error) { return day_01.PartOne(s) },
	},
}

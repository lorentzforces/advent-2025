package run

import (
	"github.com/lorentzforces/advent-2025/internal/day_01"
	"github.com/lorentzforces/advent-2025/internal/day_02"
	"github.com/lorentzforces/advent-2025/internal/day_03"
	"github.com/lorentzforces/advent-2025/internal/day_04"
)

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
	{
		Day: 1,
		Part: 2,
		InputFile: "day1-input.txt",
		Fn: func(s string) (any, error) { return day_01.PartTwo(s) },
	},
	{
		Day: 2,
		Part: 1,
		InputFile: "day2-input.txt",
		Fn: func(s string) (any, error) { return day_02.PartOne(s) },
	},
	{
		Day: 2,
		Part: 2,
		InputFile: "day2-input.txt",
		Fn: func(s string) (any, error) { return day_02.PartTwo(s) },
	},
	{
		Day: 3,
		Part: 1,
		InputFile: "day3-input.txt",
		Fn: func(s string) (any, error) { return day_03.PartOne(s) },
	},
	{
		Day: 3,
		Part: 2,
		InputFile: "day3-input.txt",
		Fn: func(s string) (any, error) { return day_03.PartTwo(s) },
	},
	{
		Day: 4,
		Part: 1,
		InputFile: "day4-input.txt",
		Fn: func(s string) (any, error) { return day_04.PartOne(s) },
	},
	{
		Day: 4,
		Part: 2,
		InputFile: "day4-input.txt",
		Fn: func(s string) (any, error) { return day_04.PartTwo(s) },
	},
}

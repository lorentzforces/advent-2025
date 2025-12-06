package day_04

import (
	"github.com/lorentzforces/advent-2025/internal/spatial"
)

func PartOne(input string) (int, error) {
	grid := spatial.ReadGrid(input)

	neighborCounts := make([][]int, grid.Height)
	for i := range len(neighborCounts) {
		neighborCounts[i] = make([]int, grid.Width)
	}

	for y := range grid.Height {
		for x := range grid.Width {
			currentSpace := spatial.Vec2d{ X: x, Y: y }
			for _, direction := range eightWayDirections {
				charAt := grid.CharAt(currentSpace.Add(direction))
				if charAt == '@' {
					neighborCounts[y][x] += 1
				}
			}
		}
	}

	accessibleRolls := 0
	for y := range grid.Height {
		for x := range grid.Width {
			if grid.CharAtXY(x, y) == '@' && neighborCounts[y][x] < 4 {
				accessibleRolls += 1
			}
		}
	}

	return accessibleRolls, nil
}

var eightWayDirections = []spatial.Vec2d{
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: 1, Y: -1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
}

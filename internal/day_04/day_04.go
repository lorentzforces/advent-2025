package day_04

import (
	"github.com/lorentzforces/advent-2025/internal/spatial"
)

func PartOne(input string) (int, error) {
	grid := spatial.ReadGrid(input)

	neighborCounts := initNeighborCounts(grid)

	for y := range grid.Height {
		for x := range grid.Width {
			currentSpace := spatial.Vec2d{ X: x, Y: y }
			if grid.CharAt(currentSpace) == '@' {
				neighborCounts.incrementNeighbors(currentSpace)
			}
		}
	}

	accessibleRolls := 0
	for y := range grid.Height {
		for x := range grid.Width {
			coords := spatial.Vec2d{ X: x, Y: y }
			if grid.CharAt(coords) == '@' && neighborCounts.countAt(coords) < 4 {
				accessibleRolls += 1
			}
		}
	}

	return accessibleRolls, nil
}

func PartTwo(input string) (int, error) {
	grid := spatial.ReadGrid(input)

	neighborCounts := initNeighborCounts(grid)

	numberOfRolls := 0
	rollLocations := make([]spatial.Vec2d, grid.Height * grid.Width)

	for y := range grid.Height {
		for x := range grid.Width {
			currentSpace := spatial.Vec2d{ X: x, Y: y }
			if grid.CharAt(currentSpace) == '@' {
				rollLocations[numberOfRolls] = currentSpace
				numberOfRolls++
				neighborCounts.incrementNeighbors(currentSpace)
			}
		}
	}

	initialNumberOfRolls := numberOfRolls

	changed := true
	for changed {
		changed = false
		for rollIndex := numberOfRolls - 1; rollIndex >= 0; rollIndex-- {
			rollLoc := rollLocations[rollIndex]
			if neighborCounts.countAt(rollLoc) < 4 {
				rollLocations[rollIndex], rollLocations[numberOfRolls - 1] =
					rollLocations[numberOfRolls - 1], rollLocations[rollIndex]
				numberOfRolls--
				neighborCounts.decrementNeighbors(rollLoc)
				changed = true
			}
		}
	}

	return initialNumberOfRolls - numberOfRolls, nil
}

type neighborCounts struct {
	gridCounts [][]int
}

func initNeighborCounts(grid spatial.Grid) neighborCounts {
	newCounts := neighborCounts{}
	newCounts.gridCounts = make([][]int, grid.Height)
	for i := range len(newCounts.gridCounts) {
		newCounts.gridCounts[i] = make([]int, grid.Width)
	}

	return newCounts
}

func (self neighborCounts) countAt(coords spatial.Vec2d) int {
	withinBounds :=
		coords.X >= 0 && coords.X < len(self.gridCounts[0]) &&
		coords.Y >= 0 && coords.Y < len(self.gridCounts)
	if !withinBounds {
		return 0
	}
	return self.gridCounts[coords.Y][coords.X]
}

func (self neighborCounts) incrementNeighbors(center spatial.Vec2d) {
	for _, direction := range eightWayDirections {
		neighborCoord := center.Add(direction)
		withinBounds :=
			neighborCoord.X >= 0 && neighborCoord.X < len(self.gridCounts[0]) &&
			neighborCoord.Y >= 0 && neighborCoord.Y < len(self.gridCounts)
		if withinBounds {
			self.gridCounts[neighborCoord.Y][neighborCoord.X] += 1
		}
	}
}

func (self neighborCounts) decrementNeighbors(center spatial.Vec2d) {
	for _, direction := range eightWayDirections {
		neighborCoord := center.Add(direction)
		withinBounds :=
			neighborCoord.X >= 0 && neighborCoord.X < len(self.gridCounts[0]) &&
			neighborCoord.Y >= 0 && neighborCoord.Y < len(self.gridCounts)
		if withinBounds {
			self.gridCounts[neighborCoord.Y][neighborCoord.X] -= 1
		}
	}
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

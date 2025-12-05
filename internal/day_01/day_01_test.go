package day_01

import (
	"testing"

	"github.com/lorentzforces/advent-2025/internal/puzzle_tools"
	"github.com/stretchr/testify/assert"
)

var testInput string =
`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(testInput)
	assert.NoError(t, err)
	puzzle_tools.BailIfFailed(t)
	assert.Equal(t, 3, result)
}

func TestPartTwoSampleInput(t *testing.T) {
	result, err := PartTwo(testInput)
	assert.NoError(t, err)
	puzzle_tools.BailIfFailed(t)
	assert.Equal(t, 6, result)
}

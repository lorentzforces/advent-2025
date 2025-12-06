package day_04

import (
	"testing"

	"github.com/lorentzforces/advent-2025/internal/puzzle_tools"
	"github.com/stretchr/testify/assert"
)

var sampleInput string =
`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(sampleInput)
	assert.NoError(t, err)
	puzzle_tools.BailIfFailed(t)
	assert.Equal(t, 13, result)
}

package day_03

import (
	"testing"

	"github.com/lorentzforces/advent-2025/internal/puzzle_tools"
	"github.com/stretchr/testify/assert"
)

var sampleInput string =
`987654321111111
811111111111119
234234234234278
818181911112111
`

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(sampleInput)
	assert.NoError(t, err)
	puzzle_tools.BailIfFailed(t)
	assert.Equal(t, 357, result)
}

func TestPartTwoSampleInput(t *testing.T) {
	result, err := PartTwo(sampleInput)
	assert.NoError(t, err)
	puzzle_tools.BailIfFailed(t)
	assert.Equal(t, uint64(3_121_910_778_619), result)
}

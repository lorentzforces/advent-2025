package day_02

import (
	"testing"

	"github.com/lorentzforces/advent-2025/internal/puzzle_tools"
	"github.com/stretchr/testify/assert"
)

var testInput =
	"11-22,95-115,998-1012,1188511880-1188511890,222220-222224," +
	"1698522-1698528,446443-446449,38593856-38593862,565653-565659," +
	"824824821-824824827,2121212118-2121212124"

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(testInput)
	assert.NoError(t, err)
	puzzle_tools.BailIfFailed(t)
	assert.Equal(t, uint(1227775554), result)
}

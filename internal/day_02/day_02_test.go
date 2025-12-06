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

func TestPartTwoSampleInput(t *testing.T) {
	result, err := PartTwo(testInput)
	assert.NoError(t, err)
	puzzle_tools.BailIfFailed(t)
	assert.Equal(t, uint(4174379265), result)
}

// the slowest - on my machine the associated puzzle solution took around 110 ms
func TestRepeatedDigitsFunc(t *testing.T) {
	assert.Equal(t, true, isRepeatedDigitsNumber(121212))
	assert.Equal(t, false, isRepeatedDigitsNumber(12121))
	assert.Equal(t, true, isRepeatedDigitsNumber(11))
	assert.Equal(t, true, isRepeatedDigitsNumber(33333))
	assert.Equal(t, true, isRepeatedDigitsNumber(100100100))
}

// by far the fastest by a factor of 3-4, usually taking around 35 ms
func TestRepeatedDigitsQuantitativeFunc(t *testing.T) {
	assert.Equal(t, true, isRepeatedDigitsNumberQuant(121212))
	assert.Equal(t, false, isRepeatedDigitsNumberQuant(12121))
	assert.Equal(t, true, isRepeatedDigitsNumberQuant(11))
	assert.Equal(t, true, isRepeatedDigitsNumberQuant(33333))
	assert.Equal(t, true, isRepeatedDigitsNumberQuant(100100100))
}

// consistently quicker than the initial version, but not by much (usually around 105 ms)
func TestRepeatedDigitsSliceFunc(t *testing.T) {
	assert.Equal(t, true, isRepeatedDigitsNumberCharSlice(121212))
	assert.Equal(t, false, isRepeatedDigitsNumberCharSlice(12121))
	assert.Equal(t, true, isRepeatedDigitsNumberCharSlice(11))
	assert.Equal(t, true, isRepeatedDigitsNumberCharSlice(33333))
	assert.Equal(t, true, isRepeatedDigitsNumberCharSlice(100100100))
}

// essentially the same as the above (usually around 105 ms)
func TestRepeatedDigitsSliceCopySubsliceFunc(t *testing.T) {
	assert.Equal(t, true, isRepeatedDigitsNumberCharSliceCopySubslice(121212))
	assert.Equal(t, false, isRepeatedDigitsNumberCharSliceCopySubslice(12121))
	assert.Equal(t, true, isRepeatedDigitsNumberCharSliceCopySubslice(11))
	assert.Equal(t, true, isRepeatedDigitsNumberCharSliceCopySubslice(33333))
	assert.Equal(t, true, isRepeatedDigitsNumberCharSliceCopySubslice(100100100))
}

// marginally faster than the others, hitting around 85 ms consistently
func TestRepeatedDigitsNumSliceFunc(t *testing.T) {
	assert.Equal(t, true, isRepeatedDigitsNumberNumSlice(121212))
	assert.Equal(t, false, isRepeatedDigitsNumberNumSlice(12121))
	assert.Equal(t, true, isRepeatedDigitsNumberNumSlice(11))
	assert.Equal(t, true, isRepeatedDigitsNumberNumSlice(33333))
	assert.Equal(t, true, isRepeatedDigitsNumberNumSlice(100100100))
}

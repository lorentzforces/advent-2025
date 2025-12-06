package day_02

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func PartOne(input string) (uint, error) {
	rangePairs, err := parseRangePairs(strings.TrimSpace(input))
	if err != nil { return 0, err }

	runningTotal := uint(0)
	for _, rangePair := range rangePairs {
		invalidIds := getDoubledValuesInRange(rangePair.left, rangePair.right)
		for _, id := range invalidIds {
			runningTotal += id
		}
	}

	return runningTotal, nil

}

type rangePair struct {
	left uint
	right uint
}

func parseRangePairs(s string) ([]rangePair, error) {
	rangeStrings := strings.Split(s, ",")
	results := make([]rangePair, 0, len(rangeStrings))

	for _, rangeString := range rangeStrings {
		valStrings := strings.Split(rangeString, "-")
		if len(valStrings) != 2 {
			return nil, fmt.Errorf("Could not parse value strings from range \"%s\"", rangeString)
		}

		left, err := strconv.Atoi(valStrings[0])
		if err != nil {
			return nil, fmt.Errorf(
				"Could not parse value strings from range \"%s\": %w",
				rangeString, err,
			)
		}
		right, err := strconv.Atoi(valStrings[1])
		if err != nil {
			return nil, fmt.Errorf(
				"Could not parse value strings from range \"%s\": %w",
				rangeString, err,
			)
		}
		results = append(results, rangePair{ left: uint(left), right: uint(right) })
	}
	return results, nil
}

func getDoubledValuesInRange(lower, higher uint) []uint {
	lowerDoubledValue := nextDoubledValueAtOrAbove(lower)
	higherDoubledValue := lastDoubledValueAtOrBelow(higher)

	if lowerDoubledValue > higherDoubledValue {
		return []uint{}
	}

	lowerValue := lowerDoubledValue / tenToPower(numDigits(lowerDoubledValue) / 2)
	higherValue := higherDoubledValue / tenToPower(numDigits(higherDoubledValue) / 2)

	results := make([]uint, 0, higherValue - lowerValue + 1)
	for i := lowerValue; i <= higherValue; i++ {
		results = append(results, doubleWrittenDigits(i))
	}

	return results
}

func lastDoubledValueAtOrBelow(n uint) uint {
	digitCount := numDigits(n)
	if digitCount % 2 == 1 {
		return getAllNinesForDigits(digitCount - 1)
	}

	if isDoubled(n) {
		return n
	}

	truncatedValue := n / tenToPower(digitCount / 2)
	doubled := doubleWrittenDigits(truncatedValue)
	if doubled < n {
		return doubled
	}
	return doubleWrittenDigits(truncatedValue - 1)
}

func isDoubled(n uint) bool {
	runes := []rune(fmt.Sprint(n))
	if len(runes) % 2 != 0 {
		return false
	}

	midpoint := len(runes) / 2
	return string(runes[0:midpoint]) == string(runes[midpoint:])
}


func nextDoubledValueAtOrAbove(n uint) uint {
	digitCount := numDigits(n)
	if digitCount % 2 == 1 {
		return getOneZeroDoubled(digitCount + 1)
	}

	if isDoubled(n) {
		return n
	}

	truncatedValue := n / tenToPower(digitCount / 2)
	doubled := doubleWrittenDigits(truncatedValue)
	if doubled > n {
		return doubled
	}
	return doubleWrittenDigits(truncatedValue + 1)
}

func getAllNinesForDigits(digitCount uint) uint {
	if digitCount < 2 {
		return 0
	}
	val := uint(9)
	for range (digitCount - 1) {
		val = val * 10 + 9
	}

	return val
}

func getOneZeroDoubled(digitCount uint) uint {
	if digitCount % 2 != 0 {
		panic(fmt.Sprintf(
			"Called digitCount with value %d which is not an even number",
			digitCount,
		))
	}

	baseValue := tenToPower(digitCount / 2 - 1)
	return doubleWrittenDigits(baseValue)
}

func numDigits(n uint) uint {
	if n == 0 { return 0 }
	return uint(math.Log10(float64(n))) + 1
}

func doubleWrittenDigits(n uint) uint {
	digitCount := numDigits(n)
	return n * tenToPower(digitCount) + n
}

func tenToPower(n uint) uint {
	if n == 0 {
		return 1
	}

	val := uint(10)
	for range (n - 1) {
		val *= 10
	}
	return val
}

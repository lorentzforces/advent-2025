package day_02

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func PartOne(input string) (uint, error) {
	rangePairs, err := parseRangePairs(strings.TrimSuffix(input, "\n"))
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

// I overengineered the hell out of part 1; for this one I'm going full grug naive iteration mode
// (it turns out the ranges aren't really that big)
func PartTwo(input string) (uint, error) {
	rangePairs, err := parseRangePairs(strings.TrimSuffix(input, "\n"))
	if err != nil { return 0, err }

	totalItemCount := uint(0)
	runningTotal := uint(0)
	for _, rangePair := range rangePairs {
		totalItemCount += rangePair.right - rangePair.left
		for i := rangePair.left; i <= rangePair.right; i++ {
			if isRepeatedDigitsNumberQuant(i) {
				runningTotal += i
			}
		}
	}

	fmt.Fprintf(os.Stderr, "Considered %d total items over all ranges\n", totalItemCount)

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

// quantitative method for determining if a number is a repeated digit pattern
func isRepeatedDigitsNumberQuant(n uint) bool {
	digitCount := int(numDigits(n))

	digitsMask := uint(1)
	for i := 1; i <= digitCount / 2; i++ {
		digitsMask *= 10
		// only evenly-divisible counts are valid since we're looking for fully-repeated elements
		if digitCount % i != 0 {
			continue
		}

		workingNumber := n
		remainder := workingNumber % digitsMask
		foundOnlyRepeats := true
		for workingNumber > 0 {
			newRemainder := workingNumber % digitsMask
			if newRemainder != remainder {
				foundOnlyRepeats = false
				break
			}
			workingNumber = workingNumber / digitsMask
		}
		if foundOnlyRepeats {
			return true
		}
	}

	return false
}

// string-based method for determining if a number is a repeated digit pattern
func isRepeatedDigitsNumber(n uint) bool {
	stringVersion := fmt.Sprint(n)
	// this is safe because we know this is all 1 byte ascii digit characters
	for i := 1; i <= len(stringVersion) / 2; i++ {
		// only evenly-divisible counts are valid since we're looking for fully-repeated elements
		if len(stringVersion) % i != 0 {
			continue
		}
		count := strings.Count(stringVersion, stringVersion[0:i])
		if count * i == len(stringVersion) {
			return true
		}
	}
	return false
}

// slice-walking method for determing if a number is a repeated digit pattern
func isRepeatedDigitsNumberCharSlice(n uint) bool {
	numberChars := []rune(fmt.Sprint(n))

	for i := 1; i <= len(numberChars) / 2; i++ {
		// only evenly-divisible counts are valid since we're looking for fully-repeated elements
		if len(numberChars) % i != 0 {
			continue
		}
		isRepeatedSequence := true
		for j := i; j < len(numberChars); j++ {
			if numberChars[j] != numberChars[j % i] {
				isRepeatedSequence = false
				break
			}
		}
		if isRepeatedSequence {
			return true
		}
	}

	return false
}

// slice-walking method for determing if a number is a repeated digit pattern
// this version copied the subslice we're checking against so we don't dance around offsets
// in a single slice
func isRepeatedDigitsNumberCharSliceCopySubslice(n uint) bool {
	numberChars := []rune(fmt.Sprint(n))

	for i := 1; i <= len(numberChars) / 2; i++ {
		// only evenly-divisible counts are valid since we're looking for fully-repeated elements
		if len(numberChars) % i != 0 {
			continue
		}
		subsequence := numberChars[:i]
		isRepeatedSequence := true
		for j := i; j < len(numberChars); j++ {
			if numberChars[j] != subsequence[j % i] {
				isRepeatedSequence = false
				break
			}
		}
		if isRepeatedSequence {
			return true
		}
	}

	return false
}

// slice-walking, but make it numeric instead of doing a string conversion
func isRepeatedDigitsNumberNumSlice(n uint) bool {
	digitCount := int(numDigits(n))
	digits := make([]int, digitCount)
	workingNumber := n
	for i := range digitCount {
		digits[digitCount - 1 - i] = int(workingNumber % 10)
		workingNumber /= 10
	}

	for i := 1; i <= digitCount / 2; i++ {
		// only evenly-divisible counts are valid since we're looking for fully-repeated elements
		if digitCount % i != 0 {
			continue
		}
		isRepeatedSequence := true
		for j := i; j < digitCount; j++ {
			if digits[j] != digits[j % i] {
				isRepeatedSequence = false
				break
			}
		}
		if isRepeatedSequence {
			return true
		}
	}

	return false
}

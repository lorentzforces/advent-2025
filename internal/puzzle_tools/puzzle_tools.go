package puzzle_tools

import (
	"strings"
	"testing"
)

func AsLines(s string) []string {
	lines := strings.Split(s, "\n")

	// trim trailing blank line (expected)
	if lines[len(lines) - 1] == "" {
		lines = lines[0:len(lines) - 1]
	}
	return lines
}

func AsLinesSplitOnBlanks(s string) [][]string {
	lines := AsLines(s)

	splits := make([][]string, 0, 1)
	start := 0
	for i, line := range lines {
		if line == "" {
			splits = append(splits, lines[start:i])
			start = i + 1
		}
	}

	splits = append(splits, lines[start:])
	return splits
}

func BailIfFailed(t *testing.T) {
	if t.Failed() { t.FailNow() }
}

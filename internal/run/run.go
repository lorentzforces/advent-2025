package run

import (
	"errors"
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
)

const (
	optionInputPath = "input-path"
	optionVerbose = "verbose"
)

func CreateRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "advent2025 [flags] [day [part]]",
		Long: "Run advent of code 2025 exercises.\n" +
			"If provided, runs the given day and part (by number). Will run all parts for a\n" +
			"day if part is not provided, and will run all puzzles if no arguments are passed.",
		SilenceUsage: true,
		DisableFlagsInUseLine: true,
		SilenceErrors: true,
		RunE: runExercises,
	}

	rootCmd.Flags().SortFlags = false
	rootCmd.InitDefaultHelpFlag()
	rootCmd.Flags().StringP(
		optionInputPath,
		"i",
		"",
		"path to a directory where puzzle input files can be found",
	)
	rootCmd.Flags().BoolP(
		optionVerbose,
		"v",
		false,
		"Show verbose output for solution processes.",
	)

	return rootCmd
}

type PuzzleRunFunc func(string) (any, error)

func FailOut(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func FailOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

type puzzleResult struct {
	day int
	part int
	output string
	err error
	duration time.Duration
}

func (pr puzzleResult) String() string {
	return fmt.Sprintf(
		"Day %02d, Part %02d output: %s  [%s]\n",
		pr.day, pr.part, pr.output, pr.duration,
	)
}

func (pr puzzleResult) PrintErr() string {
	if pr.err == nil {
		return "No error!"
	}
	return fmt.Sprintf("ERROR: %s", pr.err.Error())
}

func runExercises(cmd *cobra.Command, args []string) error {
	inputPath, err := cmd.Flags().GetString(optionInputPath)
	if err != nil { return err }

	if len(inputPath) > 0 {
		dirInfo, err := os.Stat(inputPath)
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("Provided input path does not exist.")
		}
		if err != nil { return err }
		if !dirInfo.IsDir() {
			return fmt.Errorf("Provided input path is not a directory.")
		}
	}

	if len(args) > 2 {
		return fmt.Errorf("Too many args, expect a day and part and no more.")
	}

	day := 0
	part := 0
	if len(args) > 0 {
		day, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid day value: \"%s\"", args[0])
		}
	}
	if len(args) > 1 {
		part, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("Invalid part value: \"\"", args[1])
		}
	}

	puzzlesToRun := make([]PuzzleData, 0, 2)

	if day == 0 && part == 0 {
		puzzlesToRun = listOfPuzzles
	} else {
		for _, puzzle := range listOfPuzzles {
			if puzzle.Day == day && (part == 0 || part == puzzle.Part) {
				puzzlesToRun = append(puzzlesToRun, puzzle)
			}
		}
	}

	if len(puzzlesToRun) == 0 {
		FailOut("No puzzles were found for specified day/part")
	}

	inputs, err := loadInputs(inputPath, puzzlesToRun)

	results := make([]puzzleResult, 0, len(puzzlesToRun))
	for _, puzzle := range puzzlesToRun {
		result := puzzleResult{}
		result.day = puzzle.Day
		result.part = puzzle.Part


		start := time.Now()
		output, err := puzzle.Fn(inputs[puzzle.InputFile])
		result.duration= time.Since(start)
		result.err = err
		result.output = fmt.Sprint(output)
	}

	slices.SortFunc(results, func(a, b puzzleResult) int {
		if a.day == b.day {
			return a.part - b.part
		}
		return a.day - b.day
	})

	for _, result := range results {
		fmt.Print(result)
		if result.err != nil {
			fmt.Printf("  %s\n", result.PrintErr())
		}
	}

	return fmt.Errorf("Not yet implemented")
}

func loadInputs(inputPath string, puzzles []PuzzleData) (map[string]string, error) {
	inputFiles, err := os.ReadDir(inputPath)
	if err != nil { return nil, err }

	inputs := make(map[string]string, len(puzzles) / 2 + 1)
	for _, puzzle := range puzzles {
		inputs[puzzle.InputFile] = ""
	}

	for _, foundFile := range inputFiles {
		if foundFile.IsDir() { continue }
		if _, loadFile := inputs[foundFile.Name()]; loadFile {
			input, err := getFileContents(path.Join(inputPath, foundFile.Name()))
			if err != nil { return nil, err }
			inputs[foundFile.Name()] = input
		}

	}

	return inputs, nil
}

func getFileContents(path string) (string, error) {
	fileBuf, err := os.ReadFile(path)
	if err != nil { return "", err }
	return string(fileBuf), nil
}

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

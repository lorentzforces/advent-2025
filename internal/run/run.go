package run

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

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

type PuzzleData struct {
	Day int
	Part int
	InputFile string
	Fn PuzzleRunFunc
}

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

func runExercises(cmd *cobra.Command, args []string) error {
	inputPath, err := cmd.Flags().GetString(optionInputPath)
	if err != nil { return err }

	dirInfo, err := os.Stat(inputPath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("Provided input path does not exist.")
	}
	if err != nil { return err }
	if !dirInfo.IsDir() {
		return fmt.Errorf("Provided input path is not a directory.")
	}

	// TODO: add puzzle data, then generate list of puzzles to run based on args,
	// then get file data, then run the puzzles

	return fmt.Errorf("Not yet implemented")
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

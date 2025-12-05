package main

import "github.com/lorentzforces/advent-2024/internal/run"

func main() {
	err := run.CreateRootCmd().Execute()
	run.FailOnErr(err)
}

# advent-2025

_**Advent of Code 2025**_

Casual workthrough of puzzles from Advent of Code 2025. Using Golang again this year, and again will be attempting to do as much implementation as possible with the standard library. This year I'll be using a couple more libraries than previously for the CLI plumbing, but solutions will be basically entirely only stdlib.

## Building the project

### Build requirements:

- a Golang installation (built & tested on go v1.25)
- an internet connection to download dependencies (only necessary if dependencies have changed or this is the first build)
- a `make` installation. This project is built with GNU make v4 or higher; full compatibility with other versions of make (such as that shipped by Apple) is not guaranteed, but it _should_ be broadly compatible.

To build the project, simply run `make build` in the project's root directory to build the output executable.

> _Note: running with `make` is not strictly necessary. Reference the provided `Makefile` for typical development commands._

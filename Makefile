.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
.SILENT:

# go builds are fast enough that we can just build on demand instead of trying to do any fancy
# change detection
build: clean advent2025
.PHONY: build

advent2025:
	go build ./cmd/advent2025

clean:
	rm -f ./advent2025
.PHONY: clean

check:
	go test ./...
.PHONY: check

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go test ./...                    # run all tests
go test -run TestName            # run a specific test
go run main.go -term 27          # run with a seed value
go build -o collatz              # build binary
```

## Architecture

Single-package Go CLI (`package main`) with two functions:

- `process(term uint, results []uint) []uint` — recursively builds the sequence until it reaches 1
- `calculateNewTerm(term uint) uint` — applies the Collatz rule: odd → `(n * 3) + 1`, even → `n / 2`

## Known issues

- `calculateNewTerm`: `term * 3 + 1` silently overflows for large `uint` values
- `process` is recursive — extreme seeds risk a stack overflow

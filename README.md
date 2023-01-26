# Numerals

A library for working with and converting different numerals.

## CLI

A CLI is available to format and parse numbers.

Formatting a number using roman numerals:
```bash
$ go run github.com/gnirb/numerals/cmd/numerals@latest -system=roman 2023
MMXXIII
```

Parsing roman numerals:
```bash
$ go run github.com/gnirb/numerals/cmd/numerals@latest -system=roman -parse MMXXIII
2023
```
# Numerals

A library for working with and converting different numerals.

Supported numerals:
* Roman (https://en.wikipedia.org/wiki/Roman_numerals)
* Egyptian (https://en.wikipedia.org/wiki/Egyptian_numerals)
* Etruscan (https://en.wikipedia.org/wiki/Etruscan_numerals)
* Aegan (https://en.wikipedia.org/wiki/Aegean_numerals)

## CLI

A CLI is available to format and parse numbers.

Formatting a number using roman numerals:
```bash
$ go run github.com/gnirb/numerals/cmd@latest -system=roman 2023
MMXXIII
```

Parsing roman numerals:
```bash
$ go run github.com/gnirb/numerals/cmd@latest -system=roman -parse MMXXIII
2023
```
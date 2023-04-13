# Numerals

A library for working with and converting different numeral systems in Unicode.

Supported numerals:
* Roman (https://en.wikipedia.org/wiki/Roman_numerals)
* Egyptian (https://en.wikipedia.org/wiki/Egyptian_numerals)
* Etruscan (https://en.wikipedia.org/wiki/Etruscan_numerals)
* Aegan (https://en.wikipedia.org/wiki/Aegean_numerals)
* Greek (https://en.wikipedia.org/wiki/Greek_numerals)

## CLI

A CLI is available to format and parse numbers.

Formatting a number using a numeral system:
```bash
$ go run github.com/gnirb/numerals/cmd@latest -system=roman 2023
MMXXIII

$ go run github.com/gnirb/numerals/cmd@latest -system=egyptian 2023
𓆼𓆼𓎆𓎆𓏤𓏤𓏤
```

Parsing a number using a numeral system:
```bash
$ go run github.com/gnirb/numerals/cmd@latest -system=roman -parse MMXXIII
2023

$ go run github.com/gnirb/numerals/cmd@latest -system=egyptian -parse 𓆼𓆼𓎆𓎆𓏤𓏤𓏤
2023
```
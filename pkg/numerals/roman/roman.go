package roman

import (
	"github.com/gnirb/numerals/pkg/numerals"
)

const (
	Zero            numerals.Numeral = 'N'
	One             numerals.Numeral = 'I'
	Five            numerals.Numeral = 'V'
	Ten             numerals.Numeral = 'X'
	Fifty           numerals.Numeral = 'L'
	Hundred         numerals.Numeral = 'C'
	FiveHundred     numerals.Numeral = 'D'
	Thousand        numerals.Numeral = 'M'
	FiveThousand    numerals.Numeral = 'ↁ'
	TenThousand     numerals.Numeral = 'ↂ'
	FiftyThousand   numerals.Numeral = 'ↇ'
	HundredThousand numerals.Numeral = 'ↈ'
)

var (
	values = []numerals.NumeralValue{
		{
			Numeral: []numerals.Numeral{HundredThousand},
			Value:   100000,
		},
		{
			Numeral: []numerals.Numeral{FiftyThousand},
			Value:   50000,
		},
		{
			Numeral: []numerals.Numeral{TenThousand},
			Value:   10000,
		},
		{
			Numeral: []numerals.Numeral{FiveThousand},
			Value:   5000,
		},
		{
			Numeral: []numerals.Numeral{Thousand},
			Value:   1000,
		},
		{
			Numeral: []numerals.Numeral{Hundred, Thousand},
			Value:   900,
		},
		{
			Numeral: []numerals.Numeral{FiveHundred},
			Value:   500,
		},
		{
			Numeral: []numerals.Numeral{Hundred, FiveHundred},
			Value:   400,
		},
		{
			Numeral: []numerals.Numeral{Hundred},
			Value:   100,
		},
		{
			Numeral: []numerals.Numeral{Ten, Hundred},
			Value:   90,
		},
		{
			Numeral: []numerals.Numeral{Fifty},
			Value:   50,
		},
		{
			Numeral: []numerals.Numeral{Ten, Fifty},
			Value:   40,
		},
		{
			Numeral: []numerals.Numeral{Ten},
			Value:   10,
		},
		{
			Numeral: []numerals.Numeral{One, Ten},
			Value:   9,
		},
		{
			Numeral: []numerals.Numeral{Five},
			Value:   5,
		},
		{
			Numeral: []numerals.Numeral{One, Five},
			Value:   4,
		},
		{
			Numeral: []numerals.Numeral{One},
			Value:   1,
		},
		{
			Numeral: []numerals.Numeral{Zero}, // Nulla, or None
			Value:   0,
		},
	}
	numeralMap = map[numerals.Numeral]*numerals.NumeralValue{}
)

func init() {
	for i, v := range values {
		// skip values with subtractive notation (more than one numeral)
		if len(v.Numeral) == 1 {
			numeralMap[v.Numeral[0]] = &values[i]
		}
	}
}

type romanConverter = numerals.SignValueConverter

func NewRomanConverter() *romanConverter {
	return &romanConverter{
		Values:     values,
		NumeralMap: numeralMap,
		Zero:       Zero,
	}
}

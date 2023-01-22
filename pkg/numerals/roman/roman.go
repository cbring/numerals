package roman

import (
	"strings"
)

type Numeral = rune
type Number = string

const (
	Zero            Numeral = 'N'
	One             Numeral = 'I'
	Five            Numeral = 'V'
	Ten             Numeral = 'X'
	Fifty           Numeral = 'L'
	Hundred         Numeral = 'C'
	FiveHundred     Numeral = 'D'
	Thousand        Numeral = 'M'
	FiveThousand    Numeral = 'ↁ'
	TenThousand     Numeral = 'ↂ'
	FiftyThousand   Numeral = 'ↇ'
	HundredThousand Numeral = 'ↈ'
)

type NumeralValue struct {
	numeral []Numeral
	value   int
}

var values = []NumeralValue{
	{
		numeral: []Numeral{HundredThousand},
		value:   100000,
	},
	{
		numeral: []Numeral{FiftyThousand},
		value:   50000,
	},
	{
		numeral: []Numeral{TenThousand},
		value:   10000,
	},
	{
		numeral: []Numeral{FiveThousand},
		value:   5000,
	},
	{
		numeral: []Numeral{Thousand},
		value:   1000,
	},
	{
		numeral: []Numeral{Hundred, Thousand},
		value:   900,
	},
	{
		numeral: []Numeral{FiveHundred},
		value:   500,
	},
	{
		numeral: []Numeral{Hundred, FiveHundred},
		value:   400,
	},
	{
		numeral: []Numeral{Hundred},
		value:   100,
	},
	{
		numeral: []Numeral{Ten, Hundred},
		value:   90,
	},
	{
		numeral: []Numeral{Fifty},
		value:   50,
	},
	{
		numeral: []Numeral{Ten, Fifty},
		value:   40,
	},
	{
		numeral: []Numeral{Ten},
		value:   10,
	},
	{
		numeral: []Numeral{One, Ten},
		value:   9,
	},
	{
		numeral: []Numeral{Five},
		value:   5,
	},
	{
		numeral: []Numeral{One, Five},
		value:   4,
	},
	{
		numeral: []Numeral{One},
		value:   1,
	},
	{
		numeral: []Numeral{Zero}, // Nulla, or None
		value:   0,
	},
}

// Rtoi converts roman numeral numbers to arabic numbers (integer)
func Rtoi(num Number) int {
	sum := 0

	for _, v := range values {
		for strings.HasPrefix(num, string(v.numeral)) {
			sum += v.value
			num = num[len(v.numeral):]
		}
	}

	return sum
}

func Itor(num int) Number {
	nb := strings.Builder{}

	if num == 0 {
		nb.WriteRune(Zero)
	} else {
		for _, v := range values {
			if num == 0 {
				break
			}
			if num >= v.value {
				cnt := num / v.value
				for i := 0; i < cnt; i++ {
					_, _ = nb.WriteString(string(v.numeral))
				}
				num -= cnt * v.value
			}
		}
	}

	return nb.String()
}

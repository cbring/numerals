package roman

import (
	"strings"

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

type NumeralValue struct {
	numeral []numerals.Numeral
	value   int
}

var (
	values = []NumeralValue{
		{
			numeral: []numerals.Numeral{HundredThousand},
			value:   100000,
		},
		{
			numeral: []numerals.Numeral{FiftyThousand},
			value:   50000,
		},
		{
			numeral: []numerals.Numeral{TenThousand},
			value:   10000,
		},
		{
			numeral: []numerals.Numeral{FiveThousand},
			value:   5000,
		},
		{
			numeral: []numerals.Numeral{Thousand},
			value:   1000,
		},
		{
			numeral: []numerals.Numeral{Hundred, Thousand},
			value:   900,
		},
		{
			numeral: []numerals.Numeral{FiveHundred},
			value:   500,
		},
		{
			numeral: []numerals.Numeral{Hundred, FiveHundred},
			value:   400,
		},
		{
			numeral: []numerals.Numeral{Hundred},
			value:   100,
		},
		{
			numeral: []numerals.Numeral{Ten, Hundred},
			value:   90,
		},
		{
			numeral: []numerals.Numeral{Fifty},
			value:   50,
		},
		{
			numeral: []numerals.Numeral{Ten, Fifty},
			value:   40,
		},
		{
			numeral: []numerals.Numeral{Ten},
			value:   10,
		},
		{
			numeral: []numerals.Numeral{One, Ten},
			value:   9,
		},
		{
			numeral: []numerals.Numeral{Five},
			value:   5,
		},
		{
			numeral: []numerals.Numeral{One, Five},
			value:   4,
		},
		{
			numeral: []numerals.Numeral{One},
			value:   1,
		},
		{
			numeral: []numerals.Numeral{Zero}, // Nulla, or None
			value:   0,
		},
	}
	numeralMap = map[numerals.Numeral]*NumeralValue{}
)

func init() {
	for i, v := range values {
		// skip values with subtractive notation (more than one numeral)
		if len(v.numeral) == 1 {
			numeralMap[v.numeral[0]] = &values[i]
		}
	}
}

type romanConverter struct{}

func NewRomanConverter() *romanConverter {
	return &romanConverter{}
}

// Parse converts roman numeral numbers to arabic numbers (integer)
func (r romanConverter) Parse(number numerals.Number) int {
	sum := 0

	var prev numerals.Numeral = 'N'
	var val int
	for _, r := range number {
		var num = numerals.Numeral(r)
		if num != prev {
			if numeralMap[num].value < numeralMap[prev].value {
				sum += val
			} else {
				sum -= val
			}

			val = 0
			prev = num
		}
		val += numeralMap[num].value
	}
	sum += val

	return sum
}

func (r romanConverter) Format(num int) numerals.Number {
	nb := strings.Builder{}

	if num == 0 {
		nb.WriteRune(rune(Zero))
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

	return numerals.Number(nb.String())
}

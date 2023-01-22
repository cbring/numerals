package ronum

import (
	"bytes"
	"strings"
)

type Numeral = byte
type Number = string

const (
	N Numeral = 'N'
	I Numeral = 'I'
	V Numeral = 'V'
	X Numeral = 'X'
	L Numeral = 'L'
	C Numeral = 'C'
	D Numeral = 'D'
	M Numeral = 'M'
)

type NumeralValue struct {
	numeral []Numeral
	value   int
}

var values = []NumeralValue{
	{
		numeral: []Numeral{M},
		value:   1000,
	},
	{
		numeral: []Numeral{C, M},
		value:   900,
	},
	{
		numeral: []Numeral{D},
		value:   500,
	},
	{
		numeral: []Numeral{C, D},
		value:   400,
	},
	{
		numeral: []Numeral{C},
		value:   100,
	},
	{
		numeral: []Numeral{X, C},
		value:   90,
	},
	{
		numeral: []Numeral{L},
		value:   50,
	},
	{
		numeral: []Numeral{X, L},
		value:   40,
	},
	{
		numeral: []Numeral{X},
		value:   10,
	},
	{
		numeral: []Numeral{I, X},
		value:   9,
	},
	{
		numeral: []Numeral{V},
		value:   5,
	},
	{
		numeral: []Numeral{I, V},
		value:   4,
	},
	{
		numeral: []Numeral{I},
		value:   1,
	},
	{
		numeral: []Numeral{N}, // Nulla, or None
		value:   0,
	},
}

// Rtoi converts roman numeral numbers to arabic numbers (integer)
func Rtoi(num Number) int {
	sum := 0
	b := []byte(num)

	for _, v := range values {
		for bytes.HasPrefix(b, v.numeral) {
			sum += v.value
			b = b[len(v.numeral):]
		}
	}

	return sum
}

func Itor(num int) Number {
	nb := strings.Builder{}

	if num == 0 {
		nb.WriteByte(N)
	} else {
		for _, v := range values {
			if num == 0 {
				break
			}
			if num >= v.value {
				cnt := num / v.value
				for i := 0; i < cnt; i++ {
					_, _ = nb.Write(v.numeral)
				}
				num -= cnt * v.value
			}
		}
	}

	return nb.String()
}

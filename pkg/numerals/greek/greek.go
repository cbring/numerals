package greek

import (
	"strings"

	"github.com/gnirb/numerals/pkg/numerals"
)

// Symbols are not quite sequential in unicode unfortunately,
// so we have to map them explicitly
const (
	Zero        numerals.Numeral = '𐆊'
	Keraia      numerals.Numeral = 'ʹ'
	LowerKeraia numerals.Numeral = '͵'

	One   numerals.Numeral = 'Α'
	Two   numerals.Numeral = 'Β'
	Three numerals.Numeral = 'Γ'
	Four  numerals.Numeral = 'Δ'
	Five  numerals.Numeral = 'Ε'
	Six   numerals.Numeral = 'Ϝ'
	Seven numerals.Numeral = 'Ζ'
	Eight numerals.Numeral = 'Η'
	Nine  numerals.Numeral = 'Θ'

	Ten     numerals.Numeral = 'Ι'
	Twenty  numerals.Numeral = 'Κ'
	Thirty  numerals.Numeral = 'Λ'
	Fourty  numerals.Numeral = 'Μ'
	Fifty   numerals.Numeral = 'Ν'
	Sixty   numerals.Numeral = 'Ξ'
	Seventy numerals.Numeral = 'Ο'
	Eighty  numerals.Numeral = 'Π'
	Ninety  numerals.Numeral = 'Ϙ'

	Hundred      numerals.Numeral = 'Ρ'
	TwoHundred   numerals.Numeral = 'Σ'
	ThreeHundred numerals.Numeral = 'Τ'
	FourHundred  numerals.Numeral = 'Υ'
	FiveHundred  numerals.Numeral = 'Φ'
	SixHundred   numerals.Numeral = 'Χ'
	SevenHundred numerals.Numeral = 'Ψ'
	EightHundred numerals.Numeral = 'Ω'
	NineHundred  numerals.Numeral = 'Ͳ'
)

var (
	values = []numerals.NumeralValue{
		{
			Numeral: []numerals.Numeral{NineHundred},
			Value:   900,
		},
		{
			Numeral: []numerals.Numeral{EightHundred},
			Value:   800,
		},
		{
			Numeral: []numerals.Numeral{SevenHundred},
			Value:   700,
		},
		{
			Numeral: []numerals.Numeral{SixHundred},
			Value:   600,
		},
		{
			Numeral: []numerals.Numeral{FiveHundred},
			Value:   500,
		},
		{
			Numeral: []numerals.Numeral{FourHundred},
			Value:   400,
		},
		{
			Numeral: []numerals.Numeral{ThreeHundred},
			Value:   300,
		},
		{
			Numeral: []numerals.Numeral{TwoHundred},
			Value:   200,
		},
		{
			Numeral: []numerals.Numeral{Hundred},
			Value:   100,
		},
		{
			Numeral: []numerals.Numeral{Ninety},
			Value:   90,
		},
		{
			Numeral: []numerals.Numeral{Eighty},
			Value:   80,
		},
		{
			Numeral: []numerals.Numeral{Seventy},
			Value:   70,
		},
		{
			Numeral: []numerals.Numeral{Sixty},
			Value:   60,
		},
		{
			Numeral: []numerals.Numeral{Fifty},
			Value:   50,
		},
		{
			Numeral: []numerals.Numeral{Fourty},
			Value:   40,
		},
		{
			Numeral: []numerals.Numeral{Thirty},
			Value:   30,
		},
		{
			Numeral: []numerals.Numeral{Twenty},
			Value:   20,
		},
		{
			Numeral: []numerals.Numeral{Ten},
			Value:   10,
		},
		{
			Numeral: []numerals.Numeral{Nine},
			Value:   9,
		},
		{
			Numeral: []numerals.Numeral{Eight},
			Value:   8,
		},
		{
			Numeral: []numerals.Numeral{Seven},
			Value:   7,
		},
		{
			Numeral: []numerals.Numeral{Six},
			Value:   6,
		},
		{
			Numeral: []numerals.Numeral{Five},
			Value:   5,
		},
		{
			Numeral: []numerals.Numeral{Four},
			Value:   4,
		},
		{
			Numeral: []numerals.Numeral{Three},
			Value:   3,
		},
		{
			Numeral: []numerals.Numeral{Two},
			Value:   2,
		},
		{
			Numeral: []numerals.Numeral{One},
			Value:   1,
		},
	}
	numeralMap = map[numerals.Numeral]*numerals.NumeralValue{}
)

func init() {

	for i, v := range values {
		numeralMap[v.Numeral[0]] = &values[i]
	}
}

type greekConverter struct {
	Values     []numerals.NumeralValue
	NumeralMap map[numerals.Numeral]*numerals.NumeralValue
	Zero       numerals.Numeral
	Indicators []numerals.NumeralValue
}

func (c greekConverter) Parse(number numerals.Number) int {
	const defaultMultiplier = 1
	sum := 0

	var prev = c.Zero
	var val int
	var multiplier = defaultMultiplier

	for _, r := range number {
		var num = numerals.Numeral(r)

		if m, found := c.indicator(num); found {
			multiplier = m
			continue
		}

		if num != prev {

			if prev != 0 && prev != c.Zero {

				if c.NumeralMap[num].Value < c.NumeralMap[prev].Value {
					sum += val
				} else {
					if multiplier != defaultMultiplier {
						multiplier = defaultMultiplier
						sum += val
					} else {
						sum -= val
					}
				}
			}

			val = 0
			prev = num
		}
		val += c.NumeralMap[num].Value * multiplier
	}
	sum += val

	return sum
}

func (c greekConverter) indicator(num numerals.Numeral) (int, bool) {
	for _, m := range c.Indicators {
		if m.Numeral[0] == num {
			return m.Value, true
		}
	}
	return 1, false
}

// Format formats an integer to numeral numbers
func (c greekConverter) Format(num int) numerals.Number {
	numberBuilder := strings.Builder{}

	if num == 0 {
		// handle zero if supported by the numerical system
		if last := c.Values[len(c.Values)-1]; last.Value == 0 {
			numberBuilder.WriteRune(rune(c.Zero))
		}
	} else {
		var lastMul = 0
		for _, m := range c.Indicators {
			mul := m.Value
			ind := m.Numeral[0]
			for _, v := range c.Values {
				if num == 0 {
					break
				}

				val := v.Value * mul
				if num >= val {
					if lastMul == 0 && mul > 1 {
						lastMul = mul
						_, _ = numberBuilder.WriteRune(rune(ind))
					}
					cnt := num / val
					for i := 0; i < cnt; i++ {
						_, _ = numberBuilder.WriteString(string(v.Numeral))
					}
					num -= cnt * val

					if num == 0 && mul == 1 {
						_, _ = numberBuilder.WriteRune(rune(ind))
					}

					continue
				}
			}

		}
	}

	return numerals.Number(numberBuilder.String())
}

func NewGreekConverter() *greekConverter {
	return &greekConverter{
		Values:     values,
		NumeralMap: numeralMap,
		Zero:       Zero,
		Indicators: []numerals.NumeralValue{
			{
				Numeral: []numerals.Numeral{LowerKeraia},
				Value:   1000,
			},
			{
				Numeral: []numerals.Numeral{Keraia},
				Value:   1,
			},
		},
	}
}

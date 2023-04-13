package numerals

import "strings"

type Numeral rune
type Number string

type Converter interface {
	Format(int) Number
	Parse(Number) int
}

type SignValueConverter struct {
	Values     []NumeralValue
	NumeralMap map[Numeral]*NumeralValue
	Zero       Numeral
}

type NumeralValue struct {
	Numeral []Numeral
	Value   int
}

// Parse converts sign value numeral numbers to arabic numbers (integer)
func (c SignValueConverter) Parse(number Number) int {
	sum := 0

	var prev Numeral

	if last := c.Values[len(c.Values)-1]; last.Value == 0 {
		prev = last.Numeral[0]
	}

	var val int
	for _, r := range number {
		var num = Numeral(r)
		if num != prev {
			if prev != 0 {
				if c.NumeralMap[num].Value < c.NumeralMap[prev].Value {
					sum += val
				} else {
					sum -= val
				}
			}

			val = 0
			prev = num
		}
		val += c.NumeralMap[num].Value
	}
	sum += val

	return sum
}

func (c SignValueConverter) Format(num int) Number {
	numberBuilder := strings.Builder{}

	if num == 0 {
		// handle zero if supported by the numerical system
		if last := c.Values[len(c.Values)-1]; last.Value == 0 {
			numberBuilder.WriteRune(rune(c.Zero))
		}
	} else {
		for _, v := range c.Values {
			if num == 0 {
				break
			}
			if num >= v.Value {
				cnt := num / v.Value
				for i := 0; i < cnt; i++ {
					_, _ = numberBuilder.WriteString(string(v.Numeral))
				}
				num -= cnt * v.Value
			}
		}
	}

	return Number(numberBuilder.String())
}

func Repeat(n Numeral, times int) []Numeral {
	var numerals []Numeral
	for i := 0; i < times; i++ {
		numerals = append(numerals, n)
	}
	return numerals
}

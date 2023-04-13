package aegan

import "github.com/gnirb/numerals/pkg/numerals"

const (
	One         numerals.Numeral = '𐄇'
	Ten         numerals.Numeral = '𐄐'
	Hundred     numerals.Numeral = '𐄙'
	Thousand    numerals.Numeral = '𐄢'
	TenThousand numerals.Numeral = '𐄫'
)

var (
	values     []numerals.NumeralValue
	numeralMap = map[numerals.Numeral]*numerals.NumeralValue{}
)

func init() {
	for r, i := TenThousand+8, 9; r >= TenThousand; r-- {
		values = append(values, numerals.NumeralValue{
			Numeral: []numerals.Numeral{r},
			Value:   10000 * i,
		})
		i--
	}

	for r, i := Thousand+8, 9; r >= Thousand; r-- {
		values = append(values, numerals.NumeralValue{
			Numeral: []numerals.Numeral{r},
			Value:   1000 * i,
		})
		i--
	}

	for r, i := Hundred+8, 9; r >= Hundred; r-- {
		values = append(values, numerals.NumeralValue{
			Numeral: []numerals.Numeral{r},
			Value:   100 * i,
		})
		i--
	}

	for r, i := Ten+8, 9; r >= Ten; r-- {
		values = append(values, numerals.NumeralValue{
			Numeral: []numerals.Numeral{r},
			Value:   10 * i,
		})
		i--
	}

	for r, i := One+8, 9; r >= One; r-- {
		values = append(values, numerals.NumeralValue{
			Numeral: []numerals.Numeral{r},
			Value:   i,
		})
		i--
	}

	for i, v := range values {
		numeralMap[v.Numeral[0]] = &values[i]
	}
}

type aeganConverter = numerals.SignValueConverter

func NewAeganConverter() *aeganConverter {
	return &aeganConverter{
		Values:     values,
		NumeralMap: numeralMap,
	}
}

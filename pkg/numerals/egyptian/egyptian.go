package egyptian

import "github.com/gnirb/numerals/pkg/numerals"

const (
	One             numerals.Numeral = '𓏤'
	Ten             numerals.Numeral = '𓎆'
	Hundred         numerals.Numeral = '𓍢'
	Thousand        numerals.Numeral = '𓆼'
	TenThousand     numerals.Numeral = '𓂭'
	HundredThousand numerals.Numeral = '𓆐'
	Million         numerals.Numeral = '𓁨'
)

var (
	values = []numerals.NumeralValue{
		{
			Numeral: []numerals.Numeral{Million},
			Value:   1000000,
		},
		{
			Numeral: []numerals.Numeral{HundredThousand},
			Value:   100000,
		},
		{
			Numeral: []numerals.Numeral{TenThousand},
			Value:   10000,
		},
		{
			Numeral: []numerals.Numeral{Thousand},
			Value:   1000,
		},
		{
			Numeral: []numerals.Numeral{Hundred},
			Value:   100,
		},
		{
			Numeral: []numerals.Numeral{Ten},
			Value:   10,
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

type egyptianConverter = numerals.SignValueConverter

func NewEgyptianConverter() *egyptianConverter {
	return &egyptianConverter{
		Values:     values,
		NumeralMap: numeralMap,
	}
}

package etruscan

import "github.com/gnirb/numerals/pkg/numerals"

const (
	One     numerals.Numeral = '𐌠'
	Five    numerals.Numeral = '𐌡'
	Ten     numerals.Numeral = '𐌢'
	Fifty   numerals.Numeral = '𐌣'
	Hundred numerals.Numeral = '𐌟'
)

var (
	values = []numerals.NumeralValue{
		{
			Numeral: []numerals.Numeral{Hundred},
			Value:   100,
		},
		{
			Numeral: []numerals.Numeral{Fifty},
			Value:   50,
		},
		{
			Numeral: []numerals.Numeral{Ten},
			Value:   10,
		},
		{
			Numeral: []numerals.Numeral{Five},
			Value:   5,
		},
		{
			Numeral: []numerals.Numeral{One},
			Value:   1,
		},
	}
	numeralMap = map[numerals.Numeral]*numerals.NumeralValue{}
)

func init() {
	// Etruscans would often write 17, 18, and 19 as "𐌠𐌠𐌠𐌢𐌢", "𐌠𐌠𐌢𐌢", and "𐌠𐌢𐌢" – that is,
	// "three from twenty", "two from twenty", and "one from twenty", instead of "𐌢𐌡𐌠𐌠", "𐌢𐌡𐌠𐌠𐌠", and "𐌢𐌡𐌠𐌠𐌠𐌠"
	var quirks []numerals.NumeralValue
	for i := 1; i < 5; i++ {
		quirks = append(quirks,
			numerals.NumeralValue{
				Numeral: append([]numerals.Numeral{One, Ten}, numerals.Repeat(Ten, i)...),
				Value:   9 + i*10,
			})
		quirks = append(quirks,
			numerals.NumeralValue{
				Numeral: append([]numerals.Numeral{One, One, Ten}, numerals.Repeat(Ten, i)...),
				Value:   8 + i*10,
			})
		quirks = append(quirks,
			numerals.NumeralValue{
				Numeral: append([]numerals.Numeral{One, One, One, Ten}, numerals.Repeat(Ten, i)...),
				Value:   7 + i*10,
			})
	}

	// insert quirks at the correct positions in values
	for _, v := range quirks {
		for i, vv := range values {
			if v.Value > vv.Value {
				values = append(values[:i+1], values[i:]...)
				values[i] = v
				break
			}
		}
	}

	for i, v := range values {
		// skip values with subtractive notation (more than one numeral)
		if len(v.Numeral) == 1 {
			numeralMap[v.Numeral[0]] = &values[i]
		}
	}
}

type etruscanConverter = numerals.SignValueConverter

func NewEtruscanConverter() *etruscanConverter {
	return &etruscanConverter{
		Values:     values,
		NumeralMap: numeralMap,
	}
}

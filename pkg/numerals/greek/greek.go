package greek

import "github.com/gnirb/numerals/pkg/numerals"

// Symbols are not quite sequential in unicode unfortunately,
// so we have to map them explicitly
const (
	Zero        numerals.Numeral = '𐆊'
	Keraia      numerals.Numeral = 'ʹ'
	LowerKeraia numerals.Numeral = ','

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
		{
			Numeral: []numerals.Numeral{Zero},
			Value:   0,
		},
	}
	numeralMap = map[numerals.Numeral]*numerals.NumeralValue{}
)

func init() {
	for i, v := range values {
		numeralMap[v.Numeral[0]] = &values[i]
	}
}

type greekConverter = numerals.SignValueConverter

func NewGreekConverter() *greekConverter {
	return &greekConverter{
		Values:     values,
		NumeralMap: numeralMap,
		Zero:       Zero,
	}
}

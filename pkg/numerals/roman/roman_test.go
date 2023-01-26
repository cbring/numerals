package roman

import (
	"testing"

	"github.com/gnirb/numerals/pkg/numerals"
)

type tests struct {
	name    string
	integer int
	number  numerals.Number
}

var (
	romanNumerals = []tests{
		{
			"numeral N",
			0,
			"N",
		},
		{
			"numeral I",
			1,
			"I",
		},
		{
			"numeral V",
			5,
			"V",
		},
		{
			"numeral X",
			10,
			"X",
		},
		{
			"numeral L",
			50,
			"L",
		},
		{
			"numeral C",
			100,
			"C",
		},
		{
			"numeral D",
			500,
			"D",
		},
		{
			"numeral M",
			1000,
			"M",
		},
	}
	subtractiveNotation = []tests{
		{
			"subtractive notation IV",
			4,
			"IV",
		},
		{
			"subtractive notation IX",
			9,
			"IX",
		},
		{
			"subtractive notation XL",
			40,
			"XL",
		},
		{
			"subtractive notation XC",
			90,
			"XC",
		},
		{
			"subtractive notation CD",
			400,
			"CD",
		},
		{
			"subtractive notation CM",
			900,
			"CM",
		},
	}
	additiveNotation = []tests{
		// Additive notation
		{
			"additive notation III",
			3,
			"III",
		},
		{
			"additive notation XXX",
			30,
			"XXX",
		},
		{
			"additive notation CCC",
			300,
			"CCC",
		},
		{
			"additive notation MMM",
			3000,
			"MMM",
		},
	}
	clock = []tests{
		// Clock number
		{
			"clock II",
			2,
			"II",
		},
		{
			"clock VI",
			6,
			"VI",
		},
		{
			"clock VII",
			7,
			"VII",
		},
		{
			"clock VIII",
			8,
			"VIII",
		},
		{
			"clock XI",
			11,
			"XI",
		},
		{
			"clock XII",
			12,
			"XII",
		},
	}
	samples = []tests{
		// Samples
		{
			"sample XXXIX",
			39,
			"XXXIX",
		},
		{
			"sample CCXLVI",
			246,
			"CCXLVI",
		},
		{
			"sample DCCLXXXIX",
			789,
			"DCCLXXXIX",
		},
		{
			"sample MMCDXXI",
			2421,
			"MMCDXXI",
		},
		{
			"sample CLX",
			160,
			"CLX",
		},
		{
			"sample CCVII",
			207,
			"CCVII",
		},
		{
			"sample MIX",
			1009,
			"MIX",
		},
		{
			"sample MLXVI",
			1066,
			"MLXVI",
		},
		{
			"sample MMMCMXCIX",
			3999,
			"MMMCMXCIX",
		},
	}
	bigNumbers = []tests{
		// Big numbers
		{
			"bignum ↁ",
			5000,
			"ↁ",
		},
		{
			"bignum ↂ",
			10000,
			"ↂ",
		},
		{
			"bignum ↇ",
			50000,
			"ↇ",
		},
		{
			"bignum ↈ",
			100000,
			"ↈ",
		},
	}
	all = [][]tests{
		romanNumerals,
		additiveNotation,
		subtractiveNotation,
		clock,
		samples,
		bigNumbers,
	}
	converter = NewRomanConverter()
)

func TestParse(t *testing.T) {
	for _, tt := range all {
		for _, ttt := range tt {
			t.Run(ttt.name, func(t *testing.T) {
				if got := converter.Parse(ttt.number); got != ttt.integer {
					t.Errorf("Parse() = %v, want %v", got, ttt.number)
				}
			})
		}
	}
}

func TestFormat(t *testing.T) {
	for _, tt := range all {
		for _, ttt := range tt {
			t.Run(ttt.name, func(t *testing.T) {
				if got := converter.Format(ttt.integer); got != ttt.number {
					t.Errorf("Format() = %v, want %v", got, ttt.number)
				}
			})
		}
	}
}

func TestOtherForms(t *testing.T) {
	var tests = []struct {
		name    string
		integer int
		number  numerals.Number
	}{
		// Additive
		{
			"additive IIII",
			4,
			"IIII",
		},
		{
			"additive XXXX",
			40,
			"XXXX",
		},
		{
			"additive CCCC",
			400,
			"CCCC",
		},
		{
			"additive IIIIII",
			6,
			"IIIIII",
		},
		{
			"additive XXXXXX",
			60,
			"XXXXXX",
		},
		{
			"additive CCCCLXXXX",
			490,
			"CCCCLXXXX",
		},
		// Variants
		{
			"variant MDCCCCX",
			1910,
			"MDCCCCX",
		},
		{
			"variant MDCDIII",
			1903,
			"MDCDIII",
		},
		{
			"variant CDXCIX",
			499,
			"CDXCIX",
		},
		{
			"variant LDVLIV",
			499,
			"LDVLIV",
		},
		{
			"variant XDIX",
			499,
			"XDIX",
		},
		{
			"variant VDIV",
			499,
			"VDIV",
		},
		{
			"variant ID",
			499,
			"ID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := converter.Parse(tt.number); got != tt.integer {
				t.Errorf("Parse() = %v, want %v", got, tt.number)
			}
		})
	}
}

func BenchmarkParse_numerals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range romanNumerals {
			converter.Parse(tt.number)
		}
	}
}

func BenchmarkParse_samples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range samples {
			converter.Parse(tt.number)
		}
	}
}

func BenchmarkParse_all(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range all {
			for _, ttt := range tt {
				converter.Parse(ttt.number)
			}
		}
	}
}

func BenchmarkFormat_numerals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range romanNumerals {
			converter.Format(tt.integer)
		}
	}
}

func BenchmarkFormat_samples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range samples {
			converter.Format(tt.integer)
		}
	}
}

func BenchmarkFormat_all(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range all {
			for _, ttt := range tt {
				converter.Format(ttt.integer)
			}
		}
	}
}

package greek_test

import (
	"testing"

	"github.com/gnirb/numerals/pkg/numerals"
	"github.com/gnirb/numerals/pkg/numerals/greek"
)

type tests struct {
	name    string
	integer int
	number  numerals.Number
}

var (
	greekNumerals = []tests{
		{
			"numeral 1",
			1,
			"Αʹ",
		},
		{
			"numeral 9",
			9,
			"Θʹ",
		},
		{
			"numeral 10",
			10,
			"Ιʹ",
		},
		{
			"numeral 100",
			100,
			"Ρʹ",
		},
		{
			"numeral 900000",
			900000,
			"͵Ͳ",
		},
	}
	samples = []tests{
		// Samples
		{
			"sample 2019",
			2019,
			"͵ΒΙΘʹ",
		},
		{
			"sample 22019",
			22019,
			"͵ΚΒΙΘʹ",
		},
	}
	all = [][]tests{
		greekNumerals,
		samples,
	}
	converter = greek.NewGreekConverter()
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

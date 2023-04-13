package aegan_test

import (
	"testing"

	"github.com/gnirb/numerals/pkg/numerals"
	"github.com/gnirb/numerals/pkg/numerals/aegan"
)

type tests struct {
	name    string
	integer int
	number  numerals.Number
}

var (
	aeganNumerals = []tests{
		{
			"numeral 1",
			1,
			"𐄇",
		},
		{
			"numeral 10",
			10,
			"𐄐",
		},
		{
			"numeral 100",
			100,
			"𐄙",
		},
		{
			"numeral 1000",
			1000,
			"𐄢",
		},
		{
			"numeral 10000",
			10000,
			"𐄫",
		},
		{
			"numeral 100000",
			100000,
			"𐄳𐄫",
		},
	}
	samples = []tests{
		// Samples
		{
			"sample 4622",
			4622,
			"𐄥𐄞𐄑𐄈",
		},
	}
	all = [][]tests{
		aeganNumerals,
		samples,
	}
	converter = aegan.NewAeganConverter()
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

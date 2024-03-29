package etruscan_test

import (
	"testing"

	"github.com/gnirb/numerals/pkg/numerals"
	"github.com/gnirb/numerals/pkg/numerals/etruscan"
)

type tests struct {
	name    string
	integer int
	number  numerals.Number
}

var (
	etruscanNumerals = []tests{
		{
			"numeral zero",
			0,
			"",
		},
		{
			"numeral one",
			1,
			"𐌠",
		},
		{
			"numeral five",
			5,
			"𐌡",
		},
		{
			"numeral ten",
			10,
			"𐌢",
		},
		{
			"numeral fifty",
			50,
			"𐌣",
		},
		{
			"numeral hundred",
			100,
			"𐌟",
		},
	}
	samples = []tests{
		// Samples
		{
			"sample 87",
			87,
			"𐌣𐌠𐌠𐌠𐌢𐌢𐌢𐌢",
		},
		{
			"sample 7",
			7,
			"𐌡𐌠𐌠",
		},
		{
			"sample 17",
			17,
			"𐌠𐌠𐌠𐌢𐌢",
		},
	}
	all = [][]tests{
		etruscanNumerals,
		samples,
	}
	converter = etruscan.NewEtruscanConverter()
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

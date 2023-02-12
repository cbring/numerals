package egyptian

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
	egyptianNumerals = []tests{
		{
			"numeral zero",
			0,
			"",
		},
		{
			"numeral 𓏤",
			1,
			"𓏤",
		},
		{
			"numeral 𓎆",
			10,
			"𓎆",
		},
		{
			"numeral 𓍢",
			100,
			"𓍢",
		},
		{
			"numeral 𓆼",
			1000,
			"𓆼",
		},
		{
			"numeral 𓂭",
			10000,
			"𓂭",
		},
		{
			"numeral 𓆐",
			100000,
			"𓆐",
		},
		{
			"numeral 𓁨",
			1000000,
			"𓁨",
		},
	}
	samples = []tests{
		// Samples
		{
			"sample 𓆼𓆼𓆼𓆼𓍢𓍢𓍢𓍢𓍢𓍢𓎆𓎆𓏤𓏤",
			4622,
			"𓆼𓆼𓆼𓆼𓍢𓍢𓍢𓍢𓍢𓍢𓎆𓎆𓏤𓏤",
		},
		{
			"sample 𓍢𓍢𓍢𓍢𓍢𓍢𓍢𓎆𓎆𓎆𓎆𓎆𓏤",
			751,
			"𓍢𓍢𓍢𓍢𓍢𓍢𓍢𓎆𓎆𓎆𓎆𓎆𓏤",
		},
	}
	all = [][]tests{
		egyptianNumerals,
		samples,
	}
	converter = NewEgyptianConverter()
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

func BenchmarkParse_numerals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range egyptianNumerals {
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
		for _, tt := range egyptianNumerals {
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

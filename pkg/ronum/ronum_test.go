package ronum

import "testing"

var tests = []struct {
	name  string
	num   int
	ronum Number
}{
	// Numerals
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
	// Subtractive notation
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

func TestRtoi(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rtoi(tt.ronum); got != tt.num {
				t.Errorf("Rtoi() = %v, want %v", got, tt.ronum)
			}
		})
	}
}

func TestItor(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Itor(tt.num); got != tt.ronum {
				t.Errorf("Itor() = %v, want %v", got, tt.ronum)
			}
		})
	}
}

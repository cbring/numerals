package numerals

type Numeral rune
type Number string

type Converter interface {
	Format(int) Number
	Parse(Number) int
}

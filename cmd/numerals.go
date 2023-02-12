package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/gnirb/numerals/pkg/numerals"
	"github.com/gnirb/numerals/pkg/registry"
)

func main() {
	registry := registry.Registry()

	system := flag.String("system", "roman", fmt.Sprintf("The numeral system used. \nSupported systems: %v", registry.List()))
	parse := flag.Bool("parse", false, "Parse numeral system")

	flag.Parse()

	converter := registry.Get(*system)

	var input = flag.Arg(0)
	var output string
	if *parse {
		output = strconv.Itoa(converter.Parse(numerals.Number(flag.Arg(0))))
	} else {
		num, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		output = string(converter.Format(num))
	}

	fmt.Println(output)
}

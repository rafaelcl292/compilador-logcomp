package main

import (
	"os"
	"strconv"

	"compiler/parser"
)

func main() {
	input := os.Args[1]
	result, err := parser.Parse(input)

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	os.Stdout.WriteString(strconv.Itoa(result))
}

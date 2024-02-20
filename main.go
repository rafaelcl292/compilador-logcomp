package main

import (
	"os"
	"strconv"

	"compiler/parser"
)

func main() {
	input := os.Args[1]

	result, err := parser.ParseExpression(input)

	if err != nil {
		println(err.Error())
	} else {
		os.Stdout.WriteString(strconv.Itoa(result))
	}
}

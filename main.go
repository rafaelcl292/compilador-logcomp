package main

import (
	"compiler/parser"
	"compiler/preprocessor"
	"compiler/semantic"
	"compiler/tokenizer"
	"os"
)

func main() {
	file := os.Args[1]
	bytes, err := os.ReadFile(file)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	input := preprocessor.Preprocess(string(bytes))

	tokenizer := tokenizer.CreateTokenizer(input)

	node, err := parser.Parse(tokenizer)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	st := make(semantic.SymbolTable)
	node.Eval(&st)
}

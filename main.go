package main

import (
	"compiler/parser"
	"compiler/preprocessor"
	"compiler/semantic"
	"compiler/tokenizer"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("No filename provided")
		os.Exit(1)
	}

	file := os.Args[1]
	bytes, err := os.ReadFile(file)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	input := preprocessor.Preprocess(string(bytes))

	tokenizer := tokenizer.CreateTokenizer(input)

	node := parser.Parse(tokenizer)

	st := make(semantic.SymbolTable)
	node.Eval(&st)
}

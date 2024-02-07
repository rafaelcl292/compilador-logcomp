package main

import (
	"os"

	"compiler/tokenizer"
)

func main() {
	input := os.Args[0]
	_ = tokenizer.CreateTokenizer(input)
}

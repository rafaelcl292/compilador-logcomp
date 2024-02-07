package main

import (
	"os"
	"strconv"

	"compiler/tokenizer"
)

func main() {
	input := os.Args[1]
	tok := tokenizer.CreateTokenizer(input)

	tokens := []tokenizer.Token{}
	for {
		token := tok.NextToken()
		if token.Type == tokenizer.EOF {
			break
		}
		if token.Type == tokenizer.ILLEGAL {
			println("Error: invalid token", token.Literal)
			return
		}
		tokens = append(tokens, token)
	}

	if len(tokens)%2 == 0 {
		println("Error: invalid expression")
		return
	}

	for i, token := range tokens {
		switch i % 2 {
		case 0:
			if token.Type != tokenizer.NUMBER {
				println("Error: expected a number but got", token.Literal)
				return
			}
		case 1:
			if token.Type != tokenizer.PLUS && token.Type != tokenizer.MINUS {
				println("Error: expected an operator but got", token.Literal)
				return
			}
		}
	}

	result, _ := strconv.Atoi(tokens[0].Literal)
	reg := 0
	for i := 0; i < len(tokens)-1; i += 2 {
		switch tokens[i+1].Type {
		case tokenizer.PLUS:
			reg, _ = strconv.Atoi(tokens[i+2].Literal)
			result += reg

		case tokenizer.MINUS:
			reg, _ = strconv.Atoi(tokens[i+2].Literal)
			result -= reg
		}
	}

	os.Stdout.WriteString(strconv.Itoa(result))
}

package parser

import (
	"compiler/tokenizer"
	"errors"
	"strconv"
)

func ParseExpression(input string) (int, error) {
	tok := tokenizer.CreateTokenizer(input)
	result := 0
	token := tok.NextToken()
	if token.Type != tokenizer.NUMBER {
		return 0, errors.New("Error: 1 expected a number but got " + token.Literal)
	}
	result, _ = strconv.Atoi(token.Literal)
	for {
		token = tok.NextToken()
		switch token.Type {
		case tokenizer.EOF:
			return result, nil
		case tokenizer.PLUS:
			token = tok.NextToken()
			if token.Type != tokenizer.NUMBER {
				return 0, errors.New("Error: expected a number but got " + token.Literal)
			}
			reg, _ := strconv.Atoi(token.Literal)
			result += reg
		case tokenizer.MINUS:
			token = tok.NextToken()
			if token.Type != tokenizer.NUMBER {
				return 0, errors.New("Error: expected a number but got " + token.Literal)
			}
			reg, _ := strconv.Atoi(token.Literal)
			result -= reg
		default:
			return 0, errors.New("Error: expected an operator but got " + token.Literal)
		}
	}
}

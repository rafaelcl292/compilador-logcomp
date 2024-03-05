package parser

import (
	"compiler/tokenizer"
	"errors"
	"fmt"
	"strconv"
)

func Parse(input string) (int, error) {
	tok := tokenizer.CreateTokenizer(input)
	num, err := expression(tok)
	if err != nil {
		return 0, err
	}
	if tok.Next.Type != tokenizer.EOF {
		return 0, createError(tokenizer.EOF, tok.Next)
	}
	return num, nil
}

func expression(tok *tokenizer.Tokenizer) (int, error) {
	reg, err := term(tok)
	if err != nil {
		return reg, err
	}
	for {
		switch tok.Next.Type {
		case tokenizer.PLUS:
			tok.NextToken()
			num, err := term(tok)
			if err != nil {
				return reg, err
			}
			reg += num
		case tokenizer.MINUS:
			tok.NextToken()
			num, err := term(tok)
			if err != nil {
				return reg, err
			}
			reg -= num
		default:
			return reg, nil
		}
	}
}

func term(tok *tokenizer.Tokenizer) (int, error) {
	reg, err := factor(tok)
	if err != nil {
		return reg, err
	}
	for {
		switch tok.Next.Type {
		case tokenizer.MULTIPLY:
			tok.NextToken()
			num, err := factor(tok)
			if err != nil {
				return reg, err
			}
			reg *= num
		case tokenizer.DIVIDE:
			tok.NextToken()
			num, err := factor(tok)
			if err != nil {
				return reg, err
			}
			reg /= num
		default:
			return reg, nil
		}
	}
}

func factor(tok *tokenizer.Tokenizer) (int, error) {
	switch tok.Next.Type {
	case tokenizer.NUMBER:
		num, _ := strconv.Atoi(tok.Next.Literal)
		tok.NextToken()
		return num, nil
	case tokenizer.PLUS:
		tok.NextToken()
		num, err := factor(tok)
		if err != nil {
			return 0, err
		}
		return num, nil
	case tokenizer.MINUS:
		tok.NextToken()
		num, err := factor(tok)
		if err != nil {
			return 0, err
		}
		return -num, nil
	case tokenizer.LPAREN:
		tok.NextToken()
		num, err := expression(tok)
		if err != nil {
			return 0, err
		}
		if tok.Next.Type != tokenizer.RPAREN {
			return 0, createError(tokenizer.RPAREN, tok.Next)
		}
		tok.NextToken()
		return num, nil
	default:
		// return 0, createError(tokenizer.NUMBER, tok.Next)
        return 0, errors.New("testee")
	}
}

func createError(expected tokenizer.TokenType, token tokenizer.Token) error {
	msg := fmt.Sprintf(
		"Error: expected %s but got %s '%s'",
		expected,
		token.Type,
		token.Literal,
	)
	return errors.New(msg)
}

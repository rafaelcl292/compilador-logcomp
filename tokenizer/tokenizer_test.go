package tokenizer

import (
	. "compiler/tokens"
	"testing"
)

func TestTokenizer(t *testing.T) {
	inputs := []string{
		"11  + 2",
		"1+2   - 33 - 4",
		"34--#",
		"01  - 2a",
		"1+2-3*4/5   ",
		"  * /0+-",
		"(1 + 2) / 5",
		"4/(1+1)*2",
		"print(1+2)",
		"x1 = 2\nprint(carro_especial)",
		"4dd",
	}
	tokens := [][]Token{
		{
			{Type: INTEGER, Literal: "11"},
			{Type: PLUS, Literal: "+"},
			{Type: INTEGER, Literal: "2"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: INTEGER, Literal: "1"},
			{Type: PLUS, Literal: "+"},
			{Type: INTEGER, Literal: "2"},
			{Type: MINUS, Literal: "-"},
			{Type: INTEGER, Literal: "33"},
			{Type: MINUS, Literal: "-"},
			{Type: INTEGER, Literal: "4"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: INTEGER, Literal: "34"},
			{Type: MINUS, Literal: "-"},
			{Type: MINUS, Literal: "-"},
			{Type: ILLEGAL, Literal: "#"},
		},
		{
			{Type: INTEGER, Literal: "01"},
			{Type: MINUS, Literal: "-"},
			{Type: ILLEGAL, Literal: "2a"},
		},
		{
			{Type: INTEGER, Literal: "1"},
			{Type: PLUS, Literal: "+"},
			{Type: INTEGER, Literal: "2"},
			{Type: MINUS, Literal: "-"},
			{Type: INTEGER, Literal: "3"},
			{Type: MULTIPLY, Literal: "*"},
			{Type: INTEGER, Literal: "4"},
			{Type: DIVIDE, Literal: "/"},
			{Type: INTEGER, Literal: "5"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: MULTIPLY, Literal: "*"},
			{Type: DIVIDE, Literal: "/"},
			{Type: INTEGER, Literal: "0"},
			{Type: PLUS, Literal: "+"},
			{Type: MINUS, Literal: "-"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: LPAREN, Literal: "("},
			{Type: INTEGER, Literal: "1"},
			{Type: PLUS, Literal: "+"},
			{Type: INTEGER, Literal: "2"},
			{Type: RPAREN, Literal: ")"},
			{Type: DIVIDE, Literal: "/"},
			{Type: INTEGER, Literal: "5"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: INTEGER, Literal: "4"},
			{Type: DIVIDE, Literal: "/"},
			{Type: LPAREN, Literal: "("},
			{Type: INTEGER, Literal: "1"},
			{Type: PLUS, Literal: "+"},
			{Type: INTEGER, Literal: "1"},
			{Type: RPAREN, Literal: ")"},
			{Type: MULTIPLY, Literal: "*"},
			{Type: INTEGER, Literal: "2"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: PRINT, Literal: "print"},
			{Type: LPAREN, Literal: "("},
			{Type: INTEGER, Literal: "1"},
			{Type: PLUS, Literal: "+"},
			{Type: INTEGER, Literal: "2"},
			{Type: RPAREN, Literal: ")"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: VARIABLE, Literal: "x1"},
			{Type: EQUALS, Literal: "="},
			{Type: INTEGER, Literal: "2"},
			{Type: NEWLINE, Literal: "\n"},
			{Type: PRINT, Literal: "print"},
			{Type: LPAREN, Literal: "("},
			{Type: VARIABLE, Literal: "carro_especial"},
			{Type: RPAREN, Literal: ")"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: ILLEGAL, Literal: "4d"},
		},
	}

	for i, input := range inputs {
		tok := CreateTokenizer(input)
		for j, expected := range tokens[i] {
			actual := tok.Next
			if actual != expected {
				t.Errorf(
					"Expected \"%v\" of type %v, got \"%v\" of type %v",
					expected.Literal,
					expected.Type,
					actual.Literal,
					actual.Type,
				)
			}
			if j < len(tokens[i])-1 {
				tok.NextToken()
			}
		}
	}
}

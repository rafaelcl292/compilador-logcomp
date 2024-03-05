package tokenizer

import (
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
	}
	tokens := [][]Token{
		{
			{Type: NUMBER, Literal: "11"},
			{Type: PLUS, Literal: "+"},
			{Type: NUMBER, Literal: "2"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: NUMBER, Literal: "1"},
			{Type: PLUS, Literal: "+"},
			{Type: NUMBER, Literal: "2"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "33"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "4"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: NUMBER, Literal: "34"},
			{Type: MINUS, Literal: "-"},
			{Type: MINUS, Literal: "-"},
			{Type: ILLEGAL, Literal: "#"},
		},
		{
			{Type: NUMBER, Literal: "01"},
			{Type: MINUS, Literal: "-"},
			{Type: ILLEGAL, Literal: "2a"},
		},
		{
			{Type: NUMBER, Literal: "1"},
			{Type: PLUS, Literal: "+"},
			{Type: NUMBER, Literal: "2"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "3"},
			{Type: MULTIPLY, Literal: "*"},
			{Type: NUMBER, Literal: "4"},
			{Type: DIVIDE, Literal: "/"},
			{Type: NUMBER, Literal: "5"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: MULTIPLY, Literal: "*"},
			{Type: DIVIDE, Literal: "/"},
			{Type: NUMBER, Literal: "0"},
			{Type: PLUS, Literal: "+"},
			{Type: MINUS, Literal: "-"},
			{Type: EOF, Literal: ""},
		},
        {
            {Type: LPAREN, Literal: "("},
            {Type: NUMBER, Literal: "1"},
            {Type: PLUS, Literal: "+"},
            {Type: NUMBER, Literal: "2"},
            {Type: RPAREN, Literal: ")"},
            {Type: DIVIDE, Literal: "/"},
            {Type: NUMBER, Literal: "5"},
            {Type: EOF, Literal: ""},
        },
        {
            {Type: NUMBER, Literal: "4"},
            {Type: DIVIDE, Literal: "/"},
            {Type: LPAREN, Literal: "("},
            {Type: NUMBER, Literal: "1"},
            {Type: PLUS, Literal: "+"},
            {Type: NUMBER, Literal: "1"},
            {Type: RPAREN, Literal: ")"},
            {Type: MULTIPLY, Literal: "*"},
            {Type: NUMBER, Literal: "2"},
            {Type: EOF, Literal: ""},
        },
	}

	for i, input := range inputs {
		tok := CreateTokenizer(input)
		for j, expected := range tokens[i] {
			actual := tok.Next
			if actual != expected {
				t.Errorf(
					"Expected \"%v\", got \"%v\"",
					expected.Literal,
					actual.Literal,
				)
			}
			if j < len(tokens[i])-1 {
				tok.NextToken()
			}
		}
	}
}

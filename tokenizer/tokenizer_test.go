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
	}

	for i, input := range inputs {
		tok := CreateTokenizer(input)
		for _, expected := range tokens[i] {
			actual := tok.NextToken()
			if actual != expected {
				t.Errorf(
					"Expected \"%v\", got \"%v\"",
					expected.Literal,
					actual.Literal,
				)
			}
		}
	}
}

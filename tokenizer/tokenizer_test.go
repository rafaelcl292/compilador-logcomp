package tokenizer

import (
	. "compiler/tokens"
	"os"
	"os/exec"
	"testing"
)

func TestTokenizer(t *testing.T) {
	inputs := []string{
		"11  + 2",
		"1+2   - 33 - 4",
		"1+2-3*4/5   ",
		"  * /0+-",
		"(1 + 2) / 5",
		"4/(1+1)*2",
		"print(1+2)",
		"x1 = 2\nprint(carro_especial)",
		"1==2>3<check",
		"if 1 == 2 then read(1) end",
		"while 1 < 2 do a=1 end",
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
			{Type: INTEGER, Literal: "1"},
			{Type: EQUALITY, Literal: "=="},
			{Type: INTEGER, Literal: "2"},
			{Type: GREATER, Literal: ">"},
			{Type: INTEGER, Literal: "3"},
			{Type: LESS, Literal: "<"},
			{Type: VARIABLE, Literal: "check"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: IF, Literal: "if"},
			{Type: INTEGER, Literal: "1"},
			{Type: EQUALITY, Literal: "=="},
			{Type: INTEGER, Literal: "2"},
			{Type: THEN, Literal: "then"},
			{Type: READ, Literal: "read"},
			{Type: LPAREN, Literal: "("},
			{Type: INTEGER, Literal: "1"},
			{Type: RPAREN, Literal: ")"},
			{Type: END, Literal: "end"},
			{Type: EOF, Literal: ""},
		},
		{
			{Type: WHILE, Literal: "while"},
			{Type: INTEGER, Literal: "1"},
			{Type: LESS, Literal: "<"},
			{Type: INTEGER, Literal: "2"},
			{Type: DO, Literal: "do"},
			{Type: VARIABLE, Literal: "a"},
			{Type: EQUALS, Literal: "="},
			{Type: INTEGER, Literal: "1"},
			{Type: END, Literal: "end"},
			{Type: EOF, Literal: ""},
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

func TestTokenizerError(t *testing.T) {
	inputs := []string{
		"34--#",
		"01  - 2a",
		"4dd",
	}
	for _, input := range inputs {
		flag := input
		if os.Getenv("FLAG") == flag {
			tok := CreateTokenizer(input)
			for tok.Next.Type != EOF {
				tok.NextToken()
			}
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestTokenizerError")
		cmd.Env = append(os.Environ(), "FLAG="+flag)
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			continue
		}
		t.Fatalf("process ran without error for input '%s'", input)
	}
}

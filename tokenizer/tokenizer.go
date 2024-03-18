package tokenizer

import (
	. "compiler/tokens"
	"os"
	"unicode"
	"unicode/utf8"
)

type Token struct {
	Type    TokenType
	Literal string
}

type Tokenizer struct {
	input string
	ch    rune
	Next  Token
}

func (t *Tokenizer) scan() {
	if len(t.input) == 0 {
		t.ch = 0
		return
	}
	r, size := utf8.DecodeRuneInString(t.input)
	t.input = t.input[size:]
	t.ch = r
}

func (t *Tokenizer) readNumber() {
	var number string
	for unicode.IsDigit(t.ch) {
		number += string(t.ch)
		t.scan()
	}

	if unicode.IsLetter(t.ch) {
		println("Tokenizing error: illegal number " + number + string(t.ch))
		os.Exit(1)
		return
	}

	t.Next = Token{Type: INTEGER, Literal: number}
}

func (t *Tokenizer) readIdentifier() {
	var identifier string
	for unicode.IsLetter(t.ch) || unicode.IsDigit(t.ch) || t.ch == '_' {
		identifier += string(t.ch)
		t.scan()
	}

	switch identifier {
	case "print":
		t.Next = Token{Type: PRINT, Literal: "print"}
	default:
		t.Next = Token{Type: VARIABLE, Literal: identifier}
	}
}

func (t *Tokenizer) NextToken() {
	for unicode.IsSpace(t.ch) && (t.ch != '\n') {
		t.scan()
	}

	switch t.ch {
	case '+':
		t.Next = Token{Type: PLUS, Literal: "+"}
	case '-':
		t.Next = Token{Type: MINUS, Literal: "-"}
	case '*':
		t.Next = Token{Type: MULTIPLY, Literal: "*"}
	case '/':
		t.Next = Token{Type: DIVIDE, Literal: "/"}
	case '(':
		t.Next = Token{Type: LPAREN, Literal: "("}
	case ')':
		t.Next = Token{Type: RPAREN, Literal: ")"}
	case '=':
		t.Next = Token{Type: EQUALS, Literal: "="}
	case '\n':
		t.Next = Token{Type: NEWLINE, Literal: "\n"}
	case 0:
		t.Next = Token{Type: EOF, Literal: ""}
	default:
		if unicode.IsDigit(t.ch) {
			t.readNumber()
			return
		}
		if unicode.IsLetter(t.ch) {
			t.readIdentifier()
			return
		}
        println("Tokenizing error: illegal character " + string(t.ch))
        os.Exit(1)
	}
	t.scan()

}

func CreateTokenizer(input string) *Tokenizer {
	tok := &Tokenizer{input: input}
	tok.scan()
	tok.NextToken()
	return tok
}

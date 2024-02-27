package tokenizer

import (
	"unicode"
)

type TokenType int

const (
	// Special tokens
	ILLEGAL TokenType = iota
	EOF

	// Literals
	NUMBER

	// Operators
	PLUS     // +
	MINUS    // -
	MULTIPLY // *
	DIVIDE   // /
)

type Token struct {
	Type    TokenType
	Literal string
}

type Tokenizer struct {
	input []rune
	pos   int
	Next  Token
}

func (t *Tokenizer) scan() rune {
	if t.pos >= len(t.input) {
		return 0
	}
	return t.input[t.pos]
}

func (t *Tokenizer) isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (t *Tokenizer) skipWhitespace() {
	for ch := t.scan(); t.isWhitespace(ch); ch = t.scan() {
		t.pos++
	}
}

func (t *Tokenizer) readNumber() string {
	var number string
	for ch := t.scan(); unicode.IsDigit(ch); ch = t.scan() {
		number += string(ch)
		t.pos++
	}
	return number
}

func (t *Tokenizer) NextToken() {
	t.skipWhitespace()
	ch := t.scan()

	switch {

	case ch == 0:
		t.Next = Token{Type: EOF, Literal: ""}
		return

	case ch == '+':
		t.pos++
		t.Next = Token{Type: PLUS, Literal: "+"}
		return

	case ch == '-':
		t.pos++
		t.Next = Token{Type: MINUS, Literal: "-"}
		return

	case ch == '*':
		t.pos++
		t.Next = Token{Type: MULTIPLY, Literal: "*"}
		return

	case ch == '/':
		t.pos++
		t.Next = Token{Type: DIVIDE, Literal: "/"}
		return

	case unicode.IsDigit(ch):
		number := t.readNumber()
		if ch = t.scan(); unicode.IsLetter(ch) {
			t.pos++
			t.Next = Token{Type: ILLEGAL, Literal: number + string(ch)}
			return
		}
		t.Next = Token{Type: NUMBER, Literal: number}
		return

	default:
		t.pos++
		t.Next = Token{Type: ILLEGAL, Literal: string(ch)}
		return
	}
}

func CreateTokenizer(input string) *Tokenizer {
	tok := &Tokenizer{input: []rune(input)}
	tok.NextToken()
	return tok
}

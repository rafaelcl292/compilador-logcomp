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
	PLUS  // +
	MINUS // -
)

type Token struct {
	Type    TokenType
	Literal string
}

type Tokenizer struct {
	input []rune
	pos   int
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

func (t *Tokenizer) NextToken() Token {
	t.skipWhitespace()
	ch := t.scan()

	switch {

	case ch == 0:
		return Token{Type: EOF, Literal: ""}

	case ch == '+':
		t.pos++
		return Token{Type: PLUS, Literal: "+"}

	case ch == '-':
		t.pos++
		return Token{Type: MINUS, Literal: "-"}

	case unicode.IsDigit(ch):
		number := t.readNumber()
		if ch = t.scan(); unicode.IsLetter(ch) {
			return Token{Type: ILLEGAL, Literal: number + string(ch)}
		}
		return Token{Type: NUMBER, Literal: number}

	default:
		return Token{Type: ILLEGAL, Literal: string(ch)}

	}
}

func CreateTokenizer(input string) *Tokenizer {
	return &Tokenizer{input: []rune(input)}
}

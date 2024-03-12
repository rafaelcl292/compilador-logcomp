package tokenizer

import (
	. "compiler/tokens"
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
	var number []rune
	for unicode.IsDigit(t.ch) {
		number = append(number, t.ch)
		t.scan()
	}

	if unicode.IsLetter(t.ch) {
		t.Next = Token{Type: ILLEGAL, Literal: string(number) + string(t.ch)}
		return
	}

	t.Next = Token{Type: NUMBER, Literal: string(number)}
}

func (t *Tokenizer) NextToken() {
	for unicode.IsSpace(t.ch) {
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
	case 0:
		t.Next = Token{Type: EOF, Literal: ""}
	default:
		if unicode.IsDigit(t.ch) {
			t.readNumber()
			return
		}
		t.Next = Token{Type: ILLEGAL, Literal: string(t.ch)}
	}
	t.scan()

}

func CreateTokenizer(input string) *Tokenizer {
	tok := &Tokenizer{input: input}
	tok.scan()
	tok.NextToken()
	return tok
}

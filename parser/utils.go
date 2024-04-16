package parser

import (
	"compiler/tokenizer"
	. "compiler/tokens"
	"fmt"
	"os"
)

func createError(expected string, token tokenizer.Token) {
	msg := fmt.Sprintf(
		"Parser error: expected %s but got %s '%s'",
		expected,
		token.Type,
		token.Literal,
	)
	println(msg)
	os.Exit(1)
}

func expect(tok *tokenizer.Tokenizer, expect TokenType) {
	if tok.Next.Type != expect {
		createError(string(expect), tok.Next)
	}
	tok.NextToken()
}

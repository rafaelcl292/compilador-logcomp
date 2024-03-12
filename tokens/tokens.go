package tokens

type TokenType string

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Literals
	NUMBER TokenType = "NUMBER"

	// Operators
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	MULTIPLY TokenType = "MULTIPLY"
	DIVIDE   TokenType = "DIVIDE"
	LPAREN   TokenType = "LPAREN"
	RPAREN   TokenType = "RPAREN"
)

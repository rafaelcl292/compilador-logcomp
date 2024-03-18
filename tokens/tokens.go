package tokens

type TokenType string

const (
	// Special tokens
	EOF     TokenType = "EOF"
	NEWLINE TokenType = "NEWLINE"

	// Literals
	INTEGER TokenType = "INTEGER"

	// Keywords
	PRINT TokenType = "PRINT"

	// Operators
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	MULTIPLY TokenType = "MULTIPLY"
	DIVIDE   TokenType = "DIVIDE"
	LPAREN   TokenType = "LPAREN"
	RPAREN   TokenType = "RPAREN"
	EQUALS   TokenType = "EQUALS"

	// Variables
	VARIABLE TokenType = "VARIABLE"
)

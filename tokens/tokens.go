package tokens

type TokenType string

const (
	// Special tokens
	EOF     TokenType = "EOF"
	NEWLINE TokenType = "NEWLINE"
	COMMA   TokenType = "COMMA"

	// Literals
	INTEGER TokenType = "INTEGER"
	STRING  TokenType = "STRING"

	// Keywords
	PRINT    TokenType = "PRINT"
	READ     TokenType = "READ"
	END      TokenType = "END"
	IF       TokenType = "IF"
	THEN     TokenType = "THEN"
	ELSE     TokenType = "ELSE"
	WHILE    TokenType = "WHILE"
	DO       TokenType = "DO"
	OR       TokenType = "OR"
	AND      TokenType = "AND"
	NOT      TokenType = "NOT"
	LOCAL    TokenType = "LOCAL"
	FUNCTION TokenType = "FUNCTION"
	RETURN   TokenType = "RETURN"

	// Operators
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	MULTIPLY TokenType = "MULTIPLY"
	DIVIDE   TokenType = "DIVIDE"
	LPAREN   TokenType = "LPAREN"
	RPAREN   TokenType = "RPAREN"
	EQUALS   TokenType = "EQUALS"
	LESS     TokenType = "LESS"
	GREATER  TokenType = "GREATER"
	EQUALITY TokenType = "EQUALITY"
	CONCAT   TokenType = "CONCAT"

	// Variables
	VARIABLE TokenType = "VARIABLE"
)

package token

// TokenType ...
type TokenType string

const (
	// Special
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENTIFIER TokenType = "IDENTIFIER"
	INTEGER    TokenType = "INTEGER"

	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	// LPAREN ...
	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

// Token ...
type Token struct {
	Type    TokenType
	Literal string
}

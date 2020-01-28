package token

// Token describes the physical representation of a token.
type Token struct {
	Type    Type
	Literal string
}

// Type describes the internal representation of a token.
type Type string

const (
	// -- Special --

	// ILLEGAL refers to an unknown token type.
	ILLEGAL Type = "ILLEGAL"
	// EOF is the end of line delimiter.
	EOF Type = "EOF"

	// -- Variables --

	// IDENTIFIER refers to a variable label.
	IDENTIFIER Type = "IDENTIFIER"

	// -- Literals --

	// INTEGER is a number literal
	INTEGER Type = "INTEGER"

	// -- Operators --

	// ASSIGN is the assigment operator.
	ASSIGN Type = "="
	// PLUS is the sum operator.
	PLUS Type = "+"
	// MINUS is the substract operator.
	MINUS Type = "-"
	// ASTERISK is the multiply operator.
	ASTERISK Type = "*"
	// SLASH is the division operator.
	SLASH Type = "/"
	// LT is the less than logical operator.
	LT Type = "<"
	// GT is the greater than logical operator.
	GT Type = ">"
	// NOT is the negation logical operator.
	NOT Type = "!"
	// EQ is the equal logical operator.
	EQ Type = "=="
	// NOTEQ is the not equal logical operator.
	NOTEQ Type = "!="

	// -- Delimiters --

	// COMMA is the delimiter for identifiers and literals.
	COMMA Type = ","
	// SEMICOLON is the delimiter for statements.
	SEMICOLON Type = ";"
	// LPAREN is the left parenthesis delimiter.
	LPAREN Type = "("
	// RPAREN is the right parenthesis delimiter.
	RPAREN Type = ")"
	// LBRACE is the left curly braces delimiter.
	LBRACE Type = "{"
	// RBRACE is the right curly braces delimiter.
	RBRACE Type = "}"

	// -- Keywords --

	// TRUE is the physical representation keyword
	// for true logical operations.
	TRUE Type = "TRUE"
	// FALSE is the physical representation keyword
	// for false logical operations.
	FALSE Type = "FALSE"
	// LET is the keyword for variable instancing.
	LET Type = "LET"
	// FUNCTION is a flow keyword
	// for instancing blocks of code.
	FUNCTION Type = "FUNCTION"
	// RETURN is a flow blockcode keyword.
	RETURN Type = "RETURN"
	// IF is a logical flow keyword.
	IF Type = "IF"
	// ELSE is a logical flow keyword.
	ELSE Type = "ELSE"
)

// -- List of reserverd words --
var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdentifier returns whether a literal is an Identifier
// or the corresponding Keyword Token Type.
func LookupIdentifier(identifier string) Type {
	if tokenType, isKeyword := keywords[identifier]; isKeyword {
		return tokenType
	}
	return IDENTIFIER
}

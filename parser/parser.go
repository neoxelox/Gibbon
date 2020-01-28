package parser

import (
	"Gibbon/ast"
	"Gibbon/lexer"
	"Gibbon/token"
)

// Parser describes a Parser instance.
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New creates a Parser instance.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram ...
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

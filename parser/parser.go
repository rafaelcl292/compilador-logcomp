package parser

import (
	"compiler/semantic"
	"compiler/tokenizer"
	. "compiler/tokens"
	"fmt"
	"os"
	"strconv"
)

func Parse(tok *tokenizer.Tokenizer) semantic.Node {
	node := block(tok)
	if tok.Next.Type != EOF {
		createError("EOF", tok.Next)
	}
	return node
}

func block(tok *tokenizer.Tokenizer) semantic.Node {
	stmts := make([]semantic.Node, 0)
	for tok.Next.Type != EOF {
		stmt := statement(tok)
		stmts = append(stmts, stmt)
	}
	return &semantic.Block{Stmts: stmts}
}

func statement(tok *tokenizer.Tokenizer) semantic.Node {
	switch tok.Next.Type {
	case PRINT:
		tok.NextToken()
		expect(tok, LPAREN)
		expr := expression(tok)
		expect(tok, RPAREN)
		return &semantic.UnOp{Op: "print", Expr: expr}
	case VARIABLE:
		ident := tok.Next.Literal
		tok.NextToken()
		expect(tok, EQUALS)
		expr := expression(tok)
		return &semantic.Assign{Ident: ident, Expr: expr}
	case NEWLINE:
		tok.NextToken()
		return &semantic.NoOp{}
	default:
		createError("STATEMENT", tok.Next)
		return nil
	}
}

func expression(tok *tokenizer.Tokenizer) semantic.Node {
	left := term(tok)
	for {
		if tok.Next.Type == PLUS || tok.Next.Type == MINUS {
			op := tok.Next.Literal
			tok.NextToken()
			right := term(tok)
			left = &semantic.BinOp{Op: op, Left: left, Right: right}
		} else {
			return left
		}
	}
}

func term(tok *tokenizer.Tokenizer) semantic.Node {
	left := factor(tok)
	for {
		if tok.Next.Type == MULTIPLY || tok.Next.Type == DIVIDE {
			op := tok.Next.Literal
			tok.NextToken()
			right := factor(tok)
			left = &semantic.BinOp{Op: op, Left: left, Right: right}
		} else {
			return left
		}
	}
}

func factor(tok *tokenizer.Tokenizer) semantic.Node {
	switch tok.Next.Type {
	case INTEGER:
		value, _ := strconv.Atoi(tok.Next.Literal)
		tok.NextToken()
		return &semantic.IntVal{Val: value}
	case PLUS, MINUS:
		op := tok.Next.Literal
		tok.NextToken()
		node := factor(tok)
		return &semantic.UnOp{Op: op, Expr: node}
	case VARIABLE:
		name := tok.Next.Literal
		tok.NextToken()
		return &semantic.Ident{Name: name}
	case LPAREN:
		tok.NextToken()
		node := expression(tok)
		expect(tok, RPAREN)
		return node
	default:
		createError("EXPRESSION", tok.Next)
		return nil
	}
}

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

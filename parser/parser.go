package parser

import (
	"compiler/semantic"
	"compiler/tokenizer"
	. "compiler/tokens"
	"errors"
	"fmt"
	"strconv"
)

func Parse(tok *tokenizer.Tokenizer) (semantic.Node, error) {
	node, err := block(tok)
	if err != nil {
		return nil, err
	}
	if tok.Next.Type != EOF {
		return nil, createError("EOF", tok.Next)
	}
	return node, nil
}

func block(tok *tokenizer.Tokenizer) (semantic.Node, error) {
	stmts := make([]semantic.Node, 0)
	for tok.Next.Type != EOF {
		stmt, err := statement(tok)
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
	}
	return &semantic.Block{Stmts: stmts}, nil
}

func statement(tok *tokenizer.Tokenizer) (semantic.Node, error) {
	switch tok.Next.Type {
	case PRINT:
		tok.NextToken()
		if tok.Next.Type != LPAREN {
			return nil, createError("LPAREN", tok.Next)
		}
		tok.NextToken()
		expr, err := expression(tok)
		if err != nil {
			return nil, err
		}
		if tok.Next.Type != RPAREN {
			return nil, createError("RPAREN", tok.Next)
		}
		tok.NextToken()
		return &semantic.UnOp{Op: "print", Expr: expr}, nil
	case VARIABLE:
		ident := tok.Next.Literal
		tok.NextToken()
		if tok.Next.Type != EQUALS {
			return nil, createError("EQUALS", tok.Next)
		}
		tok.NextToken()
		expr, err := expression(tok)
		if err != nil {
			return nil, err
		}
		return &semantic.Assign{Ident: ident, Expr: expr}, nil
	case NEWLINE:
		tok.NextToken()
		return &semantic.NoOp{}, nil
	default:
		return nil, createError("STATEMENT", tok.Next)
	}
}

func expression(tok *tokenizer.Tokenizer) (semantic.Node, error) {
	left, err := term(tok)
	if err != nil {
		return nil, err
	}
	for {
		if tok.Next.Type == PLUS || tok.Next.Type == MINUS {
			op := tok.Next.Literal
			tok.NextToken()
			right, err := term(tok)
			if err != nil {
				return nil, err
			}
			left = &semantic.BinOp{Op: op, Left: left, Right: right}
		} else {
			return left, nil
		}
	}
}

func term(tok *tokenizer.Tokenizer) (semantic.Node, error) {
	left, err := factor(tok)
	if err != nil {
		return nil, err
	}
	for {
		if tok.Next.Type == MULTIPLY || tok.Next.Type == DIVIDE {
			op := tok.Next.Literal
			tok.NextToken()
			right, err := factor(tok)
			if err != nil {
				return nil, err
			}
			left = &semantic.BinOp{Op: op, Left: left, Right: right}
		} else {
			return left, nil
		}
	}
}

func factor(tok *tokenizer.Tokenizer) (semantic.Node, error) {
	switch tok.Next.Type {
	case INTEGER:
		value, _ := strconv.Atoi(tok.Next.Literal)
		tok.NextToken()
		return &semantic.IntVal{Val: value}, nil
	case PLUS, MINUS:
		op := tok.Next.Literal
		tok.NextToken()
		node, err := factor(tok)
		if err != nil {
			return nil, err
		}
		return &semantic.UnOp{Op: op, Expr: node}, nil
	case VARIABLE:
		name := tok.Next.Literal
		tok.NextToken()
		return &semantic.Ident{Name: name}, nil
	case LPAREN:
		tok.NextToken()
		node, err := expression(tok)
		if err != nil {
			return nil, err
		}
		if tok.Next.Type != RPAREN {
			return nil, createError("RPAREN", tok.Next)
		}
		tok.NextToken()
		return node, nil
	default:
		return nil, createError("EXPRESSION", tok.Next)
	}
}

func createError(expected string, token tokenizer.Token) error {
	msg := fmt.Sprintf(
		"Error: expected %s but got %s '%s'",
		expected,
		token.Type,
		token.Literal,
	)
	return errors.New(msg)
}

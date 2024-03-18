package semantic

import (
	"fmt"
	"os"
)

type SymbolTable map[string]int

func (st *SymbolTable) Get(ident string) int {
	val, ok := (*st)[ident]
	if !ok {
		println("Undefined variable:", ident)
		os.Exit(1)
	}
	return val
}

func (st *SymbolTable) Set(ident string, val int) {
	(*st)[ident] = val
}

type Node interface {
	Eval(*SymbolTable) int
}

type IntVal struct {
	Val int
}

type Ident struct {
	Name string
}

type NoOp struct{}

type UnOp struct {
	Op   string
	Expr Node
}

type BinOp struct {
	Op    string
	Left  Node
	Right Node
}

type Block struct {
	Stmts []Node
}

type Assign struct {
	Ident string
	Expr  Node
}

func (n IntVal) Eval(st *SymbolTable) int {
	return n.Val
}

func (n Ident) Eval(st *SymbolTable) int {
	return st.Get(n.Name)
}

func (n NoOp) Eval(st *SymbolTable) int {
	return 0
}

func (n UnOp) Eval(st *SymbolTable) int {
	switch n.Op {
	case "+":
		return n.Expr.Eval(st)
	case "-":
		return -n.Expr.Eval(st)
	case "print":
		fmt.Println(n.Expr.Eval(st))
	}
	return 0
}

func (n BinOp) Eval(st *SymbolTable) int {
	switch n.Op {
	case "+":
		return n.Left.Eval(st) + n.Right.Eval(st)
	case "-":
		return n.Left.Eval(st) - n.Right.Eval(st)
	case "*":
		return n.Left.Eval(st) * n.Right.Eval(st)
	case "/":
		return n.Left.Eval(st) / n.Right.Eval(st)
	}
	return 0
}

func (n Block) Eval(st *SymbolTable) int {
	for _, stmt := range n.Stmts {
		stmt.Eval(st)
	}
	return 0
}

func (n Assign) Eval(st *SymbolTable) int {
	st.Set(n.Ident, n.Expr.Eval(st))
	return 0
}

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

type If struct {
	Cond Node
	Then Block
	Else Block
}

type While struct {
	Cond Node
	Do   Block
}

type Read struct{}

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
	case "not":
		if n.Expr.Eval(st) == 0 {
			return 1
		}
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
	case "and":
		if n.Left.Eval(st) != 0 && n.Right.Eval(st) != 0 {
			return 1
		}
	case "or":
		if n.Left.Eval(st) != 0 || n.Right.Eval(st) != 0 {
			return 1
		}
	case "==":
		if n.Left.Eval(st) == n.Right.Eval(st) {
			return 1
		}
	case "<":
		if n.Left.Eval(st) < n.Right.Eval(st) {
			return 1
		}
	case ">":
		if n.Left.Eval(st) > n.Right.Eval(st) {
			return 1
		}
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

func (n If) Eval(st *SymbolTable) int {
	if n.Cond.Eval(st) != 0 {
		n.Then.Eval(st)
	} else {
		n.Else.Eval(st)
	}
	return 0
}

func (n While) Eval(st *SymbolTable) int {
	for n.Cond.Eval(st) != 0 {
		n.Do.Eval(st)
	}
	return 0
}

func (n Read) Eval(st *SymbolTable) int {
	var val int
	fmt.Scan(&val)
	return val
}

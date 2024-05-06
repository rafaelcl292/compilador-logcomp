package semantic

import "reflect"

type Block struct {
	Stmts []Node
}

func (n Block) Eval(st *SymbolTable) symbol {
	for _, stmt := range n.Stmts {
		if reflect.TypeOf(stmt).String() == "*semantic.Return" {
			return stmt.Eval(st)
		}
		stmt.Eval(st)
	}
	return symbol{NONE, nil}
}

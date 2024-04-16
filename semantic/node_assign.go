package semantic

type Assign struct {
	Ident string
	Expr  Node
}

func (n Assign) Eval(st *SymbolTable) symbol {
	st.set(n.Ident, n.Expr.Eval(st))
	return symbol{NONE, nil}
}

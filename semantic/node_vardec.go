package semantic

type VarDec struct {
	Ident string
	Expr  Node
}

func (n VarDec) Eval(st *SymbolTable) symbol {
	st.create(n.Ident)
	if n.Expr != nil {
		st.set(n.Ident, n.Expr.Eval(st))
	}
	return symbol{NONE, nil}
}

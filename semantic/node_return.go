package semantic


type Return struct {
	Expr Node
}

func (n Return) Eval(st *SymbolTable) symbol {
	return n.Expr.Eval(st)
}

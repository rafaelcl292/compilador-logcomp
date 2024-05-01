package semantic

type UnOp struct {
	Op   string
	Expr Node
}

func (n UnOp) Eval(st *SymbolTable) {
	n.Expr.Eval(st)
	if n.Op == "-" {
		ASM.append("NEG EAX")
	}
}

package semantic

type VarDec struct {
	Ident string
	Expr  Node
}

func (n VarDec) Eval(st *SymbolTable) {
	ASM.append("PUSH DWORD 0")
	shift := sc.next()
	st.create(n.Ident, shift)
	if n.Expr != nil {
		Assign.Eval(Assign{n.Ident, n.Expr}, st)
	}
}

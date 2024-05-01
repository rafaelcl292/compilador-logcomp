package semantic

type Print struct {
	Expr Node
}

func (n Print) Eval(st *SymbolTable) {
	n.Expr.Eval(st)
	ASM.append("PUSH EAX")
	ASM.append("PUSH formatout")
	ASM.append("CALL printf")
	ASM.append("ADD ESP, 8")
}

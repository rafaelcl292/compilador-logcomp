package semantic

type If struct {
	Cond Node
	Then Block
	Else Block
}

func (n If) Eval(st *SymbolTable) {
	elseLabel, endif := lc.next()
	n.Cond.Eval(st)
	ASM.append("CMP EAX, 0")
	ASM.append("JE " + elseLabel)
	n.Then.Eval(st)
	ASM.append("JMP " + endif)
	ASM.append(elseLabel + ":")
	n.Else.Eval(st)
	ASM.append(endif + ":")
}

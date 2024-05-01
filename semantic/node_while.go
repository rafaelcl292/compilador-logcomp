package semantic

type While struct {
	Cond Node
	Do   Block
}

func (n While) Eval(st *SymbolTable) {
	loop, exit := lc.next()
	ASM.append(loop + ":")
	n.Cond.Eval(st)
	ASM.append("CMP EAX, 0")
	ASM.append("JE " + exit)
	n.Do.Eval(st)
	ASM.append("JMP " + loop)
	ASM.append(exit + ":")
}

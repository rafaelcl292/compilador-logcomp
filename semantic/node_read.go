package semantic

type Read struct{}

func (n Read) Eval(st *SymbolTable) {
	ASM.append("PUSH scanint")
	ASM.append("PUSH formatin")
	ASM.append("call scanf")
	ASM.append("ADD ESP, 8")
	ASM.append("MOV EAX, DWORD [scanint]")
}

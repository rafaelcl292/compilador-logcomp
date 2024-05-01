package semantic

type BinOp struct {
	Op    string
	Left  Node
	Right Node
}

func (n BinOp) Eval(st *SymbolTable) {
	n.Left.Eval(st)
	ASM.append("PUSH EAX")
	n.Right.Eval(st)
	ASM.append("MOV EBX, EAX")
	ASM.append("POP EAX")
	switch n.Op {
	case "+":
		ASM.append("ADD EAX, EBX")
	case "-":
		ASM.append("SUB EAX, EBX")
	case "*":
		ASM.append("IMUL EBX")
	case "/":
		ASM.append("MOV EDX, 0")
		ASM.append("IDIV EBX")
	case "==":
		ASM.append("CMP EAX, EBX")
		ASM.append("SETE AL")
		ASM.append("MOVZX EAX, AL")
	case "!=":
		ASM.append("CMP EAX, EBX")
		ASM.append("SETNE AL")
		ASM.append("MOVZX EAX, AL")
	case "<":
		ASM.append("CMP EAX, EBX")
		ASM.append("SETL AL")
		ASM.append("MOVZX EAX, AL")
	case ">":
		ASM.append("CMP EAX, EBX")
		ASM.append("SETG AL")
		ASM.append("MOVZX EAX, AL")
	case "or":
		ASM.append("OR EAX, EBX")
	case "and":
		ASM.append("AND EAX, EBX")
	}
}

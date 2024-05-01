package semantic

import "fmt"

func (n NoOp) Eval(st *SymbolTable) {
}

func (n UnOp) Eval(st *SymbolTable) {
	n.Expr.Eval(st)
	if n.Op == "-" {
		ASM.append("NEG EAX")
	}
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

func (n IntVal) Eval(st *SymbolTable) {
	ASM.append(fmt.Sprintf("MOV EAX, %d", n.Val))
}

func (n Ident) Eval(st *SymbolTable) {
	ASM.append(fmt.Sprintf("MOV EAX, [EBP-%d]", st.get(n.Name)))
}

func (n Assign) Eval(st *SymbolTable) {
	n.Expr.Eval(st)
	st.set(n.Ident)
	ASM.append(fmt.Sprintf("MOV [EBP-%d], EAX", st.get(n.Ident)))
}

func (n VarDec) Eval(st *SymbolTable) {
	ASM.append("PUSH DWORD 0")
	shift := sc.next()
	st.create(n.Ident, shift)
	if n.Expr != nil {
		Assign.Eval(Assign{n.Ident, n.Expr}, st)
	}
}

func (n Print) Eval(st *SymbolTable) {
	n.Expr.Eval(st)
	ASM.append("PUSH EAX")
	ASM.append("PUSH formatout")
	ASM.append("CALL printf")
	ASM.append("ADD ESP, 8")
}

func (n Read) Eval(st *SymbolTable) {
	ASM.append("PUSH scanint")
	ASM.append("PUSH formatin")
	ASM.append("call scanf")
	ASM.append("ADD ESP, 8")
	ASM.append("MOV EAX, DWORD [scanint]")
}

func (n Block) Eval(st *SymbolTable) {
	for _, stmt := range n.Stmts {
		stmt.Eval(st)
	}
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

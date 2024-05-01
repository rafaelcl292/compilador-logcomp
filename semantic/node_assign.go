package semantic

import "fmt"

type Assign struct {
	Ident string
	Expr  Node
}

func (n Assign) Eval(st *SymbolTable) {
	n.Expr.Eval(st)
	st.set(n.Ident)
	ASM.append(fmt.Sprintf("MOV [EBP-%d], EAX", st.get(n.Ident)))
}

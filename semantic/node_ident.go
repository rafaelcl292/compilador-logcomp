package semantic

import "fmt"

type Ident struct {
	Name string
}

func (n Ident) Eval(st *SymbolTable) {
	ASM.append(fmt.Sprintf("MOV EAX, [EBP-%d]", st.get(n.Name)))
}

package semantic

import "fmt"

type IntVal struct {
	Val int
}

func (n IntVal) Eval(st *SymbolTable) {
	ASM.append(fmt.Sprintf("MOV EAX, %d", n.Val))
}

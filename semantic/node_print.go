package semantic

import "fmt"

type Print struct{
	Expr Node
}

func (n Print) Eval(st *SymbolTable) symbol {
	s := n.Expr.Eval(st)
	fmt.Println(s.val)
	return s
}

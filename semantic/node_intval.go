package semantic

type IntVal struct {
	Val int
}

func (n IntVal) Eval(st *SymbolTable) symbol {
	return symbol{INT, n.Val}
}

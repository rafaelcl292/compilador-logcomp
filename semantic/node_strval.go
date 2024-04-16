package semantic


type StrVal struct {
	Val string
}

func (n StrVal) Eval(st *SymbolTable) symbol {
	return symbol{STRING, n.Val}
}

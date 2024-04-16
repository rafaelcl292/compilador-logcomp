package semantic

type Ident struct {
	Name string
}

func (n Ident) Eval(st *SymbolTable) symbol {
	return st.get(n.Name)
}

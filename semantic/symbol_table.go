package semantic

type SymbolTable map[string]symbol

func (st *SymbolTable) get(ident string) symbol {
	symbol, ok := (*st)[ident]
	if !ok {
		errorf("Undefined variable '%s'", ident)
	}
	if symbol.stype == NONE {
		errorf("Uninitialized variable '%s'", ident)
	}
	return symbol
}

func (st *SymbolTable) set(ident string, symbol symbol) {
	_, ok := (*st)[ident]
	if !ok {
		errorf("Undefined variable '%s'", ident)
	}
	(*st)[ident] = symbol
}

func (st *SymbolTable) create(ident string) {
	_, ok := (*st)[ident]
	if ok {
		errorf("Variable '%s' already exists", ident)
	}
	(*st)[ident] = symbol{NONE, 0}
}

package semantic

func (st *SymbolTable) get(ident string) int {
	s, ok := (*st)[ident]
	if !ok {
		errorf("Undefined variable '%s'", ident)
	}
	if !s.initialized {
		errorf("Uninitialized variable '%s'", ident)
	}
	return s.shift
}

func (st *SymbolTable) set(ident string) {
	s, ok := (*st)[ident]
	if !ok {
		errorf("Undefined variable '%s'", ident)
	}
	(*st)[ident] = symbol{true, s.shift}
}

func (st *SymbolTable) create(ident string, shift int) {
	_, ok := (*st)[ident]
	if ok {
		errorf("Variable '%s' already exists", ident)
	}
	(*st)[ident] = symbol{false, shift}
}

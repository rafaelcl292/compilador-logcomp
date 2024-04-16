package semantic

type If struct {
	Cond Node
	Then Block
	Else Block
}

func (n If) Eval(st *SymbolTable) symbol {
	s := n.Cond.Eval(st)
	expect(INT, s)
	if s.val.(int) != 0 {
		n.Then.Eval(st)
	} else {
		n.Else.Eval(st)
	}
	return symbol{NONE, nil}
}

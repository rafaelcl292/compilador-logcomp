package semantic

type While struct {
	Cond Node
	Do   Block
}

func (n While) Eval(st *SymbolTable) symbol {
	s := n.Cond.Eval(st)
	expect(INT, s)
	for s.val.(int) != 0 {
		n.Do.Eval(st)
		s = n.Cond.Eval(st)
		expect(INT, s)
	}
	return symbol{NONE, nil}
}

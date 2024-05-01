package semantic

type Block struct {
	Stmts []Node
}

func (n Block) Eval(st *SymbolTable) {
	for _, stmt := range n.Stmts {
		stmt.Eval(st)
	}
}

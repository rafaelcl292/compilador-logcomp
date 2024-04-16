package semantic

type UnOp struct {
	Op   string
	Expr Node
}

func (n UnOp) Eval(st *SymbolTable) symbol {
	s := n.Expr.Eval(st)
	expect(INT, s)
	switch n.Op {
	case "+":
		return s
	case "-":
		return symbol{INT, -s.val.(int)}
	case "not":
		if s.val.(int) == 0 {
			return symbol{INT, 1}
		}
	}
	return symbol{INT, 0}
}

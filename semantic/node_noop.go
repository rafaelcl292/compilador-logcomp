package semantic

type NoOp struct{}

func (n NoOp) Eval(st *SymbolTable) {
}

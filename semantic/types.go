package semantic

type stype string

const (
	NONE   stype = "NONE"
	INT    stype = "INT"
	STRING stype = "STRING"
)

type symbol struct {
	stype stype
	val   interface{}
}

type Node interface {
	Eval(*SymbolTable) symbol
}

// TODO: testar eliminar NoOp
type NoOp struct{}

func (n NoOp) Eval(st *SymbolTable) symbol {
	return symbol{NONE, 0}
}

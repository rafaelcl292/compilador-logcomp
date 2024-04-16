package semantic

import "fmt"

type Read struct{}

func (n Read) Eval(st *SymbolTable) symbol {
	var val int
	fmt.Scan(&val)
	return symbol{INT, val}
}

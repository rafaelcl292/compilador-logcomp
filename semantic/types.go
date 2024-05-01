package semantic

type symbol struct {
	initialized bool
	shift       int
}

type SymbolTable map[string]symbol

type Node interface {
	Eval(*SymbolTable)
}

type NoOp struct{}

type UnOp struct {
	Op   string
	Expr Node
}

type BinOp struct {
	Op    string
	Left  Node
	Right Node
}

type IntVal struct {
	Val int
}

type Ident struct {
	Name string
}

type Assign struct {
	Ident string
	Expr  Node
}

type VarDec struct {
	Ident string
	Expr  Node
}

type Print struct {
	Expr Node
}

type Read struct{}

type Block struct {
	Stmts []Node
}

type If struct {
	Cond Node
	Then Block
	Else Block
}

type While struct {
	Cond Node
	Do   Block
}

type labelCounter struct {
	count int
}

type shiftCounter struct {
	count int
}

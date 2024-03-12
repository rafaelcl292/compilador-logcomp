package semantic

type Node interface {
	Eval() int
}

type IntVal struct {
	Val int
}

type UnOp struct {
	Op   string
	Expr Node
}

type BinOp struct {
	Op    string
	Left  Node
	Right Node
}

func (n IntVal) Eval() int {
	return n.Val
}

func (n UnOp) Eval() int {
	switch n.Op {
	case "+":
		return n.Expr.Eval()
	case "-":
		return -n.Expr.Eval()
	}
	return 0
}

func (n BinOp) Eval() int {
	switch n.Op {
	case "+":
		return n.Left.Eval() + n.Right.Eval()
	case "-":
		return n.Left.Eval() - n.Right.Eval()
	case "*":
		return n.Left.Eval() * n.Right.Eval()
	case "/":
		return n.Left.Eval() / n.Right.Eval()
	}
	return 0
}

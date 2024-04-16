package semantic

import "strconv"

type BinOp struct {
	Op    string
	Left  Node
	Right Node
}

func (n BinOp) Eval(st *SymbolTable) symbol {
	left := n.Left.Eval(st)
	right := n.Right.Eval(st)

	switch n.Op {
	case "+", "-", "*", "/", "and", "or":
		expect(INT, left)
		expect(INT, right)
		switch n.Op {
		case "+":
			return symbol{INT, left.val.(int) + right.val.(int)}
		case "-":
			return symbol{INT, left.val.(int) - right.val.(int)}
		case "*":
			return symbol{INT, left.val.(int) * right.val.(int)}
		case "/":
			return symbol{INT, left.val.(int) / right.val.(int)}
		case "and":
			if left.val.(int) != 0 && right.val.(int) != 0 {
				return symbol{INT, 1}
			}
		case "or":
			if left.val.(int) != 0 || right.val.(int) != 0 {
				return symbol{INT, 1}
			}
		}
	case "==", "<", ">":
		if left.stype != right.stype {
			errorf("type mismatch: %s %s %s", left.stype, n.Op, right.stype)
		}
		switch n.Op {
		case "==":
			if left.val == right.val {
				return symbol{INT, 1}
			}
		case "<":
			switch left.val.(type) {
			case int:
				if left.val.(int) < right.val.(int) {
					return symbol{INT, 1}
				}
			case string:
				if left.val.(string) < right.val.(string) {
					return symbol{INT, 1}
				}
			}
		case ">":
			switch left.val.(type) {
			case int:
				if left.val.(int) > right.val.(int) {
					return symbol{INT, 1}
				}
			case string:
				if left.val.(string) > right.val.(string) {
					return symbol{INT, 1}
				}
			}
		}
	case "..":
		if left.stype == INT {
			left.val = strconv.Itoa(left.val.(int))
		}
		if right.stype == INT {
			right.val = strconv.Itoa(right.val.(int))
		}
		return symbol{STRING, left.val.(string) + right.val.(string)}

	}
	return symbol{INT, 0}
}

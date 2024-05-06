package semantic

type FuncCall struct {
	Name string
	Args Block
}

func (n FuncCall) Eval(st *SymbolTable) symbol {
	funcdec := ft.get(n.Name)
	if len(n.Args.Stmts) != len(funcdec.Vars) {
		errorf("Wrong number of arguments for function %s", n.Name)
	}
	localSt := make(SymbolTable)
	var arg symbol
	for i := 0; i < len(funcdec.Vars); i++ {
		funcdec.Vars[i].Eval(&localSt)
		arg = n.Args.Stmts[i].Eval(st)
		localSt.set(funcdec.Vars[i].Ident, arg)
	}
	return funcdec.FuncBlock.Eval(&localSt)
}

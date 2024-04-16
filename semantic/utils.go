package semantic

import (
	"fmt"
	"os"
)

func errorf(format string, args ...interface{}) {
	println(fmt.Sprintf(format, args...))
	os.Exit(1)
}

func expect(expected stype, s symbol) {
	if s.stype != expected {
		errorf("AST error: expected %s, got %s", expected, s.stype)
	}
}

package semantic

import (
	"fmt"
	"os"
)

func errorf(format string, args ...interface{}) {
	println("AST error:", fmt.Sprintf(format, args...))
	os.Exit(1)
}

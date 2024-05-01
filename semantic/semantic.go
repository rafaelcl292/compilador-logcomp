package semantic

import "fmt"

type symbol struct {
	initialized bool
	shift       int
}

type Node interface {
	Eval(*SymbolTable)
}

type labelCounter struct {
	count int
}

func (c *labelCounter) next() (string, string) {
	c.count++
	return fmt.Sprintf("L0_%d", c.count), fmt.Sprintf("L1_%d", c.count)
}

type shiftCounter struct {
	count int
}

func (c *shiftCounter) next() int {
	c.count += 4
	return c.count
}

var sc = shiftCounter{0}
var lc = labelCounter{0}
var ASM = createAsmGenerator()

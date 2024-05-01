package semantic

import "fmt"

func (c *labelCounter) next() (string, string) {
	c.count++
	return fmt.Sprintf("L0_%d", c.count), fmt.Sprintf("L1_%d", c.count)
}

func (c *shiftCounter) next() int {
	c.count += 4
	return c.count
}

var sc = shiftCounter{0}
var lc = labelCounter{0}

package preprocessor

import (
	"strings"
)

func Preprocess(input string) string {
	if !strings.Contains(input, "--") {
		return input
	}

	return strings.Split(input, "--")[0]
}

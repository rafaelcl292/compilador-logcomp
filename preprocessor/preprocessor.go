package preprocessor

import (
	"regexp"
	"strings"
)

func Preprocess(input string) string {
	lines := strings.Split(input, "\n")
	pattern := regexp.MustCompile(`--.*$`)
	processed := make([]string, len(lines))
	for i, line := range lines {
		processed[i] = pattern.ReplaceAllString(line, "")
	}
	return strings.Join(processed, "\n")
}

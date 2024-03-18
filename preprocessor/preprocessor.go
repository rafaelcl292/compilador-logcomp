package preprocessor

import (
	"regexp"
	"strings"
)

// func Preprocess(input string) string {
// 	pattern := regexp.MustCompile(`--.*$`)
// 	return pattern.ReplaceAllString(input, "")
// }

func Preprocess(input string) string {
	lines := strings.Split(input, "\n") // Split the input string into lines
	pattern := regexp.MustCompile(`--.*$`)
	processed := make([]string, len(lines))
	for i, line := range lines {
		processed[i] = pattern.ReplaceAllString(line, "") // Apply Preprocess to each line
	}
	return strings.Join(processed, "\n")
}

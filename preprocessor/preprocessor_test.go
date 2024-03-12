package preprocessor

import (
	"testing"
)

func TestPreprocess(t *testing.T) {
	inputs := []string{
		"11 + 2 -- comment",
		"hello world",
		"a b c -- comment",
		"-- comment",
	}
	expected := []string{
		"11 + 2 ",
		"hello world",
		"a b c ",
		"",
	}

	for i, input := range inputs {
        result := Preprocess(input)
        if result != expected[i] {
            t.Errorf("Preprocess('%s') != '%s', got '%s'", input, expected[i], result)
        }
	}
}

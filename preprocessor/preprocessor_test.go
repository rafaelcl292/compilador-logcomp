package preprocessor

import "testing"

func TestPreprocess(t *testing.T) {
	inputs := []string{
		"11 + 2 -- comment",
		"hello world",
		"a b c --comment",
		"-- **comment**",
		"print(1 -1) -- 1 -- comment\n--oi",
		"print(1000)--",
		`
        --print(1 -1) -- 1 -- comment
        - print(1000) --
        `,
	}
	expected := []string{
		"11 + 2 ",
		"hello world",
		"a b c ",
		"",
		"print(1 -1) \n",
		"print(1000)",
		`
        
        - print(1000) 
        `,
	}

	for i, input := range inputs {
		result := Preprocess(input)
		if result != expected[i] {
			t.Errorf("Preprocess('%s') != '%s', got '%s'", input, expected[i], result)
		}
	}
}

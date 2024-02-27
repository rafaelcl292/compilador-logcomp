package main

import (
	"os"
	"os/exec"
	"strconv"
	"testing"
)

type test struct {
	input    string
	expected string
}

func TestMain(t *testing.T) {
	tests := []test{
		{"5", "5"},
		{"10+10", "20"},
		{"1 -1   - 1", "-1"},
		{"19-9-9+1", "2"},
		{"1+2*3", "7"},
		{"10/10*10-10+10/10", "1"},
		{"\n1/1   +  2 *  1 -3+4/2", "2"},
	}

	for _, test := range tests {
		os.Args[1] = test.input
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		os.Stdout = w
		main()
		w.Close()

		out := make([]byte, 100)
		n, err := r.Read(out)
		if err != nil {
			t.Fatal(err)
		}

		if string(out[:n]) != test.expected {
			t.Fatalf(
				"expected '%s', got '%s' for input '%s'",
				test.expected, string(out[:n]), test.input,
			)
		}
	}
}

func TestMainError(t *testing.T) {
	inputs := []string{"1 2", ",*", "1 ++4", "-1-10"}

	for i, input := range inputs {
		flag := strconv.Itoa(i)
		if os.Getenv("FLAG") == flag {
			os.Args[1] = input
			main()
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestMainError")
		cmd.Env = append(os.Environ(), "FLAG="+flag)
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			continue
		}
		t.Fatalf("process ran without error for input '%s'", input)
	}
}

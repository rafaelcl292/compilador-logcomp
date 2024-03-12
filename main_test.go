package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	expected_outputs := []string{
		"5", "20", "-1", "2", "7", "1", "2", "3", "3", "4", "3", "6",
	}

	for i, expected := range expected_outputs {
		input_file := "testdata/success/" + fmt.Sprintf("%02d", i)
		os.Args[1] = input_file
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

		if string(out[:n]) != expected {
			t.Fatalf(
				"expected '%s', got '%s' for test file testdata/%s",
				expected, string(out[:n]), input_file,
			)
		}
	}
}

func TestMainError(t *testing.T) {
	for i := 0; i < 6; i++ {
		flag := fmt.Sprintf("%02d", i)
		input_file := "testdata/error/" + flag
		if os.Getenv("FLAG") == flag {
			os.Args[1] = input_file
			main()
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestMainError")
		cmd.Env = append(os.Environ(), "FLAG="+flag)
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			continue
		}
		t.Fatalf("process ran without error for input '%s'", input_file)
	}
}

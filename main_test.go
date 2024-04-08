package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	outputs := []string{
		"1\n", "30\n", "0\n1000\n", "", "9\n", "120\n0\n",
		"2\n-1\n", "1\n0\n", "", "10\n1\n2\n3\n4\n5\n6\n7\n8\n9\n0\n",
	}

	for i, expected := range outputs {
		filename := "testdata/success/" + fmt.Sprintf("%02d", i)
		os.Args[1] = filename
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		os.Stdout = w
		main()
		w.Close()

		out := make([]byte, 100)
		n, err := r.Read(out)
		if err != nil && err.Error() != "EOF" {
			t.Fatal(err, "\nFor input: ", filename)
		}

		if string(out[:n]) != expected {
			t.Fatalf(
				"expected '%s', got '%s' for test file testdata/%s",
				expected, string(out[:n]), filename,
			)
		}
	}
}

func TestMainError(t *testing.T) {
	for i := 0; i < 9; i++ {
		flag := fmt.Sprintf("%02d", i)
		filename := "testdata/error/" + flag
		if os.Getenv("FLAG") == flag {
			os.Args[1] = filename
			main()
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestMainError")
		cmd.Env = append(os.Environ(), "FLAG="+flag)
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			continue
		}
		t.Fatalf("process ran without error for input '%s'", filename)
	}
}

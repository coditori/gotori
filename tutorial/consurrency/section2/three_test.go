package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_sample3(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

    sample3()

    w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout
	if !strings.Contains(output, "$34320.00") {
		t.Errorf("expected $34430.00, got %s", output)
	}
}

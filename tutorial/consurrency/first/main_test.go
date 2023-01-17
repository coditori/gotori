package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

    go printSomethingWithWaitGroup("epislon", &wg)
	wg.Wait()

    w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "epislon") {
		t.Errorf("Expected 'epislon' to be in the output")
	}
}
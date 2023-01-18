package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("epsilon")
	wg.Wait()

	if msg != "epsilon" {
		t.Error("Expected to find epsilon, but it is not there!")
	}
}

func Test_printMessage(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "epsilon") {
		t.Error("Expected to find epsilon, but it is not there!")
	}
}

func Test_sample5(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	sample5SolutionForSample4()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find Hello, universe!, but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos!, but it is not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, world!, but it is not there")
	}
}

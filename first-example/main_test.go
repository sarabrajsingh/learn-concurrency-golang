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

	read, write, _ := os.Pipe()
	os.Stdout = write

	var wg sync.WaitGroup
	wg.Add(1)

	// we need to run the test in the background
	// make sure you're overriding the stdout capture
	go printSomething("epsilon", &wg) // always hand pointers to waitgroups
	wg.Wait()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Excepted to find epsilon, but it is not there")
	}
}

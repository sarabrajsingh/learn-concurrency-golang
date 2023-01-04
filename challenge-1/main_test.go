package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_update_message(t *testing.T) {
	wg.Add(1)

	go updateMessage("epsilon")

	wg.Wait()

	if !strings.Contains(msg, "epsilon") {
		t.Errorf("expected string epsilon")
	}
}

func Test_printMessage(t *testing.T) {
	stdout := os.Stdout

	read, write, _ := os.Pipe()
	os.Stdout = write

	printMessage()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "epsilon") {
		t.Errorf("expecting epsilon in output")
	}
}

func Test_main(t *testing.T) {
	stdout := os.Stdout

	read, write, _ := os.Pipe()
	os.Stdout = write

	main()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("expecting Hello, universe! in output")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("expecting Hello, cosmos! in output")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("expecting Hello, world! in output")
	}
}

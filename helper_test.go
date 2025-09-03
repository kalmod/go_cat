package main

import (
	"io"
	"os"
	"syscall"
	"testing"
)

func TestReadFile_validFile(t *testing.T) {
	defer redirectStdout(os.Stdout)
	os.Stdout = os.NewFile(uintptr(syscall.Stdin), os.DevNull)

	err := readInput("./test.txt")
	if err != nil {
		t.Errorf("File 'test.txt' not found in working directory::%s", err)
	}
}

func TestReadFile_invalidFile(t *testing.T) {
	err := readInput("./sadasffds.fsfs")
	if err == nil {
		t.Errorf("File './sadasffds.fsfs' does not exist. We should get error")
	}
}

func TestReadFile_testStdin(t *testing.T) {
	// create a pipe
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	// Save the original stdin
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }() // Restore original Stdin

	// Repalce stdin with the pipe's read end
	os.Stdin = r

	// Write test input to the pipe
	testInput := "This is a test input!\n"
	_, err = io.WriteString(w, testInput)
	if err != nil {
		t.Fatal(err)
	}
	w.Close() // Close the write end to signal EOF

	defer redirectStdout(os.Stdout)
	os.Stdout = os.NewFile(uintptr(syscall.Stdin), os.DevNull)

	// Now I can run the program to read from stdin
	readInputErr := readInput("-")

	if readInputErr != nil {
		t.Errorf("Could not read from stdin::%s", err)
	}
}

func redirectStdout(stdout *os.File) {
	os.Stdout = stdout
}

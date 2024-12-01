package main

import (
	"os"
	"testing"
)

func Test_Run(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	expected := 11
	got := run(file)

	if expected != got {
		t.Errorf("Expected %d got %d", expected, got)
	}
}

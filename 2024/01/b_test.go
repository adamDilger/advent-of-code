package main

import (
	"os"
	"testing"
)

func Test_RunB(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	expected := 11
	got := runB(file)

	if expected != got {
		t.Errorf("Expected %d got %d", expected, got)
	}
}

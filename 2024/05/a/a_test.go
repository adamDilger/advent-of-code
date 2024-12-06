package a

import (
	"os"
	"testing"
)

func Test_RunA(t *testing.T) {
	file, err := os.Open("../test.txt")
	if err != nil {
		panic(err)
	}

	expected := 143
	got := RunA(file)

	if expected != got {
		t.Errorf("A: Expected %d got %d", expected, got)
	}
}

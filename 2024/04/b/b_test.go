package b

import (
	"os"
	"testing"
)

func Test_RunB(t *testing.T) {
	file, err := os.Open("../test.txt")
	if err != nil {
		panic(err)
	}

	expected := 9
	got := RunB(file)

	if expected != got {
		t.Errorf("B: Expected %d got %d", expected, got)
	}
}

package weekone

import (
	"os"
	"testing"
)

const expectedA = 11
const expectedB = 31

func Test_RunA(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	got := RunA(file)

	if expectedA != got {
		t.Errorf("A: Expected %d got %d", expectedA, got)
	}
}

func Test_RunB(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	got := RunB(file)

	if expectedB != got {
		t.Errorf("B: Expected %d got %d", expectedB, got)
	}
}

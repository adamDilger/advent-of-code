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

	expected := 2
	got := RunA(file)

	if expected != got {
		t.Errorf("A: Expected %d got %d", expected, got)
	}
}

func Test_RunA_properInput(t *testing.T) {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	expected := 332
	got := RunA(file)

	if 332 != got {
		t.Errorf("A: Expected %d got %d", expected, got)
	}
}

func Test_RunA_Line(t *testing.T) {
	safeData := [][]int{
		{7, 6, 4, 2, 1},
		{1, 3, 6, 7, 9},
	}

	unsafeData := [][]int{
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
	}

	for _, s := range safeData {
		if !isSafe(s) {
			t.Errorf("A: Expected true got false for %v", s)
		}
	}

	for _, u := range unsafeData {
		if isSafe(u) {
			t.Errorf("A: Expected false got true for %v", u)
		}
	}
}

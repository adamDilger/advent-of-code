package b

import (
	"os"
	"testing"
)

const expectedB = 4

func Test_RunB(t *testing.T) {
	file, err := os.Open("../test.txt")
	if err != nil {
		panic(err)
	}

	got := RunB(file)

	if expectedB != got {
		t.Errorf("B: Expected %d got %d", expectedB, got)
	}
}

func Test_RunB_Line(t *testing.T) {
	safeData := [][]int{
		{7, 6, 4, 2, 1},
		{7, 6, 4, 2, 10},
		{1, 3, 6, 7, 9},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{100, 3, 6, 7, 9, 10},
		{65, 66, 63, 60},
	}

	unsafeData := [][]int{
		{1, 2, 100, 100},
		{100, 100, 6, 7, 9, 10},
		{100, 101, 6, 7, 9, 10},
		{100, 101, 6, 7, 99, 10},
		{1, 101, 6, 7, 9, 10},
		{1, -100, 6, 7, 99, 10},
		{1, 2, 7, 8, 9},
		{1, 2, 7, 8, 99},
		{9, 7, 6, 2, 1},
	}

	for _, s := range safeData {
		if !isSafeWithError(s) {
			t.Errorf("A: Expected true got false for %v", s)
		}
	}

	for _, u := range unsafeData {
		if isSafeWithError(u) {
			t.Errorf("A: Expected false got true for %v", u)
		}
	}
}

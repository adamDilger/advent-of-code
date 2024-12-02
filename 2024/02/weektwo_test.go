package weektwo

import (
	"os"
	"testing"
)

const expectedA = 2
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

// func Test_RunB(t *testing.T) {
// 	file, err := os.Open("./test.txt")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	got := RunB(file)
//
// 	if expectedB != got {
// 		t.Errorf("B: Expected %d got %d", expectedB, got)
// 	}
// }

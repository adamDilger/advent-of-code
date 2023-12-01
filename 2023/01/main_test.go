package main

import "testing"

func TestAdd(t *testing.T) {

	tests := []struct {
		name string
		want int
	}{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
		{"ninehello", 9},
		{"hellonine", 0},
		{"t", 0},
	}

	for _, tt := range tests {
		got := isWord([]rune(tt.name))
		want := tt.want

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
}

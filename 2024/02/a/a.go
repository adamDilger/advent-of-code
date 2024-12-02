package a

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

func RunA(f io.Reader) int {
	sc := bufio.NewScanner(f)

	var data [][]int

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		dStrings := strings.Split(t, " ")
		dInts := make([]int, 0, len(dStrings))

		for _, d := range dStrings {
			i, err := strconv.Atoi(d)
			if err != nil {
				panic(err)
			}
			dInts = append(dInts, i)
		}

		data = append(data, dInts)
	}

	safeCount := 0
	for _, d := range data {
		if isSafe(d) {
			safeCount++
		}
	}

	return safeCount
}

func isSafe(s []int) bool {
	for i := 1; i < len(s); i++ {
		diff := abs(s[i] - s[i-1])

		if diff < 1 || diff > 3 {
			return false
		}
	}

	if slices.IsSorted(s) {
		return true
	}

	slices.Reverse(s)
	if slices.IsSorted(s) {
		return true
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

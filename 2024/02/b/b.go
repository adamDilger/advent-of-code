package b

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

func RunB(f io.Reader) int {
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
		if isSafeWithError(d) {
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

func isSafeWithError(s []int) bool {
	if isSafe(s) {
		return true
	}

	sModified := make([]int, len(s)-1, len(s)-1)

	// remove one item at a time, and check again
	for skipIndex := 0; skipIndex < len(s); skipIndex++ {
		// fill array with other items
		fillingIndex := 0
		for i := 0; i < len(s); i++ {
			if i == skipIndex {
				continue
			}

			sModified[fillingIndex] = s[i]
			fillingIndex++
		}

		// now try safety check
		if isSafe(sModified) {
			return true
		}
	}

	return false
}

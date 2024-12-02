package weektwo

import (
	"bufio"
	"io"
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
	var isIncreasing bool

	switch {
	case s[0] == s[1]:
		// must differ by at least one
		return false
	case s[0] < s[1]:
		isIncreasing = true
	case s[0] > s[1]:
		isIncreasing = false
	default:
		panic("Unknown state")
	}

	lastNum := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] == lastNum {
			return false
		}

		if abs(s[i]-lastNum) > 3 {
			return false
		}

		if isIncreasing && s[i] < lastNum {
			return false
		}

		if !isIncreasing && s[i] > lastNum {
			return false
		}

		// number is ok
		lastNum = s[i]
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

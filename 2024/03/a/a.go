package a

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func RunA(f io.Reader) int {
	sc := bufio.NewScanner(f)

	var inputs []string

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		inputs = append(inputs, t)
	}

	sum := 0
	for _, input := range inputs {
		for _, m := range findMatches(input) {
			sum += m.x * m.y
		}
	}

	return sum
}

type mul struct {
	x, y int
}

func findMatches(input string) []mul {
	r, _ := regexp.Compile("mul\\((\\d+),(\\d+)\\)")

	matches := []mul{}
	for _, m := range r.FindAllStringSubmatch(input, -1) {
		xS, yS := m[1], m[2]

		x, err := strconv.Atoi(xS)
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(yS)
		if err != nil {
			panic(err)
		}

		if x > 999 || y > 999 || x < 1 || y < 1 {
			continue
		}

		matches = append(matches, mul{x: x, y: y})
	}

	return matches
}

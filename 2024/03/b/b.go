package b

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func RunB(f io.Reader) int {
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
	enabled := true

	for _, input := range inputs {
		matches, isStillEnabled := findMatches(input, enabled)
		for _, m := range matches {
			sum += m.x * m.y
		}

		enabled = isStillEnabled
	}

	return sum
}

type mul struct {
	x, y int
}

func findMatches(input string, multipleEnabled bool) ([]mul, bool) {
	r, _ := regexp.Compile("do\\(\\)|don't\\(\\)|mul\\((\\d+),(\\d+)\\)")

	matches := []mul{}

	for _, m := range r.FindAllStringSubmatch(input, -1) {
		if "don't()" == m[0] {
			multipleEnabled = false
			continue
		} else if "do()" == m[0] {
			multipleEnabled = true
			continue
		}

		if multipleEnabled == false {
			continue
		}

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

	return matches, multipleEnabled
}

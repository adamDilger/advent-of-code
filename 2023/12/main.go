package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var matchRegx = regexp.MustCompile(`(\#+)`)
var unknownRegx = regexp.MustCompile(`(\?+)`)

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sum := parseFile(file)
	fmt.Printf("sum: %d\n", sum)
}

func parseFile(f io.Reader) int {
	sc := bufio.NewScanner(f)

	sum := 0

	for sc.Scan() {
		text := sc.Text()
		fields := strings.Fields(text)

		val := fields[0]

		var matches []int
		matchGroupings := strings.Split(fields[1], ",")
		for _, mg := range matchGroupings {
			if mgInt, err := strconv.Atoi(mg); err != nil {
				panic(err)
			} else {
				matches = append(matches, mgInt)
			}
		}

		sum += calculateReport(val, matches)
	}

	return sum
}

func calculateReport(val string, matches []int) int {
	unknownPositions := unknownRegx.FindAllStringIndex(val, -1)

	curr := []string{val}

	for _, u := range unknownPositions {
		// println("--------------------------------------------------------------------------------------------------------------------")
		next := []string{}

		for _, cString := range curr {
			c := []rune(cString)
			// swap out all curr with replacements
			// fmt.Printf("u: %d, c: %s\n", u, string(c))

			options := generateAllUnknowns(u[1] - u[0])
			for _, o := range options {
				for i := 0; i < len(o); i++ {
					c[i+u[0]] = o[i]
				}

				next = append(next, string(c))
			}
		}

		curr = next
	}

	total := 0
	for _, c := range curr {
		if test(c, matches) {
			// println(c)
			total++
		}
	}
	return total
}

func test(n string, matchGroupings []int) bool {
	matches := matchRegx.FindAllStringSubmatchIndex(string(n), -1)
	if len(matches) != len(matchGroupings) {
		return false
	}

	for i, mg := range matchGroupings {
		if i >= len(matches) {
			return false
		}

		length := matches[i][1] - matches[i][0]

		if mg != length {
			return false
		}
	}

	return true
}

func generateAllUnknowns(n int) [][]rune {
	width := n
	height := int(math.Pow(2, float64(n)))

	var a [][]rune = make([][]rune, height)
	for i := range a {
		a[i] = make([]rune, width)

		for j := range a[i] {
			a[i][j] = '.'
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			k := int(math.Pow(2, float64(j+1)))
			if i%k+1 > k/2 {
				a[i][j] = '#'
			}
		}
	}

	return a
}

/*

.??..??...?##. 1,1,3


// group all unknown

??, ??, ?

then list all available options

for 1

.#
#.
##


for 2

.#
#.
##

for 3

.##
###


*/

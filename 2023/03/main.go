package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type PartNumber struct {
	num       int
	numString string
	x, y      int
	width     int
}

var partNumbers []PartNumber

var isParsingNumber bool
var currentParsedNumber []rune

var symbolMap [][]bool

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(file)

	y := 0

	for sc.Scan() {
		if sc.Err() != nil {
			panic(sc.Err())
		}

		line := sc.Text()
		fmt.Println(line)

		symbolMap = append(symbolMap, make([]bool, len(line)))

		for x, c := range line {
			if unicode.IsNumber(c) {
				isParsingNumber = true
				currentParsedNumber = append(currentParsedNumber, c)
				continue
			}

			// we've hit a non-number, if the previous scan was a number, then add that
			if isParsingNumber {
				addPartNumber(x, y)
			}

			if c != '.' {
				symbolMap[y][x] = true
			}
		}

		// end of line
		if isParsingNumber {
			addPartNumber(len(line), y)
		}

		y++
	}

	var validParts []PartNumber

outer_loop:
	for _, p := range partNumbers {
		for i := max(0, p.x-1); i < min(p.x+p.width+1, len(symbolMap[0])); i++ {
			// above
			if symbolMap[max(0, p.y-1)][i] {
				validParts = append(validParts, p)
				continue outer_loop
			}

			// below
			if p.y+1 != len(symbolMap[p.y]) {
				if symbolMap[p.y+1][i] {
					validParts = append(validParts, p)
					continue outer_loop
				}
			}
		}

		// left
		if p.x != 0 {
			if symbolMap[p.y][p.x-1] {
				validParts = append(validParts, p)
				continue outer_loop
			}
		}

		// right
		if p.x+p.width != len(symbolMap[p.y]) {
			if symbolMap[p.y][p.x+p.width] {
				validParts = append(validParts, p)
				continue outer_loop
			}
		}
	}

	total := 0
	for _, p := range validParts {
		total += p.num
	}
	fmt.Println(total)
}

func addPartNumber(x, y int) {
	width := len(currentParsedNumber)

	num, _ := strconv.Atoi(string(currentParsedNumber))

	partNumbers = append(partNumbers, PartNumber{
		num:       num,
		numString: string(currentParsedNumber),
		width:     width,
		x:         x - width,
		y:         y,
	})

	currentParsedNumber = []rune{}
	isParsingNumber = false
}

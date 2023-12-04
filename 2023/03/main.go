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

type Symbol struct {
	x, y int
	gear bool

	matches map[int]bool
}

var partNumbers []PartNumber

var isParsingNumber bool
var currentParsedNumber []rune

var symbols []Symbol
var symbolMap [][]bool

var partNumberMap [][]int

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
		partNumberMap = append(partNumberMap, make([]int, len(line)))

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
				symbols = append(symbols, Symbol{x: x, y: y, gear: c == '*', matches: make(map[int]bool)})
				symbolMap[y][x] = true
			}
		}

		// end of line
		if isParsingNumber {
			addPartNumber(len(line), y)
		}

		y++
	}

	var validSymbols []Symbol

	maxWidthIndex := len(partNumberMap[0]) - 1

	for _, s := range symbols {
		// above
		if s.y != 0 {
			for x := max(0, s.x-1); x <= min(maxWidthIndex, s.x+1); x++ {
				pn := partNumberMap[s.y-1][x]
				if pn != 0 {
					s.matches[pn] = true
				}
			}
		}

		// below
		if s.y < maxWidthIndex {
			for x := max(0, s.x-1); x <= min(maxWidthIndex, s.x+1); x++ {
				pn := partNumberMap[s.y+1][x]
				if pn != 0 {
					s.matches[pn] = true
				}
			}
		}

		// left
		if s.x != 0 {
			pn := partNumberMap[s.y][s.x-1]
			if pn != 0 {
				s.matches[pn] = true
			}
		}

		// right
		if s.x < maxWidthIndex {
			pn := partNumberMap[s.y][s.x+1]
			if pn != 0 {
				s.matches[pn] = true
			}
		}

		if len(s.matches) > 0 {
			validSymbols = append(validSymbols, s)
		}

		// fmt.Printf("[%d:%d]: %v\n", s.x, s.y, s.matches)
	}

	total := 0
	for _, s := range validSymbols {
		for m := range s.matches {
			total += m
		}
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

	for i := x - width; i < x; i++ {
		partNumberMap[y][i] = num
	}
}

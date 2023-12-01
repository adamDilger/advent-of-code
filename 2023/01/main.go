package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	total := 0

	for sc.Scan() {
		var firstNum, lastNum int
		t := sc.Text()
		line := []rune(t)

		for i, c := range line {
			if unicode.IsNumber(c) {
				firstNum, _ = strconv.Atoi(string(c))
				break
			}

			if wordNum := isWord(line[i:]); wordNum != 0 {
				firstNum = wordNum
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(line[i]) {
				lastNum, _ = strconv.Atoi(string(line[i]))
				break
			}

			if wordNum := isWord(line[i:]); wordNum != 0 {
				lastNum = wordNum
				break
			}
		}

		fmt.Printf("%d:%d - %s\n", firstNum, lastNum, t)
		total += (firstNum * 10) + lastNum
	}

	fmt.Printf("total: %d\n", total)
}

var words = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func isWord(in []rune) int {
word_loop:
	for i, word := range words {
		// check each word for a match
		if len(word) > len(in) {
			continue
		}

		for wi, wc := range word {
			if wc != in[wi] {
				continue word_loop
			}
		}

		return i + 1
	}

	return 0
}

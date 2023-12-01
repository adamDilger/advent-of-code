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

		for _, c := range line {
			if unicode.IsNumber(c) {
				firstNum, _ = strconv.Atoi(string(c))
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(line[i]) {
				lastNum, _ = strconv.Atoi(string(line[i]))
				break
			}
		}

		fmt.Printf("%d:%d - %s\n", firstNum, lastNum, t)
		total += (firstNum * 10) + lastNum
	}

	fmt.Printf("total: %d\n", total)
}

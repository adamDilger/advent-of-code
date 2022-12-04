package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func NewRange(input string) Range {
	splits := strings.Split(input, "-")
	s, _ := strconv.Atoi(splits[0])
	e, _ := strconv.Atoi(splits[1])

	return Range{start: s, end: e}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		println("------------")
		println(line)

		splits := strings.Split(line, ",")

		pair1 := NewRange(splits[0])
		pair2 := NewRange(splits[1])

		println(pair1.start, pair1.end)
		println(pair2.start, pair2.end)

		if pair1.start >= pair2.start && pair1.end <= pair2.end ||
			pair2.start >= pair1.start && pair2.end <= pair1.end {

			count++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	println(count)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	winningNumbers []int
	numbers        []int

	matches map[int]bool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var games []Game

	for sc.Scan() {
		g := Game{matches: make(map[int]bool)}

		line := sc.Text()

		left := strings.Trim(line[strings.Index(line, ":"):strings.Index(line, "|")], ": ")
		right := strings.Trim(line[strings.Index(line, "|")+2:], "| ")

		println(line)
		println(left)
		println(right)

		g.winningNumbers = parseNumberList(left)
		g.numbers = parseNumberList(right)

		games = append(games, g)
	}

	// find all matches
	for _, g := range games {
	outer_loop:
		for _, n := range g.numbers {
			for _, wn := range g.winningNumbers {
				if wn == n {
					g.matches[wn] = true
					continue outer_loop
				}
			}
		}
	}

	total := 0
	for _, g := range games {
		m := len(g.matches)
		if m == 0 {
			continue
		}

		points := 1
		for i := 1; i < m; i++ {
			points = (points * 2)
		}

		fmt.Println(points, g)
		total += points
	}
	fmt.Println(total)
}

func parseNumberList(in string) []int {
	out := []int{}

	for _, nString := range strings.Split(in, " ") {
		if nString == "" {
			continue
		}

		wn, err := strconv.Atoi(nString)
		if err != nil {
			panic(err)
		}
		out = append(out, wn)
	}

	return out
}

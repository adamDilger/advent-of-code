package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id      int
	subsets []SubSet
}

type SubSet struct {
	red   int
	blue  int
	green int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(file)

	var games []Game

	for sc.Scan() {
		if sc.Err() != nil {
			panic(sc.Err())
		}

		line := sc.Text()
		println(line)

		gameId, _ := strconv.Atoi(line[len("Game "):strings.Index(line, ":")])
		g := Game{id: gameId}

		for _, sets := range strings.Split(line[strings.Index(line, ":")+1:], ";") {

			var ss SubSet
			for _, set := range strings.Split(strings.Trim(sets, " "), ", ") {
				var amount int
				var color string
				fmt.Sscanf(set, "%d %s", &amount, &color)

				switch color {
				case "red":
					ss.red = amount
				case "blue":
					ss.blue = amount
				case "green":
					ss.green = amount
				}
			}

			g.subsets = append(g.subsets, ss)
		}

		games = append(games, g)
	}

	total_sum := 0

	for _, g := range games {
		total := 1

		red_min := 0
		green_min := 0
		blue_min := 0

		for _, ss := range g.subsets {
			if ss.red > red_min {
				red_min = ss.red
			}

			if ss.blue > blue_min {
				blue_min = ss.blue
			}

			if ss.green > green_min {
				green_min = ss.green
			}
		}

		if red_min > 0 {
			total *= red_min
		}

		if blue_min > 0 {
			total *= blue_min
		}
		if green_min > 0 {
			total *= green_min
		}

		fmt.Printf("sum: %d\n", total)

		total_sum += total
	}

	fmt.Printf("total: %d\n", total_sum)
}

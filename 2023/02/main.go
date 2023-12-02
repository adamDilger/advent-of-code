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

	// The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
	red_max := 12
	green_max := 13
	blue_max := 14

	total_sum := 0

outer_loop:
	for _, g := range games {
		for _, ss := range g.subsets {
			if ss.green > green_max || ss.red > red_max || ss.blue > blue_max {
				continue outer_loop
			}
		}

		total_sum += g.id
	}

	fmt.Printf("sum: %d\n", total_sum)
}

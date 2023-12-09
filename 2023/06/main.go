package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time, record int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	total := 1

	race := parseRaces(file)
	fmt.Println(race)

	// dist = (raceTime - btnTime) * btnTime

	isEven := race.time%2 == 0

	mid := (race.time / 2) + 1

	for btnTime := mid; btnTime <= race.time; btnTime++ {
		dist := (race.time - btnTime) * btnTime
		// fmt.Printf("btn: %d | dist: %d\n", btnTime, dist)

		if dist <= race.record {
			var count int
			if isEven {
				count = ((btnTime - mid) * 2) + 1
			} else {
				count = ((btnTime - mid) * 2)
			}

			total *= count
			fmt.Printf("FOUND: %d\n", count)
			break
		}
	}

	fmt.Println(total)
}

func parseRaces(f *os.File) Race {
	var race Race

	sc := bufio.NewScanner(f)
	sc.Scan()
	_, line, _ := strings.Cut(sc.Text(), "Time: ")

	if timeNum, err := strconv.Atoi(strings.ReplaceAll(line, " ", "")); err == nil {
		race.time = timeNum
	} else {
		panic(err)
	}

	sc.Scan()
	_, line, _ = strings.Cut(sc.Text(), "Distance: ")
	if distNum, err := strconv.Atoi(strings.ReplaceAll(line, " ", "")); err == nil {
		race.record = distNum
	} else {
		panic(err)
	}

	return race
}

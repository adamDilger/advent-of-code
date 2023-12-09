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

	races := parseRaces(file)

	total := 1

	for _, race := range races {
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
	}

	fmt.Println(total)
}

func parseRaces(f *os.File) []Race {
	var races []Race

	sc := bufio.NewScanner(f)
	sc.Scan()
	for _, time := range strings.Fields(sc.Text())[1:] {
		if timeNum, err := strconv.Atoi(time); err == nil {
			races = append(races, Race{time: timeNum})
		} else {
			panic(err)
		}
	}

	sc.Scan()
	for i, dist := range strings.Fields(sc.Text())[1:] {
		if distNum, err := strconv.Atoi(dist); err == nil {
			races[i].record = distNum
		} else {
			panic(err)
		}
	}

	return races
}

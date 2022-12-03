package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentMaxes := []int{0, 0, 0}

	var currentTotal int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for i, m := range currentMaxes {
				if currentTotal > m {
					currentMaxes[i] = currentTotal

					sort.Ints(currentMaxes)
					break
				}
			}

			currentTotal = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to parse line: %v", err)
		}

		currentTotal += cal
	}

	println(sum(currentMaxes))

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}
}

func sum(inputs []int) int {
	var t int
	for _, v := range inputs {
		t += v
	}
	return t
}

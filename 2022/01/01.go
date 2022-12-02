package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currentTotal int
	var currentMax int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currentTotal = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to parse line: %v", err)
		}

		currentTotal += cal

		if currentTotal > currentMax {
			currentMax = currentTotal
		}
	}

	println(currentMax)

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}
}

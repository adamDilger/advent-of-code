package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// "soil-to-fert" : map[seed-int]fert-int
var valueMap = make(map[string]map[int]int)

var currentKey string

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(file)

	sc.Scan()
	_, seedLine, _ := strings.Cut(sc.Text(), "seeds: ")
	seeds := parseIntList(seedLine)

	sc.Scan()

	firstLine := true
	// var dest, source string

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			firstLine = true
			continue
		}

		if firstLine {
			firstLine = false
			currentKey = string(line[0:strings.Index(line, " ")])
			valueMap[currentKey] = make(map[int]int)

			continue
		}

		ints := parseIntList(line)
		dest := ints[0]
		source := ints[1]
		length := ints[2]

		for i := 0; i < length; i++ {
			valueMap[currentKey][source+i] = dest + i
		}
	}

	// Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
	for _, seed := range seeds {
		soil := getValueFromMap(valueMap, "seed-to-soil", seed)
		fert := getValueFromMap(valueMap, "soil-to-fertilizer", soil)
		water := getValueFromMap(valueMap, "fertilizer-to-water", fert)
		light := getValueFromMap(valueMap, "water-to-light", water)
		temperature := getValueFromMap(valueMap, "light-to-temperature", light)
		humidity := getValueFromMap(valueMap, "temperature-to-humidity", temperature)
		location := getValueFromMap(valueMap, "humidity-to-location", humidity)

		fmt.Printf("Seed %d, soil %d, fertilizer %d, water %d, light %d, temperature %d, humidity %d, location %d\n", seed, soil, fert, water, light, temperature, humidity, location)
	}
}

func getValueFromMap(in map[string]map[int]int, key string, index int) int {
	if val, ok := in[key][index]; ok {
		return val
	}

	return index
}

func parseIntList(s string) []int {
	stringNumbers := strings.Split(s, " ")

	numbers := []int{}
	for _, i := range stringNumbers {
		if i == "" {
			continue
		}

		iInt, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, iInt)
	}

	return numbers
}

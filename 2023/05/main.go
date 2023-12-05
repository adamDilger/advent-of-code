package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeOffset struct {
	seedStart, seedEnd int
	offset             int
}

type SeedRange struct {
	start, end int
}

var rangeOffsetMap = make(map[string][]RangeOffset)

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
	seedRanges := parseSeedList(seedLine)
	fmt.Println(seedRanges)

	sc.Scan()

	firstLine := true

	for sc.Scan() {
		line := sc.Text()
		// println(line)

		if line == "" {
			firstLine = true
			continue
		}

		if firstLine {
			firstLine = false
			currentKey = string(line[0:strings.Index(line, " ")])
			rangeOffsetMap[currentKey] = []RangeOffset{}

			continue
		}

		ints := parseIntList(line)
		dest := ints[0]
		source := ints[1]
		length := ints[2]

		rangeOffsetMap[currentKey] = append(rangeOffsetMap[currentKey], RangeOffset{seedStart: source, seedEnd: source + length, offset: dest - source})
	}

	minSeed := 0
	minLocation := int(^uint(0) >> 1)

	for _, r := range seedRanges {
		fmt.Printf("Seed %d: %d\n", r.start, r.end)

		seed := r.start

		for seed <= r.end {
			cl := calculateLocation(seed)

			if cl < minLocation {
				minLocation = cl
				minSeed = seed
			}

			// if it ain't broke...
			if cl+1000 == calculateLocation(seed+1000) {
				seed += 1000
			} else {
				seed++
			}
		}
	}

	println(minSeed)
	println(minLocation)
}

func getValueFromRanges(in map[string][]RangeOffset, key string, seed int) int {
	// loop over all ranges for seed num, find the range, calculate the number from the given offset, otherwise return seed num
	for _, r := range in[key] {
		if r.seedStart <= seed && seed < r.seedEnd {
			val := seed + r.offset
			// fmt.Printf("%s: %d %d\n", key, seed, r.offset)

			return val
		}
	}

	return seed
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

func parseSeedList(s string) []SeedRange {
	values := parseIntList(s)

	out := []SeedRange{}

	for i := 0; i < len(values); i += 2 {
		out = append(out, SeedRange{start: values[i], end: values[i] + values[i+1]})
	}

	return out
}

var count int

func calculateLocation(seed int) int {
	soil := getValueFromRanges(rangeOffsetMap, "seed-to-soil", seed)
	fert := getValueFromRanges(rangeOffsetMap, "soil-to-fertilizer", soil)
	water := getValueFromRanges(rangeOffsetMap, "fertilizer-to-water", fert)
	light := getValueFromRanges(rangeOffsetMap, "water-to-light", water)
	temperature := getValueFromRanges(rangeOffsetMap, "light-to-temperature", light)
	humidity := getValueFromRanges(rangeOffsetMap, "temperature-to-humidity", temperature)
	location := getValueFromRanges(rangeOffsetMap, "humidity-to-location", humidity)

	// fmt.Printf("Seed %d, soil %d, fertilizer %d, water %d, light %d, temperature %d, humidity %d, location %d\n", seed, soil, fert, water, light, temperature, humidity, location)
	// fmt.Printf("Seed %d, location %d\n", seed, location)

	return location
}

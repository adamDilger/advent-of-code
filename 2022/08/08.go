package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		println(lines[len(lines)-1])
	}
	println()

	width := len(lines[0])
	height := len(lines)

	count := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if isVisible(x, y, lines) {
				count++
			}
		}

		// if y == 2 {
		// 	break
		// }
	}

	println(count)
}

func isVisible(x, y int, lines []string) bool {
	char, _ := strconv.Atoi(string(lines[y][x]))
	// println(char)

	obstructed := false

	for i := y - 1; i >= 0; i-- {
		char1, _ := strconv.Atoi(string(lines[i][x]))
		if char1 >= char {
			// println("FUK", char, char1)
			obstructed = true
			break
		} else {
			// println("NO", char, char1)
		}
	}

	if obstructed == false {
		// println(x, y, char)
		return true
	}

	obstructed = false

	// down
	for i := y + 1; i < len(lines); i++ {
		char1, _ := strconv.Atoi(string(lines[i][x]))
		if char1 >= char {
			// println("FUK", char, char1)
			obstructed = true
			break
		} else {
			// println("NO", char, char1)
		}
	}

	if obstructed == false {
		// println(x, y, char)
		return true
	}

	obstructed = false

	// left
	for i := x - 1; i >= 0; i-- {
		char1, _ := strconv.Atoi(string(lines[y][i]))
		if char1 >= char {
			// println("FUK", char, char1)
			obstructed = true
			break
		} else {
			// println("NO", char, char1)
		}
	}

	if obstructed == false {
		// println(x, y, char)
		return true
	}

	obstructed = false

	// right
	for i := x + 1; i < len(lines[0]); i++ {
		char1, _ := strconv.Atoi(string(lines[y][i]))
		if char1 >= char {
			// println("FUK", char, char1)
			obstructed = true
			break
		} else {
			// println("NO", char, char1)
		}
	}

	if obstructed == false {
		// println(x, y, char)
		return true
	}

	obstructed = false

	return false
}

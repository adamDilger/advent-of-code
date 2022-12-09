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

	highest := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			score := isVisible(x, y, lines)
			if score > highest {
				highest = score
			}
			// println("SCORE", score)
		}

		// if y == 2 {
		// 	break
		// }
	}

	println(highest)
}

func isVisible(x, y int, lines []string) int {
	char, _ := strconv.Atoi(string(lines[y][x]))
	// println(char)

	var up, down, left, right int

	for i := y - 1; i >= 0; i-- {
		up++
		char1, _ := strconv.Atoi(string(lines[i][x]))
		if char1 >= char {
			// println("FUK", char, char1)
			break
		} else {
		}
	}

	// down
	for i := y + 1; i < len(lines); i++ {
		down++
		char1, _ := strconv.Atoi(string(lines[i][x]))
		if char1 >= char {
			// println("FUK", char, char1)
			break
		} else {
		}
	}

	// left
	for i := x - 1; i >= 0; i-- {
		left++
		char1, _ := strconv.Atoi(string(lines[y][i]))
		if char1 >= char {
			// println("FUK", char, char1)
			break
		} else {
		}
	}

	// right
	for i := x + 1; i < len(lines[0]); i++ {
		right++
		char1, _ := strconv.Atoi(string(lines[y][i]))
		if char1 >= char {
			// println("FUK", char, char1)
			break
		} else {
		}
	}

	return up * down * left * right
}

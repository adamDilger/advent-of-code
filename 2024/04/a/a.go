package a

import (
	"bufio"
	"io"
)

func RunA(f io.Reader) int {
	sc := bufio.NewScanner(f)

	grid := Grid{}

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		grid = append(grid, []rune(t))
	}

	sum := 0
	for y, l := range grid {
		for x := range l {
			sum += search(grid, x, y)
		}
	}

	return sum
}

func getNextLetter(l rune) rune {
	switch l {
	case 'X':
		return 'M'
	case 'M':
		return 'A'
	case 'A':
		return 'S'
	}

	return ' '
}

type XY struct {
	x, y int
}

func search(g Grid, x, y int) int {
	if g.get(x, y) != 'X' {
		return 0
	}

	next := getNextLetter('X')
	if next == ' ' {
		panic("Invalid")
	}

	// look for all adjacent letters that match
	positions := [][2]int{
		{-1, -1}, // top left
		{-1, 0},  // left
		{-1, 1},  // bottom left
		{0, 1},   // bottom
		{0, -1},  // top
		{1, -1},  // top right
		{1, 0},   // right
		{1, 1},   // bottom right
	}

	totalForThisChar := 0

	for _, p := range positions {
		if findDirection(g, x, y, XY{x: p[0], y: p[1]}, next) {
			totalForThisChar += 1
		}
	}

	return totalForThisChar
}

func findDirection(g Grid, x, y int, direction XY, needle rune) bool {
	x += direction.x
	y += direction.y

	if g.get(x, y) != needle {
		return false
	}

	if needle == 'S' {
		return true
	}

	next := getNextLetter(needle)
	if next == ' ' {
		panic("Invalid")
	}

	return findDirection(g, x, y, direction, next)
}

type Grid [][]rune

func (g Grid) get(x, y int) rune {
	if x < 0 || y < 0 || y >= len(g) || x >= len(g[y]) {
		return ' '
	}

	return g[y][x]
}

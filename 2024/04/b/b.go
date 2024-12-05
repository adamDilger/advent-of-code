package b

import (
	"bufio"
	"io"
)

func RunB(f io.Reader) int {
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

type XY struct {
	x, y int
}

type DIR int

const (
	T DIR = iota
	B
	L
	R
)

func (d DIR) opposite() DIR {
	switch d {
	case T:
		return B
	case B:
		return T
	case L:
		return R
	case R:
		return L
	}
	panic("Invalid DIR")
}

func search(g Grid, x, y int) int {
	if g.get(x, y) != 'A' {
		return 0
	}

	// look for all adjacent letters that match
	positions := []DIR{T, B, L, R}

	totalForThisChar := 0

	for _, p := range positions {
		if findDirection(g, x, y, p, 'M') {
			totalForThisChar += 1
		}
	}

	return totalForThisChar
}

func findDirection(g Grid, x, y int, direction DIR, needle rune) bool {
	x1, x2, y1, y2 := 0, 0, 0, 0

	switch direction {
	case T:
		{
			x1 = x - 1
			x2 = x + 1

			y1 = y - 1
			y2 = y - 1
		}
	case B:
		{
			x1 = x - 1
			x2 = x + 1

			y1 = y + 1
			y2 = y + 1
		}
	case L:
		{
			x1 = x - 1
			x2 = x - 1

			y1 = y - 1
			y2 = y + 1
		}
	case R:
		{
			x1 = x + 1
			x2 = x + 1

			y1 = y - 1
			y2 = y + 1
		}
	}

	if g.get(x1, y1) != needle || g.get(x2, y2) != needle {
		return false
	}

	if needle == 'S' {
		return true
	}

	return findDirection(g, x, y, direction.opposite(), 'S')
}

type Grid [][]rune

func (g Grid) get(x, y int) rune {
	if x < 0 || y < 0 || y >= len(g) || x >= len(g[y]) {
		return ' '
	}

	return g[y][x]
}

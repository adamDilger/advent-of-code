package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// file, err := os.Open("test1.txt")
	// file, err := os.Open("test2.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	grid := parseFile(file)

	fmt.Println(grid)

	x, y := grid.startX, grid.startY

	seen := make(map[int]map[int]bool)
	for i := 0; i < len(grid.coords); i++ {
		seen[i] = make(map[int]bool)
	}

	seen[x][y] = true

	path := findPath(x, y, &grid, seen, []Coordinate{})
	for _, c := range path {
		fmt.Println(c)
	}

	fmt.Println("done", (len(path)/2)+1)
}

func findPath(x, y int, g *Grid, seenNodes map[int]map[int]bool, path []Coordinate) []Coordinate {
	// base case
	if g.Get(x, y).tile == 0 || g.Get(x, y).tile == Empty {
		return path
	}

	dirs := [][]int{}

	curr := g.Get(x, y)

	if curr.tile == Start ||
		curr.tile == TopLeft ||
		curr.tile == TopRight ||
		curr.tile == Vertical {
		dirs = append(dirs, []int{0, 1}) // down
	}

	if curr.tile == Start ||
		curr.tile == BotLeft ||
		curr.tile == BotRight ||
		curr.tile == Vertical {
		dirs = append(dirs, []int{0, -1}) // up
	}

	if curr.tile == Start ||
		curr.tile == BotLeft ||
		curr.tile == TopLeft ||
		curr.tile == Horizontal {
		dirs = append(dirs, []int{1, 0}) // right
	}

	if curr.tile == Start ||
		curr.tile == BotRight ||
		curr.tile == TopRight ||
		curr.tile == Horizontal {
		dirs = append(dirs, []int{-1, 0}) // left
	}

	for _, dir := range dirs {
		newX := x + dir[0]
		newY := y + dir[1]

		if newX < 0 || newY < 0 {
			continue
		}

		if newX >= len(g.coords[0]) || newY >= len(g.coords) {
			continue
		}

		if _, ok := seenNodes[newX][newY]; ok {
			continue // seen this node already
		}

		seenNodes[newX][newY] = true

		c := g.Get(newX, newY)
		if c.tile == Empty {
			continue
		}

		path = append(path, c)

		return findPath(newX, newY, g, seenNodes, path)
	}

	// should there be some path.pop() somewhere?

	return path
}

type Coordinate struct {
	x, y int
	tile Tile
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d) %c", c.x, c.y, c.tile)
}

type Tile rune

const (
	Empty      Tile = '.'
	TopLeft    Tile = 'F'
	TopRight   Tile = '7'
	BotLeft    Tile = 'L'
	BotRight   Tile = 'J'
	Vertical   Tile = '|'
	Horizontal Tile = '-'

	Start Tile = 'S'
)

/*

.....
.S-7.
.|.|.
.L-J.
.....

*/

type Grid struct {
	startX, startY int
	coords         [][]Coordinate
}

func (g Grid) Get(x, y int) Coordinate {
	if len(g.coords) <= y {
		return Coordinate{tile: Empty}
	}

	if len(g.coords[0]) <= x {
		return Coordinate{tile: Empty}
	}

	return g.coords[y][x]
}

func (g Grid) String() string {
	sb := strings.Builder{}

	for _, row := range g.coords {
		for _, c := range row {
			sb.WriteString(string(c.tile))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func parseFile(f io.Reader) Grid {
	var grid Grid

	sc := bufio.NewScanner(f)

	y := 0
	for sc.Scan() {
		line := sc.Text()
		row := make([]Coordinate, len(line))

		for i, c := range line {
			cc := Coordinate{tile: Tile(c), x: i, y: y}
			if cc.tile == Start {
				grid.startX = i
				grid.startY = y
			}

			row[i] = cc
		}

		grid.coords = append(grid.coords, row)
		y++
	}

	return grid
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const debug = true

const (
	START = 'S'
	END   = 'E'
)

func log(args ...any) {
	if debug {
		fmt.Println(args...)
	}
}

type Point struct {
	x, y int
	key  string
}

func (p *Point) getKey() string {
	if p.key == "" {
		p.key = fmt.Sprintf("%d:%d", p.x, p.y)
	}
	return p.key
}

func (p *Point) left() Point {
	return Point{x: p.x - 1, y: p.y}
}

func (p *Point) right() Point {
	return Point{x: p.x + 1, y: p.y}
}

func (p *Point) up() Point {
	return Point{x: p.x, y: p.y - 1}
}

func (p *Point) down() Point {
	return Point{x: p.x, y: p.y + 1}
}

type Grid struct {
	X, Y int
	Grid [][]rune
}

// copy a map
func copyMap(input map[string]struct{}) map[string]struct{} {
	c := make(map[string]struct{})
	for k, v := range input {
		c[k] = v
	}
	return c
}

func (g *Grid) printGrid() {
	for y := range g.Grid {
		for x := range g.Grid[y] {
			print(string(g.Grid[y][x]))
		}

		println()
	}
}

func (g *Grid) getStart() Point {
	for y := range g.Grid {
		for x, c := range g.Grid[y] {
			if c == START {
				return Point{x: x, y: y}
			}
		}
	}

	panic("No start found")
}

func elev(i rune) rune {
	if i == START {
		return 'a'
	} else if i == END {
		return 'z'
	}

	return i
}

func (g *Grid) search(prev, curr Point, pathLength int, visited map[string]struct{}) int {
	if curr.y < 0 || curr.y == len(g.Grid) {
		return 0
	}

	if curr.x < 0 || curr.x == len(g.Grid[curr.y]) {
		return 0
	}

	// point is in grid
	char := g.Grid[curr.y][curr.x]
	prevChar := g.Grid[prev.y][prev.x]

	charElev := elev(char)
	prevCharElev := elev(prevChar)

	if charElev-1 > prevCharElev {
		// too high
		return 0
	}

	log("FUK", curr.x, curr.y, prev.x, prev.y)
	if _, ok := visited[curr.getKey()]; ok {
		return 0
	}

	pathLength++
	visited[curr.getKey()] = struct{}{}

	if char == END {
		fmt.Println("END found", pathLength)
		return pathLength
	}

	foundRight := g.search(curr, curr.right(), pathLength, visited)
	foundDown := g.search(curr, curr.down(), pathLength, visited)
	foundLeft := g.search(curr, curr.left(), pathLength, visited)
	foundUp := g.search(curr, curr.up(), pathLength, visited)

	if foundRight+foundDown+foundLeft+foundUp == 0 {
		return 0
	}

	res := []int{}
	if foundRight != 0 {
		res = append(res, foundRight)
	}
	if foundLeft != 0 {
		res = append(res, foundLeft)
	}
	if foundDown != 0 {
		res = append(res, foundDown)
	}
	if foundUp != 0 {
		res = append(res, foundUp)
	}

	sort.Ints(res)

	return res[0]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var grid Grid

	for scanner.Scan() {
		line := scanner.Text()
		log(line)

		grid.Grid = append(grid.Grid, []rune(line))
	}

	startPoint := grid.getStart()

	visited := make(map[string]struct{})
	log(grid.search(startPoint, startPoint.down(), 0, visited))
}

package main

import (
	"bufio"
	"fmt"
	"os"
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
	X, Y  int
	Grid  [][]rune
	Nodes map[string]*Node
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

func (g *Grid) newNode(point Point) *Node {
	x, y := point.x, point.y
	char := g.Grid[y][x]

	n := &Node{
		Point: point,
		Value: char,
	}
	n.Edges = make(map[string]int)

	if char == START {
		n.Start = true
		char = 'a'
	} else if char == END {
		n.End = true
		char = 'z'
	}

	checkEdge := func(p Point) {
		nextChar := g.Grid[p.y][p.x]

		var weight int

		if nextChar < char {
			weight = int(char-nextChar) + 2
		} else if nextChar == char {
			weight = 2
		} else if nextChar == char+1 {
			weight = 1
		}

		if weight != 0 {
			n.Edges[p.getKey()] = weight
		}
	}

	if y != 0 {
		checkEdge(point.up())
	}

	if y < len(g.Grid)-2 {
		checkEdge(point.down())
	}

	if x != 0 {
		checkEdge(point.left())
	}

	if x < len(g.Grid[y])-2 {
		checkEdge(point.right())
	}

	return n
}

type Node struct {
	Point
	Value      rune
	Start, End bool
	Edges      map[string]int
}

type Graph struct {
	Nodes map[string]Node
}

func elev(i rune) rune {
	if i == START {
		return 'a'
	} else if i == END {
		return 'z'
	}

	return i
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var grid Grid
	grid.Nodes = make(map[string]*Node)

	for scanner.Scan() {
		line := scanner.Text()
		log(line)

		grid.Grid = append(grid.Grid, []rune(line))
	}

	for y := range grid.Grid {
		for x := range grid.Grid[y] {
			n := grid.newNode(Point{x: x, y: y})
			grid.Nodes[n.getKey()] = n
		}
	}

	sp := grid.getStart()
	fmt.Println(grid.Nodes[sp.getKey()])
}

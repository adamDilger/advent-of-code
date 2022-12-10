package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	HEAD  = 'H'
	TAIL  = 'N'
	START = 's'
	UP    = 'U'
	DOWN  = 'D'
	LEFT  = 'L'
	RIGHT = 'R'

	WIDTH  = 100
	HEIGHT = 100
)

type Move struct {
	Direction rune
	Value     int
}

type Point struct {
	x, y int
}

type Game struct {
	Head, Tail Point
	Moves      []Move

	TailPositions map[string]struct{}
}

func (g *Game) logTail() {
	if g.TailPositions == nil {
		g.TailPositions = make(map[string]struct{})
	}

	pos := fmt.Sprintf("%d,%d", g.Tail.x, g.Tail.y)
	g.TailPositions[pos] = struct{}{}
}

func (g *Game) printMap() {
	// get the max point of both head and tail, and print a loop

	maxX := g.Head.x
	if g.Tail.x > g.Head.x {
		maxX = g.Tail.x
	}

	maxY := g.Head.y
	if g.Tail.y > g.Head.y {
		maxY = g.Tail.y
	}

	maxX += 4
	maxY += 3

	for x := 0; x <= maxX; x++ {
		if x == 0 {
			print("  ")
		}
		print(x)
	}

	println()

	for y := maxY; y >= 0; y-- {
		fmt.Printf("%2d", y)
		for x := 0; x <= maxX; x++ {
			if g.Head.x == x && g.Head.y == y {
				print("H")
			} else if g.Tail.x == x && g.Tail.y == y {
				print("T")
			} else {
				print(".")
			}
		}
		println()
	}
}

func (g *Game) alignX() {
	if g.Tail.x < g.Head.x {
		g.Tail.x++
	} else if g.Tail.x > g.Head.x {
		g.Tail.x--
	}
}

func (g *Game) alignY() {
	if g.Tail.y < g.Head.y {
		g.Tail.y++
	} else if g.Tail.y > g.Head.y {
		g.Tail.y--
	}
}

func (g *Game) up() {
	g.Head.y++

	if g.Tail.y+1 == g.Head.y {
		return
	}

	if g.Tail.y+2 == g.Head.y {
		g.Tail.y++
		g.alignX()
	}
}

func (g *Game) down() {
	g.Head.y--

	if g.Tail.y-1 == g.Head.y {
		return
	}

	if g.Tail.y-2 == g.Head.y {
		g.Tail.y--
		g.alignX()
	}
}

func (g *Game) left() {
	g.Head.x--

	if g.Tail.x-1 == g.Head.x {
		return
	}

	if g.Tail.x-2 == g.Head.x {
		g.Tail.x--
		g.alignY()
	}
}

func (g *Game) right() {
	g.Head.x++

	if g.Tail.x+1 == g.Head.x {
		return
	}

	if g.Tail.x+2 == g.Head.x {
		g.Tail.x++
		g.alignY()
	}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var game Game
	scanner := bufio.NewScanner(input)

	game.Head = Point{x: 0, y: 0}
	game.Tail = Point{x: 0, y: 0}

	for scanner.Scan() {
		var dir string
		var val int
		line := scanner.Text()
		fmt.Fscanf(strings.NewReader(line), "%s %d", &dir, &val)
		m := Move{Direction: rune(dir[0]), Value: val}
		game.Moves = append(game.Moves, m)
	}

	for _, m := range game.Moves {
		// fmt.Printf("\n%c %d\n", m.Direction, m.Value)

		var fn func()

		switch m.Direction {
		case UP:
			fn = game.up
		case DOWN:
			fn = game.down
		case LEFT:
			fn = game.left
		case RIGHT:
			fn = game.right
		}

		for i := 0; i < m.Value; i++ {
			fn()
			game.logTail()
		}

		// game.printMap()
	}

	println(len(game.TailPositions))
}

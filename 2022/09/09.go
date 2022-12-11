package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	HEAD  = 'H'
	TAIL  = 'T'
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
	Points [10]Point
	Moves  []Move

	TailPositions map[string]struct{}
}

func (g *Game) logTail() {
	if g.TailPositions == nil {
		g.TailPositions = make(map[string]struct{})
	}

	tailPoint := g.Points[len(g.Points)-1]

	pos := fmt.Sprintf("%d,%d", tailPoint.x, tailPoint.y)
	g.TailPositions[pos] = struct{}{}
}

func (g *Game) printMap() {
	// get the max point of both head and tail, and print a loop

	maxX := 0
	for _, x := range g.Points {
		if x.x > maxX {
			maxX = x.x
		}
	}

	maxY := 0
	for _, y := range g.Points {
		if y.y > maxY {
			maxY = y.y
		}
	}

	maxX += 5
	maxY += 5

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
			m := false

			for pi, headPoint := range g.Points {
				if headPoint.x == x && headPoint.y == y {
					switch pi {
					case 0:
						print(string(HEAD))
					case len(g.Points) - 1:
						print(string(TAIL))
					default:
						print(pi)
					}

					m = true
					break
				}
			}

			if !m {
				print(".")
			}
		}
		println()
	}
}

func (g *Game) adjustY(index int) {
	headPoint := &g.Points[index-1]
	tailPoint := &g.Points[index]

	if headPoint.y > tailPoint.y {
		tailPoint.y++
	} else if headPoint.y < tailPoint.y {
		tailPoint.y--
	}
}

func (g *Game) adjustX(index int) {
	headPoint := &g.Points[index-1]
	tailPoint := &g.Points[index]

	if headPoint.x > tailPoint.x {
		tailPoint.x++
	} else if headPoint.x < tailPoint.x {
		tailPoint.x--
	}
}

func (g *Game) align(index int) {
	headPoint := &g.Points[index-1]
	tailPoint := &g.Points[index]

	// X
	if headPoint.x < tailPoint.x {
		if tailPoint.x-headPoint.x > 1 {
			tailPoint.x--
			g.adjustY(index)
		}
	}

	if headPoint.x > tailPoint.x {
		if headPoint.x-tailPoint.x > 1 {
			tailPoint.x++
			g.adjustY(index)
		}
	}

	// Y
	if headPoint.y < tailPoint.y {
		if tailPoint.y-headPoint.y > 1 {
			tailPoint.y--
			g.adjustX(index)
		}
	}

	if headPoint.y > tailPoint.y {
		if headPoint.y-tailPoint.y > 1 {
			tailPoint.y++
			g.adjustX(index)
		}
	}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var game Game
	for i := range game.Points {
		game.Points[i].x = 30
		game.Points[i].y = 30
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		var dir string
		var val int
		line := scanner.Text()
		fmt.Fscanf(strings.NewReader(line), "%s %d", &dir, &val)
		m := Move{Direction: rune(dir[0]), Value: val}
		game.Moves = append(game.Moves, m)
	}

	// for mIndex, m := range game.Moves {
	for _, m := range game.Moves {
		// fmt.Printf("\n%c %d\n", m.Direction, m.Value)

		head := &game.Points[0]

		for i := 0; i < m.Value; i++ {
			switch m.Direction {
			case UP:
				head.y++
			case DOWN:
				head.y--
			case LEFT:
				head.x--
			case RIGHT:
				head.x++
			}

			for j := 1; j < len(game.Points); j++ {
				game.align(j)
			}

			game.logTail()
		}

		// if mIndex == 3 {
		// 	break
		// }

		game.printMap()
	}

	println(len(game.TailPositions))
}

package a

import (
	"bufio"
	"fmt"
	"io"
)

type DIR int

const (
	UP DIR = iota
	DOWN
	LEFT
	RIGHT
)

type XY struct {
	x, y int
}

type Guard struct {
	pos XY
	dir DIR
}

func (g Guard) nextPos() XY {
	switch g.dir {
	case UP:
		return XY{x: g.pos.x, y: g.pos.y - 1}
	case DOWN:
		return XY{x: g.pos.x, y: g.pos.y + 1}
	case LEFT:
		return XY{x: g.pos.x - 1, y: g.pos.y}
	case RIGHT:
		return XY{x: g.pos.x + 1, y: g.pos.y}
	}

	panic("Invalid direction")
}

func (g *Guard) rotate() {
	switch g.dir {
	case UP:
		g.dir = RIGHT
	case DOWN:
		g.dir = LEFT
	case LEFT:
		g.dir = UP
	case RIGHT:
		g.dir = DOWN
	default:
		panic("Invalid direction")
	}
}

type Map struct {
	obstaclesByPos map[int]map[int]struct{}
	positionsSeen  map[int]map[int]struct{}
	guard          Guard
	rows, cols     int
}

func (m *Map) addObstacle(pos XY) {
	if _, ok := m.obstaclesByPos[pos.y]; !ok {
		m.obstaclesByPos[pos.y] = make(map[int]struct{})
	}

	m.obstaclesByPos[pos.y][pos.x] = struct{}{}
}

func (m *Map) addSeen(pos XY) {
	if _, ok := m.positionsSeen[pos.y]; !ok {
		m.positionsSeen[pos.y] = make(map[int]struct{})
	}

	m.positionsSeen[pos.y][pos.x] = struct{}{}
}

func (m *Map) get(pos XY) rune {
	if pos.x < 0 || pos.y < 0 ||
		pos.x >= m.cols || pos.y >= m.rows {
		return ' '
	}

	if _, ok := m.obstaclesByPos[pos.y][pos.x]; ok {
		return '#'
	}

	if m.guard.pos.x == pos.x && m.guard.pos.y == pos.y {
		switch m.guard.dir {
		case UP:
			return '^'
		case DOWN:
			return 'v'
		case LEFT:
			return '<'
		case RIGHT:
			return '>'
		}
	}

	if _, ok := m.positionsSeen[pos.y][pos.x]; ok {
		return 'X'
	}

	return '.'
}

func (m *Map) print() {
	fmt.Println("----------------------", m.guard.pos, m.guard.dir)
	for y := range m.rows {
		for x := range m.cols {
			fmt.Print(string(m.get(XY{x: x, y: y})))
		}
		fmt.Println()
	}
	fmt.Println("----------------------")
}

func RunA(f io.Reader) int {
	sc := bufio.NewScanner(f)

	m := Map{
		obstaclesByPos: make(map[int]map[int]struct{}),
		positionsSeen:  make(map[int]map[int]struct{}),
		guard: Guard{
			dir: UP,
			pos: XY{x: 0, y: 0},
		},
	}

	for sc.Scan() {
		t := sc.Text()

		if m.cols == 0 {
			m.cols = len(t)
		}

		if sc.Err() != nil {
			panic(sc.Err())
		}

		for x, xC := range t {
			if xC == '#' {
				m.addObstacle(XY{x: x, y: m.rows})
			} else if xC == '^' {
				m.guard.pos.x = x
				m.guard.pos.y = m.rows
			}
		}

		m.rows++
	}

	m.addSeen(m.guard.pos)

	// m.print()

	for {
		n := m.guard.nextPos()
		nC := m.get(m.guard.nextPos())

		if nC == '#' {
			m.guard.rotate()
			continue
		} else if nC == ' ' {
			break
		}

		m.guard.pos = n
		m.addSeen(m.guard.pos)
		// m.print()
	}

	totalSeen := 0
	for y := range m.rows {
		totalSeen += len(m.positionsSeen[y])
	}

	// 	m.print()
	return totalSeen
}

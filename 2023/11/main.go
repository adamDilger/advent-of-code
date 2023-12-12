package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	grid := parseFile(file)

	fmt.Println(grid)

	for i, g := range grid.galaxies {
		fmt.Println(i+1, g)
	}

	expandDist := 1_000_000 - 1 // why minus 1?

	// expand all columns
	for x := 0; x < grid.width; x++ {
		isEmpty := true
		for _, g := range grid.galaxies {
			if g[0] == x {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			fmt.Printf("Expanding column %d\n", x)

			// loop through all galaxies and +1 to their X values if they are >= x
			for i, g := range grid.galaxies {
				if g[0] >= x {
					fmt.Printf("Expanding galaxy [%d:%d]\n", g[0], g[1])
					grid.galaxies[i][0] += expandDist
				}
			}

			grid.width += expandDist
			x += expandDist
		}
	}

	// expand all rows
	for y := 0; y < grid.height; y++ {
		isEmpty := true
		for _, g := range grid.galaxies {
			if g[1] == y {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			fmt.Printf("Expanding row %d\n", y)

			// loop through all galaxies and +1 to their Y values if they are >= y
			for i, g := range grid.galaxies {
				if g[1] >= y {
					fmt.Printf("Expanding galaxy [%d:%d]\n", g[0], g[1])
					grid.galaxies[i][1] += expandDist
				}
			}

			grid.height += expandDist
			y += expandDist
		}
	}

	fmt.Println(grid.height, grid.width)

	pairsCalculated := make(map[string]bool)

	var shortests []int

	for i1, g1 := range grid.galaxies {
		for i2, g2 := range grid.galaxies {
			key := fmt.Sprintf("%d:%d", i1, i2)

			if i1 == i2 {
				// same galaxy
				continue
			}

			if pairsCalculated[key] {
				// already calculated
				continue
			}

			keyRev := fmt.Sprintf("%d:%d", i2, i1)
			pairsCalculated[key] = true
			pairsCalculated[keyRev] = true

			s := calculateShortest(g1, g2)
			fmt.Printf("[%s]: %d\n", key, s)
			shortests = append(shortests, s)
		}
	}

	fmt.Println(len(shortests))

	total := 0
	for _, s := range shortests {
		total += s
	}

	fmt.Println(total)
}

func calculateShortest(g1, g2 Galaxy) int {
	return abs(g1[0]-g2[0]) + abs(g1[1]-g2[1])
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Galaxy []int // [x, y]

type Grid struct {
	width, height int
	galaxies      []Galaxy
}

func (g Grid) String() string {
	s := ""

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			found := false
			for _, galaxy := range g.galaxies {
				if galaxy[0] == x && galaxy[1] == y {
					s += "#"
					found = true
					break
				}
			}

			if !found {
				s += "."
			}
		}
		s += "\n"
	}

	return s
}

func parseFile(f io.Reader) Grid {
	sc := bufio.NewScanner(f)

	grid := Grid{}

	y := 0
	for sc.Scan() {
		text := sc.Text()
		line := []rune(text)
		grid.width = len(line)

		for i, c := range line {
			if c == '.' {
				continue
			}

			if c == '#' {
				grid.galaxies = append(grid.galaxies, Galaxy{i, y})
			} else {
				panic("unknown character: " + string(c))
			}
		}

		y++
	}

	grid.height = y

	return grid
}

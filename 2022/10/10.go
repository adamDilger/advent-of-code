package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	NOOP = "noop"
	ADDX = "addx"

	WIDTH  = 40
	HEIGHT = 6
)

type Comms struct {
	Cycle, X, StrengthSignal int
}

func (c *Comms) tick() {
	c.Cycle++
	c.draw()
}

func (c *Comms) addX(val string) {
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	c.tick()
	c.tick()
	c.X += i
}

func (c *Comms) draw() {
	cx := (c.Cycle - 1) % 40

	if cx == 0 {
		println()
	}

	if c.X+1 == cx {
		print("#")
	} else if c.X-1 == cx {
		print("#")
	} else if c.X == cx {
		print("#")
	} else {
		print(".")
	}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var comms Comms
	comms.X = 1

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		cmd := line[0]

		switch cmd {
		case NOOP:
			comms.tick()
		case ADDX:
			comms.addX(line[1])
		}
	}

	println()
}

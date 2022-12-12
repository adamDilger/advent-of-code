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
)

type Comms struct {
	Cycle, X, StrengthSignal int
}

func (c *Comms) tick() {
	c.Cycle++

	if c.Cycle%40 == 20 {
		println(c.Cycle, c.X, c.Cycle*c.X)
		c.StrengthSignal += c.Cycle * c.X
	}
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

	println(comms.StrengthSignal)
}

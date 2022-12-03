package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3

	LOSS = 0
	DRAW = 3
	WIN  = 6
)

var LABELS = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",

	"X": "Loss",
	"Y": "Draw",
	"Z": "Win",
}

var MOVE = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,

	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
}

var MOVE_NUMBER = []string{"", "A", "B", "C"}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Split(line, " ")
		m1 := lineSplit[0]
		wld := lineSplit[1]

		m2 := MOVE_NUMBER[calculateDesiredMove(m1, wld)]

		r := calculateScore(m1, m2)

		total += r.Score
	}

	fmt.Printf("++++++++++++++++++\n\n%d", total)

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
}

type result struct {
	Win   bool
	Draw  bool
	M1    int
	M2    int
	Score int
}

func calculateDesiredMove(a, winLossOrDraw string) int {
	m1 := MOVE[a]
	wld := MOVE[winLossOrDraw]

	if wld == DRAW {
		return m1
	}

	if wld == LOSS {
		if m1 == ROCK {
			return SCISSORS
		} else if m1 == PAPER {
			return ROCK
		} else {
			return PAPER
		}
	}

	if m1 == ROCK {
		return PAPER
	} else if m1 == PAPER {
		return SCISSORS
	} else {
		return ROCK
	}
}

func calculateScore(a, b string) result {
	m1 := MOVE[a]
	m2 := MOVE[b]

	r := result{
		M1:    m1,
		M2:    m2,
		Score: m2,
	}

	if m1 == m2 {
		r.Draw = true
		r.Score += DRAW
		return r
	}

	if m1 == ROCK && m2 == PAPER ||
		m1 == PAPER && m2 == SCISSORS ||
		m1 == SCISSORS && m2 == ROCK {
		r.Win = true
		r.Score += WIN
		return r
	}

	r.Score += LOSS
	return r
}

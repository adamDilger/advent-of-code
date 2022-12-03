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
	"X": "Rock",
	"B": "Paper",
	"Y": "Paper",
	"C": "Scissors",
	"Z": "Scissors",
}

var MOVE = map[string]int{
	"A": ROCK,
	"X": ROCK,
	"B": PAPER,
	"Y": PAPER,
	"C": SCISSORS,
	"Z": SCISSORS,
}

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
		m2 := lineSplit[1]

		fmt.Printf("GAME --- \n")
		fmt.Printf("m1 %v: %v\n", m1, LABELS[m1])
		fmt.Printf("m2 %v: %v\n", m2, LABELS[m2])

		r := calculateScore(m1, m2)

		if r.Win {
			fmt.Printf("WIN : ")
		} else if r.Draw {
			fmt.Printf("DRAW: ")
		} else {
			fmt.Printf("LOSS: ")
		}

		fmt.Printf("%d", r.Score)

		fmt.Println()

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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	run(file)
}

func runB(f io.Reader) int {
	sc := bufio.NewScanner(f)

	var l1, l2 []int

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		var left, right int
		fmt.Sscanf(t, "%d   %d\n", &left, &right)
		l1 = append(l1, left)
		l2 = append(l2, right)
	}

	slices.Sort(l1)
	slices.Sort(l2)
	// fmt.Println(l1, l2)

	sum := 0
	for item := range l1 {
		count := 0
		idx := 0

		fmt.Println(sum)
		return sum
	}
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

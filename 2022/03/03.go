package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		length := len(line) / 2
		l := line[:length]
		r := line[length:]

		var dupe rune

	RowLoop:
		for _, lc := range l {
			for _, rc := range r {
				if rc == lc {
					dupe = lc
					break RowLoop
				}
			}
		}

		total += charToPriority(dupe)
	}

	println("----------------------------")
	println(total)
}

func charToPriority(a rune) int {
	if 65 <= a && a <= 90 {
		return int(a) - 64 + 26
	}

	if 97 <= a && a <= 122 {
		return int(a) - 96
	}

	panic("invalid char")
}

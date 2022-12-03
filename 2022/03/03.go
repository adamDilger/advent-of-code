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
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		var dupe rune

	RowLoop:
		for _, one := range line1 {
			for _, two := range line2 {
				for _, three := range line3 {
					if one == two && two == three {
						dupe = one
						break RowLoop
					}
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

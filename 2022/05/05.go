package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var crateLines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if rune(line[1]) >= 48 && rune(line[1]) <= 57 {
			scanner.Scan()
			break
		}

		crateLines = append(crateLines, line)
	}

	var stacks [9][]rune

	for i := len(crateLines) - 1; i >= 0; i-- {
		currentStack := 0
		for j := 1; j < len(crateLines[i])-1; j += 4 {
			c := crateLines[i][j]
			if c == ' ' {
				currentStack++
				continue
			}
			stacks[currentStack] = append(stacks[currentStack], rune(c))
			currentStack++
		}
	}

	printStacks(stacks)

	for scanner.Scan() {
		line := scanner.Text()
		println(line)

		var amount, from, to int
		fmt.Fscanf(strings.NewReader(line), "move %d from %d to %d", &amount, &from, &to)

		stack := stacks[from-1]
		toStack := stacks[to-1]

		tmp := stack[len(stack)-amount:]
		reverse(tmp)

		stacks[from-1] = stack[:len(stack)-amount]
		stacks[to-1] = append(toStack, tmp...)
	}

	printStacks(stacks)

	printStackSum(stacks)
}

func printStack(s []rune) {
	for _, i := range s {
		print(string(i))
	}
	println()
}

func printStacks(stacks [9][]rune) {
	for k, s := range stacks {
		print(k+1, "- ")
		printStack(s)
		println()
	}
}

func printStackSum(stacks [9][]rune) {
	for _, s := range stacks {
		print(string(s[len(s)-1]))
	}
	println()
}

func reverse(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

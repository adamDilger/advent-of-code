package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	List    bool
	Val     int
	ListVal []Item
}

func buildItems(nodes []Node, index *int, currItemsPtr *[]*Item) *Item {
	node := nodes[*index]

	*index++

	n := len((*currItemsPtr)) - 1

	// end
	if node.end {
		e := (*currItemsPtr)[n]
		(*currItemsPtr) = (*currItemsPtr)[:n]

		return e
	}

	// value
	if node.val != 0 {
		(*currItemsPtr)[n].ListVal = append((*currItemsPtr)[n].ListVal, Item{Val: node.val})
	} else {
		// start
		(*currItemsPtr) = append((*currItemsPtr), &Item{List: true})
	}

	return buildItems(nodes, index, currItemsPtr)
}

/* either [, ], or a number */
type Node struct {
	val   int
	start bool
	end   bool
}

func parseLine(line string) []Node {
	var nodes []Node

	for x := 0; x < len(line); x++ {
		c := line[x]

		if c == ',' {
			continue
		} else if c == '[' {
			nodes = append(nodes, Node{start: true})
		} else if c == ']' {
			nodes = append(nodes, Node{end: true})
		} else {
			number := string(c)

			for line[x+1] != ']' && line[x+1] != ',' {
				number += string(line[x+1])
				x++
			}

			a, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			nodes = append(nodes, Node{val: a})
		}
	}

	return nodes
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()

		fmt.Println(line1)
		fmt.Println(line2)

		ind := 0
		fmt.Println(buildItems(parseLine(line1), &ind, &[]*Item{}))
		ind = 0
		fmt.Println(buildItems(parseLine(line2), &ind, &[]*Item{}))

		scanner.Scan()
	}
}

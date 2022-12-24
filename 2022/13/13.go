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
	ListVal []*Item
}

func (i *Item) printAll(newline bool) {
	for _, c := range i.ListVal {
		if c.List {
			fmt.Print("[")
			c.printAll(false)
			fmt.Print("]")
		} else {
			fmt.Print(c.Val, ",")
		}
	}

	if newline {
		fmt.Println()
	}
}

func buildItems(nodes []Node, currItemsPtr []*Item) {
	// fmt.Printf("%v\n", *currItemsPtr)
	if len(nodes) == 0 {
		return
	}

	node := nodes[0]
	nodes = nodes[1:]

	lastIndex := len(currItemsPtr) - 1

	if node.start {
		// add an item to the list
		i := Item{List: true}

		// append to current item
		currItemsPtr[lastIndex].ListVal = append(currItemsPtr[lastIndex].ListVal, &i)
		currItemsPtr = append(currItemsPtr, &i)
	} else if node.end {
		currItemsPtr = currItemsPtr[:lastIndex]
	} else {
		currItemsPtr[lastIndex].ListVal = append(currItemsPtr[lastIndex].ListVal, &Item{Val: node.val})
		// fmt.Println(lastItem)
	}

	buildItems(nodes, currItemsPtr)
}

/* either [, ], or a number */
type Node struct {
	val   int
	start bool
	end   bool
}

func parseLine(line string) []Node {
	var nodes []Node

	for x := 1; x < len(line)-1; x++ {
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

		ind1 := &Item{List: true}
		buildItems(parseLine(line1), []*Item{ind1})
		ind1.printAll(true)

		ind2 := &Item{List: true}
		buildItems(parseLine(line2), []*Item{ind2})
		ind2.printAll(true)

		scanner.Scan()
		fmt.Println("------------------------------")
	}
}

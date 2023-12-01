package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Item struct {
	IsList  bool
	Val     int
	ListVal []*Item
}

func (i *Item) String() string {
	if !i.IsList {
		return fmt.Sprintf("%d", i.Val)
	}

	out := "["
	for idx, v := range i.ListVal {
		if idx > 0 {
			out += ","
		}
		out += fmt.Sprintf("%v", v.String())
	}

	out += "]"

	return out
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
		i := Item{IsList: true}

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
	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	scanner := bufio.NewScanner(
		strings.NewReader("[[1],[2,3,4]]\n[[1],4]\n"),
	)

	count := 1

	for scanner.Scan() {
		fmt.Printf("== Pair %d ==\n", count)

		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()

		ind1 := &Item{IsList: true}
		buildItems(parseLine(line1), []*Item{ind1})

		ind2 := &Item{IsList: true}
		buildItems(parseLine(line2), []*Item{ind2})

		fmt.Printf("- Compare %s vs %s\n", line1, line2)

		fmt.Printf("%v\n", ind1)
		fmt.Printf("%v\n", ind2)

		indexes := compare(ind1, ind2, 0, []int{})
		fmt.Println(indexes)

		scanner.Scan()
		count++
		println()
		break
	}
}

func compare(i1, i2 *Item, curIndex int, idx []int) []int {
	if !i1.IsList && !i2.IsList {
		// compare the values
		if i1.Val > i2.Val {
			idx = append(idx, curIndex)
		}

		return idx
	}

	return idx
}

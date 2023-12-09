package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	dir, allNodes := parseNodes(file)
	fmt.Println(dir)

	nodes := make(map[string]Node)
	for _, n := range allNodes {
		nodes[n.Name] = n
		fmt.Println(n)
	}

	fmt.Println(doThing(0, 0, dir, nodes, nodes["AAA"]))
}

func doThing(moveCount int, dirIndex int, dirs []Dir, nodes map[string]Node, node Node) int {
	if node.Name == "ZZZ" {
		fmt.Println("Found", moveCount)
		return moveCount
	}

	fmt.Println(dirs)
	dir := dirs[dirIndex%len(dirs)]

	moveCount++
	dirIndex++

	if dir == L {
		fmt.Println("Going left", nodes[node.Left])
		return doThing(moveCount, dirIndex, dirs, nodes, nodes[node.Left])
	} else {
		fmt.Println("Going right", nodes[node.Right])
		return doThing(moveCount, dirIndex, dirs, nodes, nodes[node.Right])
	}
}

type Node struct {
	Name        string
	Left, Right string
}

type Dir string

const (
	L Dir = "L"
	R Dir = "R"
)

func parseNodes(f io.Reader) ([]Dir, []Node) {
	var nodes []Node

	sc := bufio.NewScanner(f)
	sc.Scan()

	var dir []Dir
	dirRune := []rune(sc.Text())
	for _, d := range dirRune {
		dir = append(dir, Dir(d))
	}

	sc.Scan()

	// AAA = (BBB, CCC)
	r := regexp.MustCompile("([A-Z]+) = \\(([A-Z]+), ([A-Z]+)\\)")
	for sc.Scan() {
		line := sc.Text()

		res := r.FindStringSubmatch(line)
		nodes = append(nodes, Node{
			Name:  res[1],
			Left:  res[2],
			Right: res[3],
		})
	}

	return dir, nodes
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
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

	var ghosts []Node
	for _, n := range allNodes {
		if n.Start {
			ghosts = append(ghosts, n)
		}
	}

	var dirs []Dir
	for _, d := range dir {
		dirs = append(dirs, d)
	}

	moveCount := 0

	for {
		println("Move", moveCount)
		isFinished := true
		for _, n := range ghosts {
			if !n.End {
				isFinished = false
				break
			}
		}

		if isFinished {
			fmt.Println("Finished", moveCount)
			break
		}

		dir := dirs[moveCount%len(dirs)]
		moveCount++

		// move all nodes
		for i := range ghosts {
			if dir == L {
				// fmt.Println("Going left", nodes[node.Left])
				ghosts[i] = nodes[ghosts[i].Left]
			} else {
				// fmt.Println("Going right", nodes[node.Right])
				ghosts[i] = nodes[ghosts[i].Right]
			}
		}
	}
}

type Node struct {
	Name        string
	Left, Right string
	End, Start  bool
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
	r := regexp.MustCompile("([A-Z1-9]+) = \\(([A-Z1-9]+), ([A-Z1-9]+)\\)")
	for sc.Scan() {
		line := sc.Text()

		res := r.FindStringSubmatch(line)
		nodes = append(nodes, Node{
			Name:  res[1],
			Left:  res[2],
			Right: res[3],
			End:   strings.HasSuffix(res[1], "Z"),
			Start: strings.HasSuffix(res[1], "A"),
		})
	}

	return dir, nodes
}

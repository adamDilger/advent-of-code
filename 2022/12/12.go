package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

const debug = true

const (
	START = 'S'
	END   = 'E'
)

func log(args ...any) {
	if debug {
		fmt.Println(args...)
	}
}

type Point struct {
	x, y int
	key  string
}

func (p *Point) getKey() string {
	if p.key == "" {
		p.key = fmt.Sprintf("%d:%d", p.x, p.y)
	}
	return p.key
}

func (p *Point) left() Point {
	return Point{x: p.x - 1, y: p.y}
}

func (p *Point) right() Point {
	return Point{x: p.x + 1, y: p.y}
}

func (p *Point) up() Point {
	return Point{x: p.x, y: p.y - 1}
}

func (p *Point) down() Point {
	return Point{x: p.x, y: p.y + 1}
}

type Grid struct {
	X, Y  int
	Grid  [][]rune
	Nodes map[string]*Node
}

func (g *Grid) printGrid() {
	for y := range g.Grid {
		for x := range g.Grid[y] {
			print(string(g.Grid[y][x]))
		}

		println()
	}
}

func (g *Grid) endPoint() *Node {
	for _, n := range g.Nodes {
		if n.End {
			return n
		}
	}

	panic("No end found")
}

func (g *Grid) startPoint() *Node {
	for _, n := range g.Nodes {
		if n.Start {
			return n
		}
	}

	panic("No start found")
}

func (g *Grid) newNode(point Point) *Node {
	x, y := point.x, point.y
	char := g.Grid[y][x]

	n := &Node{
		Point: point,
		Value: char,

		ShortestPathVal: math.MaxInt,
		ShortestPathVia: nil,

		HeapIndex: -1,
	}
	n.Edges = make(map[string]int)

	if char == START {
		n.Start = true
		n.ShortestPathVal = 0
		n.Value = 'a'
	} else if char == END {
		n.End = true
		n.Value = 'z'
	}

	checkEdge := func(p Point) {
		nextChar := g.Grid[p.y][p.x]
		if nextChar == START {
			nextChar = 'a'
		} else if nextChar == END {
			nextChar = 'z'
		}

		weight := -1

		// current 83
		// next    82
		//  83 - 82 = 1 | 1 + 1 = 1

		// current 83
		// next    83
		//  83 - 83 = 0 | 0 + 1 = 1

		// current 83
		// next    84
		// 1

		// 83 - 82 = 1 + 2 = 3

		if nextChar == n.Value {
			weight = 1
			// } else if nextChar < n.Value {
			// 	weight = int(n.Value-nextChar) + 1
		} else if nextChar == n.Value+1 {
			weight = 0
		}

		if weight != -1 {
			n.Edges[p.getKey()] = weight
		}
	}

	if y != 0 {
		checkEdge(point.up())
	}

	if y < len(g.Grid)-1 {
		checkEdge(point.down())
	}

	if x != 0 {
		checkEdge(point.left())
	}

	if x < len(g.Grid[y])-1 {
		checkEdge(point.right())
	}

	return n
}

type Node struct {
	Point
	Value      rune
	Start, End bool
	Edges      map[string]int

	ShortestPathVal int
	ShortestPathVia *Node

	HeapIndex int
}

type Graph struct {
	Nodes map[string]Node

	Seen map[string]struct{}
}

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var grid Grid
	grid.Nodes = make(map[string]*Node)

	for scanner.Scan() {
		line := scanner.Text()
		log(line)

		grid.Grid = append(grid.Grid, []rune(line))
	}

	for y := range grid.Grid {
		for x := range grid.Grid[y] {
			n := grid.newNode(Point{x: x, y: y})
			grid.Nodes[n.getKey()] = n
		}
	}

	visited := make(map[string]struct{})

	startNode := grid.startPoint()
	startNode.HeapIndex = 0
	heapNodes := &NodeHeap{}
	heap.Push(heapNodes, startNode)

	c := 0

	for heapNodes.Len() != 0 {
		n := heap.Pop(heapNodes).(*Node)

		visited[n.getKey()] = struct{}{}

		if n.End {
			break
		}

		c++
		println(c)

		for eKey, eWeight := range n.Edges {
			if _, ok := visited[eKey]; ok {
				continue
			}

			eNode := grid.Nodes[eKey]

			if n.ShortestPathVal+eWeight < eNode.ShortestPathVal {
				eNode.ShortestPathVal = n.ShortestPathVal + eWeight
				eNode.ShortestPathVia = n
			}

			heap.Push(heapNodes, eNode)
			heap.Fix(heapNodes, eNode.HeapIndex)
		}
	}

	fmt.Println()

	fmt.Println(grid.startPoint())
	fmt.Println(grid.endPoint())

	count := 0
	p := grid.endPoint()
	for !p.Start {
		fmt.Printf("%c at %s\n", p.Value, p.key)
		p = p.ShortestPathVia
		count++
	}

	println(count)
}

type NodeHeap []*Node

func (h NodeHeap) Len() int { return len(h) }
func (h NodeHeap) Less(i, j int) bool {
	return h[i].ShortestPathVal < h[j].ShortestPathVal
}
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].HeapIndex = i
	h[j].HeapIndex = j
}

func (h *NodeHeap) Push(x any) {
	// *h = append(*h, x.(*Node))
	n := len(*h)
	item := x.(*Node)
	item.HeapIndex = n
	*h = append(*h, item)
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil   // avoid memory leak
	x.HeapIndex = -1 // for safety
	*h = old[0 : n-1]
	return x
}

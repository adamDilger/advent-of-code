package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	CD  = "cd"
	LS  = "ls"
	DIR = "dir"

	PREV_DIR = ".."
)

type File struct {
	Name     string
	IsDir    bool
	Size     int
	Children []*File

	Depth int

	Parent *File
}

func (f *File) p() {
	for i := 0; i < f.Depth; i++ {
		print("- ")
	}

	fmt.Printf("/%s : %d\n", f.Name, f.Size)

	for _, c := range f.Children {
		if c.IsDir {
			c.p()
		} else {
			for i := 0; i < f.Depth; i++ {
				// print("- ")
			}
			// fmt.Printf("  f %s (%d)\n", c.Name, c.Size)
		}
	}
}

func (f *File) calculateDirectoryTotals() int {
	for _, c := range f.Children {
		if c.IsDir {
			f.Size += c.calculateDirectoryTotals()
		} else {
			f.Size += c.Size
		}
	}

	return f.Size
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var root *File
	var current *File

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		args := strings.Split(line, " ")

		if args[0] == "$" {
			fmt.Printf("cmd: %s\n", line)

			if args[1] == CD {
				loc := args[2]

				if loc == PREV_DIR {
					fmt.Printf("cd to PREV_DIR\n")
					current = current.Parent
				} else {
					var n *File

					if current == nil {
						root = &File{IsDir: true, Name: loc, Depth: 0}
						n = root
					} else {
						for _, c := range current.Children {
							if c.Name == loc {
								n = c
								break
							}
						}
					}

					current = n
				}
			} else if args[1] == LS {
				// fmt.Printf("ls\n")
			}
		} else {
			newFile := &File{Name: args[1], Parent: current, Depth: current.Depth + 1}
			current.Children = append(current.Children, newFile)

			// fmt.Printf("[%s] %s\n", currentDir.Name, line)

			if args[0] == DIR {
				fmt.Printf("[%s]\n", args[1])
				newFile.IsDir = true
			} else {
				size, _ := strconv.Atoi(args[0])
				fmt.Printf("{%d} %s\n", size, args[1])
				newFile.Size = size
			}
		}
	}

	println("-----------------------")
	root.calculateDirectoryTotals()
	root.p()

	var maxes []*File
	bfs(root, &maxes)

	count := 0
	for _, m := range maxes {
		count += m.Size
	}
	println(count)
}

func bfs(dir *File, maxes *[]*File) {
	if !dir.IsDir {
		return
	}

	if dir.Size <= 100000 {
		*maxes = append(*maxes, dir)
	}

	for _, child := range dir.Children {
		bfs(child, maxes)
	}
}

package main

import (
	"fmt"
	"os"
	"weektwo/a"
	"weektwo/b"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 || os.Args[1] == "a" {
		fmt.Println("A:", a.RunA(file))
	} else {
		fmt.Println("A:", b.RunB(file))
	}
}

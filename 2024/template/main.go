package main

import (
	"fmt"
	"os"
	"time"
	"weekthree/a"
	"weekthree/b"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := time.Now()

	var res int
	running := "A"

	if len(os.Args) < 2 || os.Args[1] == "a" {
		res = a.RunA(file)
	} else {
		running = "B"
		res = b.RunB(file)
	}

	fmt.Println("Time taken:", time.Since(s))
	fmt.Println("Running", running)
	fmt.Println(res)
}

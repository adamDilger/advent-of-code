package main

import (
	weekone "2024/01"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 || os.Args[1] == "a" {
		fmt.Println("A:", weekone.RunA(file))
	} else {
		fmt.Println("A:", weekone.RunB(file))
	}
}

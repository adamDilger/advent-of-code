package main

import (
	"fmt"
	"os"
	"week<week_number>"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 || os.Args[1] == "a" {
		fmt.Println("A:", <week_number>.RunA(file))
	} else {
		fmt.Println("A:", <week_number>.RunB(file))
	}
}

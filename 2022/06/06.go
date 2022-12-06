package main

import (
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	res := make(map[rune]bool)

MainLoop:
	for i := 0; i < len(input); i++ {
		println(string(input[i]))

		for j := 0; j < 14; j++ {
			print(string(input[i+j]))

			if res[rune(input[i+j])] {
				res = make(map[rune]bool)
				continue MainLoop
			}

			res[rune(input[i+j])] = true
		}

		println("MATCH", i+14)
		break
		// if i == 2 {
		// 	break
		// }
	}
}

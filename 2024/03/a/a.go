package a

import (
	"bufio"
	"fmt"
	"io"
)

func RunA(f io.Reader) int {
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		fmt.Println(t)
	}

	return 0
}

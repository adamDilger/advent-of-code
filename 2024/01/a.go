package weekone

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

func RunA(f io.Reader) int {
	sc := bufio.NewScanner(f)

	var l1, l2 []int

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		var left, right int
		fmt.Sscanf(t, "%d   %d\n", &left, &right)
		l1 = append(l1, left)
		l2 = append(l2, right)
	}

	slices.Sort(l1)
	slices.Sort(l2)
	// fmt.Println(l1, l2)

	sum := 0
	for i := 0; i < len(l1); i++ {
		l, r := l1[i], l2[i]

		// fmt.Println(l, r, abs(l-r))
		sum += abs(l - r)
	}

	return sum
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

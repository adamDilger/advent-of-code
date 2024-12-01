package weekone

import (
	"bufio"
	"fmt"
	"io"
)

func RunB(f io.Reader) int {
	sc := bufio.NewScanner(f)

	var l1 []int
	l2 := make(map[int]int)

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		var left, right int
		fmt.Sscanf(t, "%d   %d\n", &left, &right)
		l1 = append(l1, left)

		l2[right] = (l2[right] + 1)
	}

	sum := 0
	for _, i := range l1 {
		if l2[i] == 0 {
			continue
		}

		sum += l2[i] * i
	}

	return sum
}

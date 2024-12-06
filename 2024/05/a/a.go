package a

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Page struct {
	pre, post map[string]struct{}
}

func RunA(f io.Reader) int {
	sc := bufio.NewScanner(f)

	linesPageOrdering := []string{}
	linesUpdates := []string{}

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		if t == "" {
			break
		}

		linesPageOrdering = append(linesPageOrdering, t)
	}

	for sc.Scan() {
		t := sc.Text()

		if sc.Err() != nil {
			panic(sc.Err())
		}

		linesUpdates = append(linesUpdates, t)
	}

	// fmt.Println(linesPageOrdering)
	// fmt.Println(linesUpdates)

	pages := make(map[string]Page)

	for _, l := range linesPageOrdering {
		splits := strings.Split(l, "|")
		n1, n2 := splits[0], splits[1]

		// add first number
		if _, ok := pages[n1]; !ok {
			pages[n1] = Page{
				pre:  make(map[string]struct{}),
				post: make(map[string]struct{}),
			}
		}

		pages[n1].post[n2] = struct{}{}

		// add first number
		if _, ok := pages[n2]; !ok {
			pages[n2] = Page{
				pre:  make(map[string]struct{}),
				post: make(map[string]struct{}),
			}
		}

		pages[n2].pre[n1] = struct{}{}
	}

	// for k, v := range pages {
	// 	for pre := range v.pre {
	// 		fmt.Printf("%s ", pre)
	// 	}
	// 	fmt.Printf("|| %s || ", k)
	// 	for post := range v.post {
	// 		fmt.Printf("%s ", post)
	// 	}
	// 	fmt.Println()
	// }

	sum := 0
	for _, l := range linesUpdates {
		numbers := strings.Split(l, ",")
		valid := isValid(pages, numbers)
		if valid {
			middle := numbers[len(numbers)/2]
			middleInt, err := strconv.Atoi(middle)
			if err != nil {
				panic(err)
			}

			sum += middleInt
		}
	}

	return sum
}

func isValid(pages map[string]Page, numbers []string) bool {
	for i, n := range numbers {
		// fmt.Println(n)

		// look left
		for x := i - 1; x >= 0; x-- {
			// fmt.Println("<", i, n, numbers[x])
			if _, ok := pages[n].pre[numbers[x]]; !ok {
				return false
			}
		}

		// look right
		for x := i + 1; x < len(numbers); x++ {
			// fmt.Println(">", i, n, numbers[x])
			if _, ok := pages[n].post[numbers[x]]; !ok {
				return false
			}
		}
	}

	return true
}

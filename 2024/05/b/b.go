package b

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Page struct {
	value     string
	pre, post map[string]struct{}
}

type PageSlice []Page

func (p PageSlice) Len() int {
	return len(p)
}

func (p PageSlice) Less(i, j int) bool {
	p1, p2 := p[i], p[j]

	if _, ok := p2.pre[p1.value]; !ok {
		// p1 is supposed to be before p2
		return true
	}

	return false
}

func (p PageSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func RunB(f io.Reader) int {
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

	pages := make(map[string]Page)

	for _, l := range linesPageOrdering {
		splits := strings.Split(l, "|")
		n1, n2 := splits[0], splits[1]

		// add first number
		if _, ok := pages[n1]; !ok {
			pages[n1] = Page{
				value: n1,
				pre:   make(map[string]struct{}),
				post:  make(map[string]struct{}),
			}
		}

		pages[n1].post[n2] = struct{}{}

		// add first number
		if _, ok := pages[n2]; !ok {
			pages[n2] = Page{
				value: n2,
				pre:   make(map[string]struct{}),
				post:  make(map[string]struct{}),
			}
		}

		pages[n2].pre[n1] = struct{}{}
	}

	sum := 0
	for _, l := range linesUpdates {
		numbers := strings.Split(l, ",")
		pageSlice := PageSlice{}
		for _, n := range numbers {
			pageSlice = append(pageSlice, pages[n])
		}

		valid := isValid(pages, pageSlice)
		if valid {
			continue
		}

		// sort the list
		sort.Sort(pageSlice)

		middle := pageSlice[len(pageSlice)/2]
		middleInt, err := strconv.Atoi(middle.value)
		if err != nil {
			panic(err)
		}

		sum += middleInt
	}

	return sum
}

func isValid(pages map[string]Page, numbers PageSlice) bool {
	for i, n := range numbers {
		// fmt.Println(n)

		// look left
		for x := i - 1; x >= 0; x-- {
			// fmt.Println("<", i, n, numbers[x])
			if _, ok := pages[n.value].pre[numbers[x].value]; !ok {
				return false
			}
		}

		// look right
		for x := i + 1; x < len(numbers); x++ {
			// fmt.Println(">", i, n, numbers[x])
			if _, ok := pages[n.value].post[numbers[x].value]; !ok {
				return false
			}
		}
	}

	return true
}

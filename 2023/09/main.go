package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	histories := parseFile(file)

	for i := range histories {
		calculateDiffs(&histories[i])
		predictHistories(&histories[i])
		fmt.Println(histories[i].String())
		fmt.Printf("\n%s\n\n", strings.Repeat("-", 80))
	}

	total := 0
	for i := range histories {
		total += histories[i].prediction
	}

	println(total)
}

func calculateDiffs(h *History) {
	end := true
	for i := range h.values {
		if h.values[i] != 0 {
			end = false
			break
		}
	}

	if end {
		return
	}

	h.child = &History{}

	for i, v := range h.values {
		if i+1 == len(h.values) {
			break
		}

		h.child.values = append(h.child.values, h.values[i+1]-v)
	}

	calculateDiffs(h.child)
}

func predictHistories(h *History) {
	end := h.child == nil

	if end {
		h.prediction = 0
		h.values = append(h.values, h.prediction)
		return
	}

	predictHistories(h.child)

	h.prediction = h.child.prediction + h.values[len(h.values)-1]
	h.values = append(h.values, h.prediction)
}

type History struct {
	values []int

	child *History

	prediction int
}

func (h History) String() string {
	var sb strings.Builder

	h.InnerString(0, &sb)

	return sb.String()
}

func (h History) InnerString(indentLevel int, sb *strings.Builder) string {
	indent := strings.Repeat(" ", indentLevel*4)
	sb.WriteString(indent)

	for _, v := range h.values {
		sb.WriteString(fmt.Sprintf("%8d", v))
	}

	if h.child != nil {
		sb.WriteString("\n")
		h.child.InnerString(indentLevel+1, sb)
	}

	return sb.String()
}

func parseFile(f io.Reader) []History {
	var histories []History

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		h := History{}

		line := sc.Text()
		stringVals := strings.Fields(line)

		for _, stringVal := range stringVals {
			intVal, err := strconv.Atoi(stringVal)
			if err != nil {
				panic(err)
			}

			h.values = append(h.values, intVal)
		}

		histories = append(histories, h)
	}

	return histories
}

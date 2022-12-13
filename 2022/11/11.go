package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	MULTIPLY = "*"
	ADD      = "+"
	SUBTRACT = "+"
)

type Test struct {
	ByValue                 int
	TrueTarget, FalseTarget int
}

type Monkey struct {
	TotalItemsInspected int
	Number              int
	Items               []int
	Operation           string
	OperationValue      int
	OperationValueOld   bool
	Test                Test
}

const debug = false

func log(args ...any) {
	if debug {
		fmt.Println(args...)
	}
}

func NewMonkey(scanner *bufio.Scanner) Monkey {
	m := Monkey{}

	var err error

	var parseLine = func(format string, params ...any) {
		if err != nil {
			return
		}

		t := strings.Trim(scanner.Text(), " ")
		_, err = fmt.Fscanf(strings.NewReader(t), format, params...)

		if err != nil {
			panic(err)
		}

		scanner.Scan()
	}

	parseLine("Monkey %d:", &m.Number)

	itemsString := strings.TrimLeft(scanner.Text(), "Starting items: ")
	for _, v := range strings.Split(itemsString, ", ") {
		intV, _ := strconv.Atoi(v)
		m.Items = append(m.Items, intV)
	}
	scanner.Scan()

	var op string
	parseLine("Operation: new = old %s %s", &m.Operation, &op)
	if op == "old" {
		m.OperationValueOld = true
	} else {
		m.OperationValue, _ = strconv.Atoi(op)
	}

	parseLine("Test: divisible by %d", &m.Test.ByValue)
	parseLine("If true: throw to monkey %d", &m.Test.TrueTarget)
	parseLine("If false: throw to monkey %d", &m.Test.FalseTarget)

	if err != nil {
		panic(err)
	}

	return m
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var monkeys []Monkey

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		m := NewMonkey(scanner)
		monkeys = append(monkeys, m)
	}

	for i := 0; i < 20; i++ {

		for mi := range monkeys {
			m := &monkeys[mi]

			loopingItems := m.Items
			m.Items = []int{}

			for _, ii := range loopingItems {
				newValue := ii

				m.TotalItemsInspected++

				opVal := m.OperationValue
				if m.OperationValueOld {
					opVal = newValue
				}

				log(newValue)

				switch m.Operation {
				case ADD:
					newValue = newValue + opVal
				case MULTIPLY:
					newValue = newValue * opVal
				}

				log(newValue)
				newValue = newValue / 3
				log(newValue)

				if newValue%m.Test.ByValue == 0 {
					monkeys[m.Test.TrueTarget].Items = append(monkeys[m.Test.TrueTarget].Items, newValue)
					log("TRUE", m.Test.TrueTarget)
				} else {
					log("FALSE", m.Test.FalseTarget)
					monkeys[m.Test.FalseTarget].Items = append(monkeys[m.Test.FalseTarget].Items, newValue)
				}

				log()
			}

			log("---------")
		}

		printMonkeys(i, monkeys)
	}

	printMonkeyTotals(monkeys)

	var highest []int
	for _, m := range monkeys {
		highest = append(highest, m.TotalItemsInspected)
	}

	sort.Ints(highest)
	fmt.Println(highest)
	highest = highest[len(highest)-2:]
	fmt.Println(highest)

	fmt.Println(highest[0] * highest[1])
}

func printMonkeys(round int, monkeys []Monkey) {
	fmt.Printf("Round: %d\n---------------------\n", round+1)

	for _, m := range monkeys {
		fmt.Printf("Monkey %d: ", m.Number)

		for _, i := range m.Items {
			fmt.Printf("%d, ", i)
		}

		fmt.Println()
	}

	fmt.Printf("---------------------\n")
}

func printMonkeyTotals(monkeys []Monkey) {
	fmt.Printf("----\nTOTALS\n---------------------\n")

	for _, m := range monkeys {
		fmt.Printf("Monkey %d: %d\n", m.Number, m.TotalItemsInspected)
	}

	fmt.Printf("---------------------\n")
}

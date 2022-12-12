package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Thrower struct {
	Divisor       int
	TargetIfTrue  int
	TargetIfFalse int
}

type Monkey struct {
	Items          []int
	WorryOperation string
	Thrower        Thrower
	InspectCount   int
}

func (m Monkey) hasItems() bool {
	return len(m.Items) != 0
}

func (m *Monkey) popItem() int {
	m.InspectCount++
	item := m.Items[0]
	m.Items = m.Items[1:]
	return item
}

func (m *Monkey) receive(i int) {
	m.Items = append(m.Items, i)
}

func (m Monkey) worry(worryLevel int) int {
	componentParts := strings.Split(m.WorryOperation, " ")

	left, err := strconv.Atoi(componentParts[0])
	if err != nil {
		left = worryLevel
	}

	right, err := strconv.Atoi(componentParts[2])
	if err != nil {
		right = worryLevel
	}
	operator := componentParts[1]

	switch operator {
	case "+":
		return left + right
	case "*":
		return left * right
	case "-":
		return left - right
	}

	log.Fatal("Some worrying went wrong")
	return 0
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err.Error())
	}

	fullMonkeys, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	monkeyMatch := regexp.MustCompile(
		"(?m)Monkey [0-9]+:\\n *Starting items: (([0-9]+(, )?)+)\\n" +
			" *Operation: new = (.*)\\n" +
			" *Test: divisible by ([0-9]+)\\n" +
			".* true: throw to monkey ([0-9]+)\\n" +
			".*throw to monkey ([0-9]+)",
	)

	monkeyMatches := monkeyMatch.FindAllStringSubmatch(string(fullMonkeys), 1000)
	var monkeys []*Monkey
	for _, monkeyItems := range monkeyMatches {
		itemString := monkeyItems[1]
		items := strings.Split(itemString, ", ")
		var itemInts []int
		for _, i := range items {
			item, err := strconv.Atoi(i)
			if nil != err {
				log.Fatal(err)
			}
			itemInts = append(itemInts, item)
		}

		m := &Monkey{
			Items:          itemInts,
			WorryOperation: monkeyItems[4],
		}

		divisor, err := strconv.Atoi(monkeyItems[5])
		if nil != err {
			log.Fatal(err)
		}
		targetIfTrue, err := strconv.Atoi(monkeyItems[6])
		if nil != err {
			log.Fatal(err)
		}
		targetIfFalse, err := strconv.Atoi(monkeyItems[7])
		if nil != err {
			log.Fatal(err)
		}
		t := Thrower{
			Divisor:       divisor,
			TargetIfTrue:  targetIfTrue,
			TargetIfFalse: targetIfFalse,
		}

		m.Thrower = t
		monkeys = append(monkeys, m)
	}
	commonDivisor := 1
	for _, m := range monkeys {
		commonDivisor *= m.Thrower.Divisor
	}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			for m.hasItems() {
				worryLevel := m.popItem()
				worryLevel = m.worry(worryLevel)
				worryLevel = worryLevel % commonDivisor
				if worryLevel%m.Thrower.Divisor == 0 {
					monkeys[m.Thrower.TargetIfTrue].receive(worryLevel)
				} else {
					monkeys[m.Thrower.TargetIfFalse].receive(worryLevel)
				}
			}
		}
	}

	sort.Slice(monkeys, func(a, b int) bool {
		return monkeys[a].InspectCount > monkeys[b].InspectCount
	})

	res := 1
	for _, m := range monkeys[:2] {
		res *= m.InspectCount
	}
	fmt.Println(res)
}

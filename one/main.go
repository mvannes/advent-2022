package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Foods []int
}

func (e Elf) totalCalories() int {
	var total int
	for _, f := range e.Foods {
		total += f
	}
	return total
}

func main() {

	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}

	var elves []Elf
	var bestElf Elf
	var currentElf *Elf
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		if nil == currentElf {
			currentElf = &Elf{}
		}

		line := scan.Text()

		if "" == line {
			elves = append(elves, *currentElf)
			if bestElf.totalCalories() < currentElf.totalCalories() {
				bestElf = *currentElf
			}
			currentElf = nil
			continue
		}
		calories, err := strconv.ParseInt(line, 10, 0)
		if nil != err {
			log.Fatal(err.Error())
		}
		currentElf.Foods = append(currentElf.Foods, int(calories))
	}

	sort.Slice(elves, func(a, b int) bool {
		return elves[a].totalCalories() > elves[b].totalCalories()
	})
	var totalCarried int
	for _, e := range elves[0:3] {
		fmt.Println(e.totalCalories())
		totalCarried += e.totalCalories()
	}

	fmt.Println(totalCarried)
	fmt.Println(len(elves))
	fmt.Println(bestElf)
	fmt.Println(bestElf.totalCalories())
}

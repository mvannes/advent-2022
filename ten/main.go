package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatalln(err.Error())
	}
	s := bufio.NewScanner(f)
	cycleCount := 0
	register := 1
	addRegex := regexp.MustCompile("addx ([-0-9]+)")

	cycleToPrint := 20
	var sumCycles int
	for s.Scan() {
		cycleCount++

		if cycleToPrint == cycleCount {
			sumCycles += cycleCount * register
			cycleToPrint += 40
		}

		line := s.Text()
		if line == "noop" {
			continue
		}

		cycleCount++
		if cycleToPrint == cycleCount {
			sumCycles += cycleCount * register
			cycleToPrint += 40
		}
		addMatch := addRegex.FindStringSubmatch(line)
		registerAdd, err := strconv.Atoi(addMatch[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		register += registerAdd
	}

	fmt.Println(sumCycles)
}

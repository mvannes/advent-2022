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

	var screen [6][40]string

	cycleCount := 0
	register := 1
	addRegex := regexp.MustCompile("addx ([-0-9]+)")
	for s.Scan() {
		cycleCount++
		draw(cycleCount, &screen, register)

		line := s.Text()
		if line == "noop" {
			continue
		}
		cycleCount++

		addMatch := addRegex.FindStringSubmatch(line)
		registerAdd, err := strconv.Atoi(addMatch[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		register += registerAdd
		draw(cycleCount, &screen, register)

	}

	for _, r := range screen {
		fmt.Println(r)
	}
}

func draw(cycleCount int, screen *[6][40]string, register int) {
	rowNumber := ((cycleCount - 1) / 40)
	rowPosition := cycleCount % 40

	var absReg int
	if rowPosition >= register {
		absReg = rowPosition - register
	} else {
		absReg = register - rowPosition
	}
	if absReg <= 1 {
		screen[rowNumber][rowPosition] = "#"
	} else {
		screen[rowNumber][rowPosition] = "."
	}
}

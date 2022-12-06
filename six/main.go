package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		resOne := findSignal(line, 4)
		fmt.Println(resOne)
		resTwo := findSignal(line, 14)
		fmt.Println(resTwo)
	}
}

func findSignal(s string, amountOfChars int) int {
	for i := amountOfChars; i < len(s); i++ {
		sub := s[i-amountOfChars : i]
		set := map[rune]bool{}
		for _, s := range sub {
			set[s] = true
		}
		if len(set) == amountOfChars {
			return i
		}
	}
	return 0
}

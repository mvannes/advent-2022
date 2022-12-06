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
		res := findSignal(line)
		fmt.Println(res)
	}
}

func findSignal(s string) int {
	for i := 4; i < len(s); i++ {
		sub := s[i-4 : i]
		set := map[rune]bool{}
		for _, s := range sub {
			set[s] = true
		}
		fmt.Println(set)
		if len(set) == 4 {
			return i
		}
	}
	return 0
}

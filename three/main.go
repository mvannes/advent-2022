package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var validCharacters = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}
	var prioSum int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		totalItems := len(line)
		halfWay := totalItems / 2
		compartmentOne := line[0:halfWay]
		compartmentTwo := line[halfWay:]

		for _, c := range compartmentOne {
			if strings.ContainsRune(compartmentTwo, c) {
				prioSum += runeToPrio(c)
				break
			}
		}
	}
	fmt.Println(prioSum)
}

func runeToPrio(input rune) int {
	for prio, r := range validCharacters {
		if input == r {
			// 0 indexed.
			return (prio + 1)
		}
	}
	return -1
}

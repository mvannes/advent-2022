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
	var currentGroup []string
	var counter int
	for s.Scan() {
		line := s.Text()
		counter++
		currentGroup = append(currentGroup, line)
		if counter < 3 {
			continue
		}

		compartmentOne := currentGroup[0]
		compartmentTwo := currentGroup[1]
		compartmentThree := currentGroup[2]

		for _, c := range compartmentOne {
			if strings.ContainsRune(compartmentTwo, c) && strings.ContainsRune(compartmentThree, c) {
				prioSum += runeToPrio(c)
				break
			}
		}
		counter = 0
		currentGroup = []string{}
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

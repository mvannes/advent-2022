package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Stack struct {
	crates []string
}

func (s Stack) peek() string {
	return s.crates[len(s.crates)-1]
}

func (s *Stack) pop() string {
	poppedElement := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return poppedElement
}

func (s *Stack) put(crate string) {
	s.crates = append(s.crates, crate)
}

func (s *Stack) popMultiple(takeAmount int) []string {
	poppedElements := s.crates[len(s.crates)-1-takeAmount : len(s.crates)]
	s.crates = s.crates[:len(s.crates)-1-takeAmount]
	return poppedElements
}

func main() {
	//[N]             [R]             [C]
	//[T] [J]         [S] [J]         [N]
	//[B] [Z]     [H] [M] [Z]         [D]
	//[S] [P]     [G] [L] [H] [Z]     [T]
	//[Q] [D]     [F] [D] [V] [L] [S] [M]
	//[H] [F] [V] [J] [C] [W] [P] [W] [L]
	//[G] [S] [H] [Z] [Z] [T] [F] [V] [H]
	//[R] [H] [Z] [M] [T] [M] [T] [Q] [W]
	// 1   2   3   4   5   6   7   8   9
	stacks := []Stack{
		{crates: []string{"R", "G", "H", "Q", "S", "B", "T", "N"}},
		{crates: []string{"H", "S", "F", "D", "P", "Z", "J"}},
		{crates: []string{"Z", "H", "V"}},
		{crates: []string{"M", "Z", "J", "F", "G", "H"}},
		{crates: []string{"T", "Z", "C", "D", "L", "M", "S", "R"}},
		{crates: []string{"M", "T", "W", "V", "H", "Z", "J"}},
		{crates: []string{"T", "F", "P", "L", "Z"}},
		{crates: []string{"Q", "V", "W", "S"}},
		{crates: []string{"W", "H", "L", "M", "T", "D", "N", "C"}},
	}

	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		var amountToMove, startStackIndex, endStackIndex int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &amountToMove, &startStackIndex, &endStackIndex)
		if nil != err {
			log.Fatal(err)
		}
		// Down one to match slice indices.
		startStackIndex--
		endStackIndex--
		startStack := &stacks[startStackIndex]
		endStack := &stacks[endStackIndex]
		for i := 0; i < amountToMove; i++ {
			inMoveElement := startStack.pop()
			endStack.put(inMoveElement)
		}
	}
	for _, s := range stacks {
		fmt.Print(s.peek())
	}

}

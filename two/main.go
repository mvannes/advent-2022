package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Shape string

const (
	Rock     Shape = "rock"
	Paper          = "Paper"
	Scissors       = "Scissors"
)

func (s Shape) scoreAgainst(opponent Shape) int {
	return shapeResultMap[s][opponent]
}
func (s Shape) selfScore() int {
	return chosenShapeScoreMap[s]
}

var shapeResultMap = map[Shape]map[Shape]int{
	Rock:     {Rock: 3, Paper: 0, Scissors: 6},
	Paper:    {Rock: 6, Paper: 3, Scissors: 0},
	Scissors: {Rock: 0, Paper: 6, Scissors: 3},
}

var chosenShapeScoreMap = map[Shape]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}
var tokenToShapeMap = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err)
	}

	var totalScore int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		actions := strings.Split(line, " ")
		opponent := tokenToShapeMap[actions[0]]
		me := tokenToShapeMap[actions[1]]

		score := me.selfScore() + me.scoreAgainst(opponent)
		totalScore += score
	}
	fmt.Println(totalScore)
}

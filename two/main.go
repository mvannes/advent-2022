package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Result string

const (
	Win  Result = "win"
	Draw        = "draw"
	Lose        = "lose"
)

func (r Result) score() int {
	switch r {
	case Win:
		return 6
	case Draw:
		return 3
	case Lose:
		return 0
	default:
		return -1
	}
}

type Shape string

const (
	Rock     Shape = "rock"
	Paper          = "Paper"
	Scissors       = "Scissors"
)

func (s Shape) match(opponent Shape) Result {
	return shapeResultMap[s][opponent]
}

func (s Shape) selfScore() int {
	return chosenShapeScoreMap[s]
}

func (s Shape) findOpponentForResult(r Result) (Shape, error) {
	for o, res := range shapeResultMap[s] {
		if res == r {
			return o, nil
		}
	}
	return "", errors.New("Invalid result for this shape.")
}

var shapeResultMap = map[Shape]map[Shape]Result{
	Rock:     {Rock: Draw, Paper: Lose, Scissors: Win},
	Paper:    {Rock: Win, Paper: Draw, Scissors: Lose},
	Scissors: {Rock: Lose, Paper: Win, Scissors: Draw},
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

var tokenToResultMap = map[string]Result{
	"X": Win,
	"Y": Draw,
	"Z": Lose,
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}

	var partOneTotal int
	var totalScore int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		actions := strings.Split(line, " ")
		opponent := tokenToShapeMap[actions[0]]
		partOneMe := tokenToShapeMap[actions[1]]
		expectedResult := tokenToResultMap[actions[1]]
		me, err := opponent.findOpponentForResult(expectedResult)
		if nil != err {
			log.Fatal(err.Error())
		}
		score := me.selfScore() + me.match(opponent).score()
		totalScore += score
		partOneTotal += (partOneMe.selfScore() + partOneMe.match(opponent).score())
	}
	fmt.Println(partOneTotal)
	fmt.Println(totalScore)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const startingCoordinates = 0
const amountOfRopeParts = 10

type Coordinates struct {
	X, Y int
}

type RopePart struct {
	CurrentCoordinates Coordinates
	Next               *RopePart
	VisitedCoordinates []Coordinates
}

func (r *RopePart) Move(direction string) {
	c := Coordinates{
		X: r.CurrentCoordinates.X,
		Y: r.CurrentCoordinates.Y,
	}
	switch direction {
	case "U":
		c.Y++
		break
	case "D":
		c.Y--
		break
	case "L":
		c.X--
		break
	case "R":
		c.X++
		break
	case "Stay":
		break
	}
	r.MoveToCoordinates(c)
}
func (r *RopePart) MoveToCoordinates(c Coordinates) {
	r.CurrentCoordinates = c
	r.VisitedCoordinates = append(r.VisitedCoordinates, c)

	if r.Next == nil {
		return
	}

	distanceFromNext := chebyshev(r.CurrentCoordinates, r.Next.CurrentCoordinates)
	if distanceFromNext > 1 {
		newCoords := Coordinates{
			X: r.Next.CurrentCoordinates.X,
			Y: r.Next.CurrentCoordinates.Y,
		}

		if r.CurrentCoordinates.Y > newCoords.Y {
			newCoords.Y++
		} else if r.CurrentCoordinates.Y < newCoords.Y {
			newCoords.Y--
		}

		if r.CurrentCoordinates.X > newCoords.X {
			newCoords.X++
		} else if r.CurrentCoordinates.X < newCoords.X {
			newCoords.X--
		}
		r.Next.MoveToCoordinates(newCoords)
	}
}

func chebyshev(a, b Coordinates) int {
	var xDistance int
	if a.X > b.X {
		xDistance = a.X - b.X
	} else {
		xDistance = b.X - a.X
	}
	var yDistance int
	if a.Y > b.Y {
		yDistance = a.Y - b.Y
	} else {
		yDistance = b.Y - a.Y
	}

	if xDistance > yDistance {
		return xDistance
	}

	return yDistance
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatalln(err.Error())
	}
	s := bufio.NewScanner(f)
	ropeParts := []*RopePart{}

	for i := 0; i < amountOfRopeParts; i++ {
		part := &RopePart{
			CurrentCoordinates: Coordinates{
				X: startingCoordinates,
				Y: startingCoordinates,
			},
			VisitedCoordinates: []Coordinates{},
		}
		if i != 0 {
			part.Next = ropeParts[i-1]
		}
		part.Move("Stay")
		ropeParts = append(ropeParts, part)
	}

	head := ropeParts[len(ropeParts)-1]

	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " ")
		// X is horizontal
		// Y is vertical
		direction := parts[0]
		count, err := strconv.Atoi(parts[1])
		if nil != err {
			log.Fatalln(err.Error())
		}

		for i := 0; i < count; i++ {
			head.Move(direction)
		}
	}

	var uniqueVisits []Coordinates
	for _, v := range ropeParts[0].VisitedCoordinates {
		found := false
		for _, u := range uniqueVisits {
			if u == v {
				found = true
				break
			}
		}
		if !found {
			uniqueVisits = append(uniqueVisits, v)
		}
	}
	fmt.Println(len(ropeParts[0].VisitedCoordinates))
	fmt.Println(len(uniqueVisits))
}

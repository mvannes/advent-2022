package main

import "C"
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

type Grid struct {
	innerGrid map[int]map[int]string
}

func (g *Grid) addCoordinate(c Coordinates, kind string) {
	if _, ok := g.innerGrid[c.Y]; !ok {
		g.innerGrid[c.Y] = map[int]string{}
	}
	g.innerGrid[c.Y][c.X] = kind
}

func (g *Grid) isCoordinateTaken(x, y int) bool {
	if _, ok := g.innerGrid[y]; !ok {
		return false
	}
	if _, ok := g.innerGrid[y][x]; !ok {
		return false
	}
	return g.innerGrid[y][x] != ""
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}

	s := bufio.NewScanner(f)
	grid := Grid{innerGrid: map[int]map[int]string{}}
	for s.Scan() {
		coordinateList := coordinateListFromString(s.Text())
		for i, c := range coordinateList {
			grid.addCoordinate(c, "ROCK")
			if i == 0 {
				continue
			}
			prevCoordinate := coordinateList[i-1]
			for _, lc := range coordinatesBetweenTwoCoordinates(c, prevCoordinate) {
				grid.addCoordinate(lc, "ROCK")
			}
		}
	}
	yMax := 0
	for y, _ := range grid.innerGrid {
		if y > yMax {
			yMax = y
		}
	}
	yMax += 2
	fmt.Println(yMax)
	sandCount := 0
	// Await overflow.
	overflow := false
	for !overflow {
		// Do create sand
		x := 500
		y := 0
		sandCount++
		for {
			if grid.isCoordinateTaken(500, 0) {
				fmt.Println("found")
				overflow = true
				break
			}
			if y == yMax-1 {
				grid.addCoordinate(Coordinates{
					X: x,
					Y: y,
				}, "SAND")
				break
			}
			if !grid.isCoordinateTaken(x, y+1) {
				// Can we move down?
				y = y + 1
				continue
			} else if !grid.isCoordinateTaken(x-1, y+1) {
				// Can we move diagonal left
				x = x - 1
				y = y + 1
				continue
			} else if !grid.isCoordinateTaken(x+1, y+1) {
				// Can we move diagonal right
				x = x + 1
				y = y + 1
				continue
			} else {
				grid.addCoordinate(Coordinates{
					X: x,
					Y: y,
				}, "SAND")
				break
			}
		}
	}
	fmt.Println(sandCount - 1)
}

func coordinatesBetweenTwoCoordinates(c, c2 Coordinates) []Coordinates {
	results := []Coordinates{}
	if c.Y < c2.Y {
		current := c.Y + 1
		for current < c2.Y {
			results = append(results, Coordinates{
				X: c.X,
				Y: current,
			})
			current++
		}
		return results
	}
	if c.Y > c2.Y {
		current := c.Y - 1
		for current > c2.Y {
			results = append(results, Coordinates{
				X: c.X,
				Y: current,
			})
			current--
		}
		return results
	}
	if c.X < c2.X {
		current := c.X + 1
		for current < c2.X {
			results = append(results, Coordinates{
				X: current,
				Y: c.Y,
			})
			current++
		}
		return results
	}
	if c.X > c2.X {
		current := c.X - 1
		for current > c2.X {
			results = append(results, Coordinates{
				X: current,
				Y: c.Y,
			})
			current--
		}
		return results
	}
	return results
}

func coordinateListFromString(line string) []Coordinates {
	res := []Coordinates{}
	for _, d := range strings.Split(line, " -> ") {
		coordinates := strings.Split(d, ",")
		x, err := strconv.Atoi(coordinates[0])
		if nil != err {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(coordinates[1])
		if nil != err {
			log.Fatal(err)
		}
		res = append(res, Coordinates{
			X: x,
			Y: y,
		})
	}
	return res
}

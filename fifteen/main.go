package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Beacon struct {
	X int
	Y int
}

type Sensor struct {
	X             int
	Y             int
	ClosestBeacon Beacon
}

type Range struct {
	minX int
	maxX int
}

func (r Range) contains(i int) bool {
	return r.minX <= i && i <= r.maxX
}

func (s Sensor) Manhattan() int {
	diffX := s.X - s.ClosestBeacon.X
	diffY := s.Y - s.ClosestBeacon.Y
	return abs(diffX) + abs(diffY)
}

func (s Sensor) IsPointInRange(x, y int) bool {
	return s.RangeForRow(y).contains(x)
}

func (s Sensor) RangeForRow(y int) Range {
	manhattan := s.Manhattan()
	diffY := abs(s.Y - y)
	maxDistance := manhattan - diffY
	if maxDistance < 0 {
		return Range{minX: math.MaxInt, maxX: math.MinInt}
	}

	return Range{
		minX: s.X - maxDistance,
		maxX: s.X + maxDistance,
	}
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return n * -1
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(f)
	}
	s := bufio.NewScanner(f)

	var sensors []Sensor
	for s.Scan() {
		l := s.Text()
		var sX, sY, bX, bY int
		_, err := fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &sX, &sY, &bX, &bY)
		if nil != err {
			log.Fatal(err)
		}

		sensors = append(sensors, Sensor{
			X: sX,
			Y: sY,
			ClosestBeacon: Beacon{
				X: bX,
				Y: bY,
			},
		})
	}
	inRangeCount := 0
	y := 2000000

	sX := math.MaxInt
	mX := math.MinInt
	for _, s := range sensors {
		if sX > s.X-s.Manhattan() {
			sX = s.X - s.Manhattan()
		}
		if mX < s.X+s.Manhattan() {
			mX = s.X + s.Manhattan()
		}
	}

	for x := sX; x <= mX; x++ {
		isBeacon := false
		for _, s := range sensors {
			if s.ClosestBeacon.X == x && s.ClosestBeacon.Y == y {
				isBeacon = true
				break
			}
		}
		if isBeacon {
			continue
		}
		for _, s := range sensors {
			if s.IsPointInRange(x, y) {
				inRangeCount++
				break
			}
		}
	}
	fmt.Println(inRangeCount)
}

func drawBeacon(smallestX, largestX, smallestY, largestY int, sensor Sensor) {
	for y := smallestY; y <= largestY; y++ {
		r := sensor.RangeForRow(y)
		for x := smallestX; x <= largestX; x++ {
			if x == smallestX {
				fmt.Print(y, " ")
			}
			if x == sensor.X && y == sensor.Y {
				fmt.Print("S")
			} else if r.contains(x) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()

	}
}

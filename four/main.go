package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}

	var numberOfFullIntersections int
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		schedules := strings.Split(line, ",")
		left := intRangeFromSchedule(schedules[0])
		right := intRangeFromSchedule(schedules[1])

		intersect := intersection(left, right)
		if len(intersect) == len(left) || len(intersect) == len(right) {
			numberOfFullIntersections++
		}
	}
	fmt.Println(numberOfFullIntersections)
}

func intRangeFromSchedule(schedule string) []int {
	bounds := strings.Split(schedule, "-")
	min, err := strconv.Atoi(bounds[0])
	if nil != err {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(bounds[1])
	if nil != err {
		log.Fatal(err)
	}
	rangeSlice := make([]int, max-min+1)
	for i := range rangeSlice {
		rangeSlice[i] = i + min
	}
	return rangeSlice
}

func intersection(a, b []int) []int {
	lookup := map[int]bool{}
	for _, aVal := range a {
		lookup[aVal] = true
	}
	var result []int
	for _, bVal := range b {
		if ok := lookup[bVal]; ok {
			result = append(result, bVal)
		}
	}

	return result
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Item struct {
	Type         string
	intValue     int
	sliceOfItems []Item
}

func (i Item) hasIndex(index int) bool {
	return len(i.sliceOfItems) > index
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}

	s := bufio.NewScanner(f)
	var set [2]Item
	var sets [][2]Item
	var counter int
	for s.Scan() {
		line := s.Bytes()
		if "" == string(line) {
			counter = 0
			sets = append(sets, set)
			continue
		}

		var content interface{}
		err := json.Unmarshal(line, &content)
		if err != nil {
			log.Fatal(err.Error())
		}
		set[counter] = createItem(content)
		counter++
	}

	sum := 0
	for i, s := range sets {
		correct := compareItems(s[0], s[1])
		fmt.Println(i)
		if correct == 1 {
			sum += i + 1
		}
		fmt.Println(s)
	}
	fmt.Println(sum)
}

func compareItems(a, b Item) int {
	if a.Type == "Int" && b.Type == "Int" {
		if a.intValue < b.intValue {
			return 1
		}
		if a.intValue > b.intValue {
			return -1
		}
		return 0
	}

	if a.Type != "Int" && b.Type == "Int" {
		newB := Item{Type: "ItemSlice", intValue: 0, sliceOfItems: []Item{b}}
		return compareItems(a, newB)
	}
	if a.Type == "Int" && b.Type != "Int" {
		newA := Item{
			Type:         "ItemSlice",
			intValue:     0,
			sliceOfItems: []Item{a},
		}
		return compareItems(newA, b)
	}

	// At this point, everything should be a slice.
	counter := 0
	for {
		if !a.hasIndex(counter) && b.hasIndex(counter) {
			return 1
		}
		if a.hasIndex(counter) && !b.hasIndex(counter) {
			return -1
		}

		if !a.hasIndex(counter) && !b.hasIndex(counter) {
			return 0
		}
		aI := a.sliceOfItems[counter]
		bI := b.sliceOfItems[counter]
		result := compareItems(aI, bI)
		if result != 0 {
			return result
		}

		counter++

	}
}

func createItem(i any) Item {
	var v int
	var vSlice []int
	var iSlice []interface{}
	bytes, err := json.Marshal(i)

	// int
	err = json.Unmarshal(bytes, &v)
	if nil == err {
		return Item{
			Type:         "Int",
			intValue:     v,
			sliceOfItems: nil,
		}
	}
	// Int slice
	err = json.Unmarshal(bytes, &vSlice)
	if err == nil {
		item := Item{
			Type:         "ItemSlice",
			intValue:     0,
			sliceOfItems: []Item{},
		}
		for _, i := range vSlice {
			intItem := Item{
				Type:         "Int",
				intValue:     i,
				sliceOfItems: nil,
			}
			item.sliceOfItems = append(item.sliceOfItems, intItem)
		}
		return item
	}

	// Slice of items
	err = json.Unmarshal(bytes, &iSlice)
	if err != nil {
		log.Fatal(err)
	}
	item := Item{
		Type:         "ItemSlice",
		intValue:     0,
		sliceOfItems: []Item{},
	}
	for _, i := range iSlice {
		iItem := createItem(i)
		item.sliceOfItems = append(item.sliceOfItems, iItem)
	}

	return item
}

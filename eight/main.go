package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Tree struct {
	size  int
	left  *Tree
	right *Tree
	up    *Tree
	down  *Tree
}

func treeByDirection(t *Tree, direction string) *Tree {
	switch direction {
	case "left":
		return t.left
	case "right":
		return t.right
	case "up":
		return t.up
	case "down":
		return t.down
	}
	return nil
}

func (t *Tree) isVisible() bool {
	for _, direction := range []string{"up", "down", "left", "right"} {
		sideTree := treeByDirection(t, direction)
		for {
			if sideTree == nil {
				return true
			}
			if sideTree.size >= t.size {
				break
			}
			sideTree = treeByDirection(sideTree, direction)
		}
	}
	return false
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatalln(err.Error())
	}
	s := bufio.NewScanner(f)

	grid := map[int][]*Tree{}

	var lineCount int
	for s.Scan() {
		line := s.Text()
		slice := strings.Split(line, "")
		for rowI, s := range slice {
			i, err := strconv.ParseInt(s, 10, 0)
			if nil != err {
				log.Fatalln(err.Error())
			}

			tree := &Tree{
				size:  int(i),
				left:  nil,
				right: nil,
				up:    nil,
				down:  nil,
			}

			// If its in the first column, the map key won't exist.
			if _, ok := grid[lineCount]; !ok {
				grid[lineCount] = []*Tree{tree}
				continue
			} else {
				grid[lineCount] = append(grid[lineCount], tree)
			}
			// If there is something before it, ensure the tree match
			if rowI-1 >= 0 {
				leftTree := grid[lineCount][rowI-1]
				tree.left = leftTree
				leftTree.right = tree
			}
			// If there is something above it, ensure the tree match.
			if lineCount-1 >= 0 {
				// Tree or dog?
				upTree := grid[lineCount-1][rowI]
				tree.up = upTree
				upTree.down = tree
			}
		}
		lineCount++
	}

	var wg sync.WaitGroup
	visibleTreeChan := make(chan int)

	for _, row := range grid {
		wg.Add(1)
		go func(c chan int, trees []*Tree) {
			defer wg.Done()
			for _, tree := range trees {
				if tree.isVisible() {
					c <- 1
				}
			}

		}(visibleTreeChan, row)
	}
	go func() {
		wg.Wait()
		close(visibleTreeChan)
	}()

	var visibleTrees int

	for _ = range visibleTreeChan {
		visibleTrees++
	}

	fmt.Println(visibleTrees)
}

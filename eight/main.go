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

func (t *Tree) calculateScenicScore() int {
	var scores []int
	for _, direction := range []string{"up", "down", "left", "right"} {
		var directionalScore int
		sideTree := treeByDirection(t, direction)
		for {
			if sideTree == nil {
				break
			}
			directionalScore++
			if sideTree.size >= t.size {
				break
			}
			sideTree = treeByDirection(sideTree, direction)
		}
		scores = append(scores, directionalScore)
	}
	score := scores[0]
	for _, ds := range scores[1:] {
		score = score * ds
	}
	return score
}

type TreeResult struct {
	visible     bool
	scenicScore int
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
	treeResultChan := make(chan TreeResult)
	for _, row := range grid {
		wg.Add(1)
		go func(c chan TreeResult, trees []*Tree) {
			defer wg.Done()
			for _, tree := range trees {
				c <- TreeResult{
					visible:     tree.isVisible(),
					scenicScore: tree.calculateScenicScore(),
				}
			}

		}(treeResultChan, row)
	}

	go func() {
		wg.Wait()
		close(treeResultChan)
	}()

	var visibleTrees int
	var mostScenicTree int
	for t := range treeResultChan {
		if t.visible {
			visibleTrees++
		}
		if t.scenicScore > mostScenicTree {
			mostScenicTree = t.scenicScore
		}
	}
	fmt.Println(visibleTrees)
	fmt.Println(mostScenicTree)
}

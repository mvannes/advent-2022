package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	size  int
	left  *Tree
	right *Tree
	up    *Tree
	down  *Tree
}

func (t Tree) isVisible() bool {
	left := t.left
	for {
		if left == nil {
			return true
		}
		if left.size >= t.size {
			break
		}
		left = left.left
	}
	right := t.right
	for {
		if right == nil {
			return true
		}
		if right.size >= t.size {
			break
		}
		right = right.right
	}

	up := t.up
	for {
		if up == nil {
			return true
		}
		if up.size >= t.size {
			break
		}
		up = up.up
	}

	down := t.down
	for {
		if down == nil {
			return true
		}
		if down.size >= t.size {
			break
		}
		down = down.down
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

	var visibleTrees int
	for _, row := range grid {
		for _, tree := range row {
			if tree.isVisible() {
				visibleTrees++
			}
		}
	}
	fmt.Println(visibleTrees)
}

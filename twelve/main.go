package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type Node struct {
	id        int
	Char      string
	Elevation int
	Edges     []Edge
}

type Edge struct {
	Node   *Node
	Weight int
}

func calcWeight(a, b *Node) int {
	if b.Elevation-1 > a.Elevation {
		return math.MaxInt64
	}
	return 1
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}

	s := bufio.NewScanner(f)
	nodes := map[int][]*Node{}
	nodeSlice := []*Node{}
	var lineCount int
	var nodeCount int
	for s.Scan() {
		line := s.Text()
		nodeStrings := strings.Split(line, "")
		for i, v := range nodeStrings {
			node := &Node{
				id:        nodeCount,
				Char:      v,
				Elevation: calculateElevation(v),
				Edges:     []Edge{},
			}
			// Get left and add it in.
			if i > 0 {
				leftNode := nodes[lineCount][i-1]
				leftNode.Edges = append(leftNode.Edges, Edge{
					Node:   node,
					Weight: calcWeight(leftNode, node),
				})
				node.Edges = append(node.Edges, Edge{
					Node:   leftNode,
					Weight: calcWeight(node, leftNode),
				})

			}
			// Get up and add it in.
			if lineCount > 0 {
				upNode := nodes[lineCount-1][i]
				upNode.Edges = append(upNode.Edges, Edge{
					Node:   node,
					Weight: calcWeight(upNode, node),
				})
				node.Edges = append(node.Edges, Edge{
					Node:   upNode,
					Weight: calcWeight(node, upNode),
				})
			}
			nodeCount++
			nodes[lineCount] = append(nodes[lineCount], node)
			nodeSlice = append(nodeSlice, node)
		}
		lineCount++
	}
	calculateTraveled(nodeSlice, "E")
}

type PrioNode struct {
	node   *Node
	weight int
	prevId int
}

type PrioQueue struct {
	nodes []PrioNode
}

func (p *PrioQueue) addNode(n *Node, weight int, prevId int) {
	p.nodes = append(p.nodes, PrioNode{
		node:   n,
		weight: weight,
		prevId: prevId,
	})
	sort.Slice(p.nodes, func(a, b int) bool {
		return p.nodes[b].weight > p.nodes[a].weight
	})
}

func (p *PrioQueue) hasMore() bool {
	return len(p.nodes) != 0
}

func (p *PrioQueue) pop() PrioNode {
	res := p.nodes[:1]
	p.nodes = p.nodes[1:]
	return res[0]
}

func calculateTraveled(nodes []*Node, characterToFind string) {
	var endNode *Node
	var visited = map[int]bool{}
	var dists = map[int]int{}
	for _, n := range nodes {
		if n.Char == characterToFind {
			endNode = n
		}

		dists[n.id] = math.MaxInt64
	}
	pq := PrioQueue{}
	pq.addNode(endNode, 0, -1)

	for pq.hasMore() {
		node := pq.pop()
		id := node.node.id
		if visited[id] {
			continue
		}
		visited[id] = true

		prevCost := 0
		if node.prevId != -1 && dists[node.prevId] != math.MaxInt64 {
			prevCost = dists[node.prevId]
		}

		if currentDist, ok := dists[id]; !ok {
			dists[id] = node.weight
		} else {
			if currentDist > node.weight {
				dists[id] = node.weight
			}
		}
		dists[id] += prevCost

		for _, e := range node.node.Edges {
			pq.addNode(e.Node, e.Weight, id)
		}
	}

	for _, n := range nodes {
		if n.Char == "E" {
			fmt.Println(dists[n.id])
			fmt.Println(n.id)
		}
	}
	fmt.Println("done!")

}

func calculateElevation(char string) int {
	if char == "S" {
		return int('a')
	}
	if char == "E" {
		return int('z')
	}

	return int(char[0])
}

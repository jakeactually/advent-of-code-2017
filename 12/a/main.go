package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	regex, _ := regexp.Compile("\\d+")
	g := simple.NewUndirectedGraph()

	for scanner.Scan() {
		strs := regex.FindAllString(scanner.Text(), -1)

		var nums []int64
		for _, chr := range strs {
			num, _ := strconv.Atoi(chr)
			nums = append(nums, int64(num))
		}

		node := nodeOrNew(g, nums[0])
		for _, num := range nums[1:] {
			neighbor := nodeOrNew(g, num)
			if node != neighbor {
				g.SetEdge(g.NewEdge(node, neighbor))
			}
		}
	}

	count := 0
	bf := traverse.BreadthFirst{
		Visit: func(n graph.Node) {
			count++
		},
	}
	bf.Walk(g, nodeOrNew(g, 0), nil)
	fmt.Println(count)
}

type Node struct {
	id int64
}

func (node Node) ID() int64 {
	return node.id
}

func nodeOrNew(g *simple.UndirectedGraph, id int64) graph.Node {
	node := g.Node(id)
	if node == nil {
		newNode := Node{id}
		g.AddNode(newNode)
		return newNode
	}
	return node
}

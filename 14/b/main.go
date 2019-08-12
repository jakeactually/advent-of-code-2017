package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

const size = 256
const key = "uugsqrei"

var suffix = []int{17, 31, 73, 47, 23}

func main() {
	var bits [128][128]int

	for y := 0; y < 128; y++ {
		entry := fmt.Sprint(key, "-", y)

		var bytes []int
		for _, chr := range entry {
			bytes = append(bytes, int(chr))
		}
		bytes = append(bytes, suffix...)

		for i, col := range hash(bytes) {
			for j := 0; j < 8; j++ {
				bits[y][i*8+j] = col >> uint(7-j) & 1
			}
		}
	}

	g := simple.NewUndirectedGraph()
	for y := int64(0); y < 128; y++ {
		for x := int64(0); x < 128; x++ {
			if bits[y][x] == 1 {
				node := nodeOrNew(g, x, y)

				if y < 127 && bits[y+1][x] == 1 {
					neighbor := nodeOrNew(g, x, y+1)
					g.SetEdge(g.NewEdge(node, neighbor))
				}

				if x < 127 && bits[y][x+1] == 1 {
					neighbor := nodeOrNew(g, x+1, y)
					g.SetEdge(g.NewEdge(node, neighbor))
				}
			}
		}
	}

	fmt.Println(len(topo.ConnectedComponents(g)))
}

func hash(lengths []int) [16]int {
	var array [size]int
	for i := 0; i < size; i++ {
		array[i] = i
	}

	offset := 0
	skip := 0

	for r := 0; r < 64; r++ {
		for _, length := range lengths {
			for i := 0; i < length/2; i++ {
				src := (offset + i) % size
				dst := (offset + length - i - 1) % size

				temp := array[src]
				array[src] = array[dst]
				array[dst] = temp
			}

			offset = (offset + length + skip) % size
			skip++
		}
	}

	var hash [16]int
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			hash[i] ^= array[i*16+j]
		}
	}

	return hash
}

func nodeOrNew(g *simple.UndirectedGraph, x, y int64) graph.Node {
	node := g.Node(y*128 + x)
	if node == nil {
		newNode := Node{x, y}
		g.AddNode(newNode)
		return newNode
	}
	return node
}

type Node struct {
	x, y int64
}

func (node Node) ID() int64 {
	return node.y*128 + node.x
}

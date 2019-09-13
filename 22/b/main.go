package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var quadrant1 quadrant
	size := 0
	for scanner.Scan() {
		quadrant1 = append(quadrant1, []state{})
		size++

		line := scanner.Text()

		for _, char := range line {
			st := clean

			if char == '#' {
				st = infected
			}

			quadrant1[size-1] = append(quadrant1[size-1], st)
		}
	}

	quadrants := []quadrant{
		quadrant1,
		quadrant{{}},
		quadrant{{}},
		quadrant{{}},
	}

	gy := len(quadrant1) / 2
	gx := len(quadrant1[0]) / 2

	var d direction
	count := 0
	for i := 0; i < 10000000; i++ {
		index := getQuadrant(gx, gy)
		q := quadrants[index]
		x, y := abs(gx), abs(gy)
		q = q.realloc(x, y)

		switch q[y][x] {
		case clean:
			d = d.left()
		case weakened:
			count++
		case infected:
			d = d.right()
		case flagged:
			d = d.reverse()
		}

		q[y][x] = q[y][x].next()

		ox, oy := d.toSteps()
		gx, gy = gx+ox, gy+oy
		quadrants[index] = q
	}

	fmt.Println(count)
}

type direction int

const (
	up    direction = 0
	right direction = 1
	down  direction = 2
	left  direction = 3
)

func (d direction) left() direction {
	return (d + 4 - 1) % 4
}

func (d direction) right() direction {
	return (d + 1) % 4
}

func (d direction) reverse() direction {
	return (d + 2) % 4
}

type state int

const (
	clean    state = 0
	weakened state = 1
	infected state = 2
	flagged  state = 3
)

func (s state) next() state {
	return (s + 1) % 4
}

func (d direction) toSteps() (int, int) {
	switch d {
	case up:
		return 0, -1
	case right:
		return 1, 0
	case down:
		return 0, 1
	case left:
		return -1, 0
	default:
		return 0, 0
	}
}

func (d direction) toString() string {
	switch d {
	case up:
		return "up"
	case right:
		return "right"
	case down:
		return "down"
	case left:
		return "left"
	default:
		return ""
	}
}

type quadrant [][]state

func (q quadrant) realloc(x, y int) quadrant {
	dy := y - len(q) + 1

	for i := 0; i < dy; i++ {
		q = append(q, []state{})
	}

	dx := x - len(q[y]) + 1

	for i := 0; i < dx; i++ {
		q[y] = append(q[y], clean)
	}

	return q
}

func getQuadrant(x, y int) int {
	if x < 0 {
		if y < 0 {
			return 2
		}

		return 1
	}

	if y < 0 {
		return 3
	}

	return 0
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

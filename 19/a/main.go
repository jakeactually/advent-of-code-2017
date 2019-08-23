package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	x, y := 1, 0
	var d dir

	for {
		char := lines[y][x]

		if char >= 65 && char <= 90 {
			fmt.Print(string(char))
		}

		if char == '+' {
			if d == down || d == up {
				if lines[y][x-1] == '-' {
					d = left
				} else {
					d = right
				}
			} else {
				if lines[y-1][x] == '|' {
					d = up
				} else {
					d = down
				}
			}
		}

		switch d {
		case down:
			y++
		case left:
			x--
		case up:
			y--
		case right:
			x++
		}

		if x < 0 || y < 0 || x > 200 || y > 200 {
			break
		}
	}
}

type dir int

const (
	down  dir = 0
	left  dir = 1
	up    dir = 2
	right dir = 3
)

package main

import (
	"fmt"
	"math"
)

// based on my own gist
// https://gist.github.com/jakeactually/4325133c5d66384c4be944d1d0ec7111

const goal int = 347991

func main() {
	size := int(math.Ceil(math.Sqrt(float64(goal))))

	for y := range make([]int, size) {
		for x := range make([]int, size) {
			result := plot(size, x, y)
			// fmt.Printf("%2d ", result)

			if result == goal {
				str := "result %d at a x %d y %d manhathan %d"
				center := size / 2
				fmt.Printf(str, result, x, y, manhathan(center, center, x, y))
			}
		}
		// fmt.Println()
	}
}

func manhathan(x1, y1, x2, y2 int) int {
	m := math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
	return int(m)
}

func plot(size, x, y int) int {
	l := getLevel(size, x, y)*2 + 1
	l_ := l - 1
	z := getZone(size, x, y)
	mapping := getMapping(size, l_, z, x, y)
	return l*l - (z-1)*l_ - mapping
}

func getLevel(size, x, y int) int {
	half := size / 2
	dx, dy := dist(half, half, x, y)
	if dx > dy {
		return dx
	}
	return dy
}

func dist(x1, y1, x2, y2 int) (int, int) {
	return int(math.Abs(float64(x1 - x2))), int(math.Abs(float64(y1 - y2)))
}

func getZone(size, x, y int) int {
	if y <= size/2 {
		if y-x > 0 {
			return 2
		} else if size-y-2 < x {
			return 4
		} else {
			return 3
		}
	} else {
		if size-y > x {
			return 2
		} else if x-y > 0 {
			return 4
		} else {
			return 1
		}
	}
}

func getMapping(size, level, zone, x, y int) int {
	margin := (size - level) / 2

	switch zone {
	case 1:
		return level - x + margin
	case 2:
		return level - y + margin
	case 3:
		return x - margin
	case 4:
		return y - margin
	default:
		return 0
	}
}

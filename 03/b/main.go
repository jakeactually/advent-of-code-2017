package main

import (
	"fmt"
)

func main() {
	const trySize = 9
	const limit = trySize * trySize
	const size = trySize + 2

	var spiral [size][size]int
	x, y := size/2, size/2
	spiral[y][x] = 1
	i := 2

	direction := 1
	length := 1
	done := false

	for !done {
		for j := 0; j < length; j++ {
			if i > limit {
				done = true
				break
			}

			switch direction {
			case 0:
				y++
			case 1:
				x++
			case 2:
				y--
			case 3:
				x--
			default:
			}

			if i == 1 {
				continue
			}

			spiral[y][x] = spiral[y+1][x] +
				spiral[y][x+1] +
				spiral[y+1][x+1] +
				spiral[y-1][x] +
				spiral[y][x-1] +
				spiral[y-1][x-1] +
				spiral[y+1][x-1] +
				spiral[y-1][x+1]

			i++
		}

		direction = (direction + 1) % 4

		if direction == 3 || direction == 1 {
			length++
		}
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fmt.Printf("%2d ", spiral[y][x])
		}
		fmt.Println()
	}
}

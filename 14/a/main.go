package main

import (
	"fmt"
)

const size = 256
const key = "uugsqrei"

var suffix = []int{17, 31, 73, 47, 23}

func main() {
	sum := 0

	for i := 0; i < 128; i++ {
		entry := fmt.Sprint(key, "-", i)

		var bytes []int
		for _, chr := range entry {
			bytes = append(bytes, int(chr))
		}
		bytes = append(bytes, suffix...)

		for _, col := range hash(bytes) {
			for j := 0; j < 8; j++ {
				sum += col >> uint(j) & 1
			}
		}
	}

	fmt.Println(sum)
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

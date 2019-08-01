package main

import (
	"fmt"
)

const size = 256

func main() {
	var array [size]int
	for i := 0; i < size; i++ {
		array[i] = i
	}

	input := "192,69,168,160,78,1,166,28,0,83,198,2,254,255,41,12"
	suffix := []int{17, 31, 73, 47, 23}
	var lengths []int
	for _, char := range input {
		lengths = append(lengths, int(char))
	}
	lengths = append(lengths, suffix...)

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
		fmt.Printf("%02x", hash[i])
	}
}

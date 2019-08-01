package main

import "fmt"

const size = 256

func main() {
	var array [size]int
	lengths := []int{192, 69, 168, 160, 78, 1, 166, 28, 0, 83, 198, 2, 254, 255, 41, 12}

	for i := 0; i < size; i++ {
		array[i] = i
	}

	offset := 0
	skip := 0
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

	fmt.Println(array[0] * array[1])
}

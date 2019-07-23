package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	// Remove newline
	input = input[:len(input)-1]
	length := len(input)
	numbers := make([]int, length)

	for i, char := range input {
		number, _ := strconv.Atoi(string(char))
		numbers[i] = number
	}

	sum := 0
	half := length / 2

	for i, number := range numbers {
		if number == numbers[(i+half)%length] {
			sum += number
		}
	}

	fmt.Println(sum)
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	text, _ := ioutil.ReadFile("input.txt")
	trimmed := strings.Trim(string(text), "\n")
	lines := strings.Split(trimmed, "\n")
	length := len(lines)
	numbers := make([]int, length)

	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	i := 0
	c := 0

	for i < length {
		to := numbers[i]
		numbers[i]++
		i += to
		c++
	}

	fmt.Println(c)
}

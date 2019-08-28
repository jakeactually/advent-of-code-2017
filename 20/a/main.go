package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

// the answer is
// the one with the less acceleration
// if many
// the one whose speed is less in sync with that acceleration

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("\\d+")
	min := 100.0
	var i int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := r.FindAllString(line, -1)
		accel := numbers[6:9]

		x, _ := strconv.Atoi(accel[0])
		y, _ := strconv.Atoi(accel[1])
		z, _ := strconv.Atoi(accel[2])

		manhathan := math.Abs(float64(x)) +
			math.Abs(float64(y)) +
			math.Abs(float64(z))

		if manhathan <= min {
			fmt.Println(line, i)
			min = manhathan
		}

		i++
	}
}

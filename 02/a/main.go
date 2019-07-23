package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		min := 10000
		max := 0

		for _, number := range strings.Split(scanner.Text(), "\t") {
			n, _ := strconv.Atoi(number)

			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}

		sum += max - min
	}

	fmt.Println(sum)
}

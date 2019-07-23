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
		strs := strings.Split(scanner.Text(), "\t")
		nums := make([]int, len(strs))
		for i, str := range strs {
			num, _ := strconv.Atoi(str)
			nums[i] = num
		}

		for _, x := range nums {
			for _, y := range nums {
				if x != y && x%y == 0 {
					sum += x / y
				}
			}
		}
	}

	fmt.Println(sum)
}

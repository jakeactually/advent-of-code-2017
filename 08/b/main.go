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

	regs := make(map[string]int)
	max := 0

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		b, _ := strconv.Atoi(words[6])

		if cond(regs[words[4]], words[5], b) {
			sum, _ := strconv.Atoi(words[2])

			if words[1] == "inc" {
				regs[words[0]] += sum
			} else {
				regs[words[0]] -= sum
			}

			if regs[words[0]] > max {
				max = regs[words[0]]
			}
		}
	}

	fmt.Println(max)
}

func cond(reg int, op string, num int) bool {
	switch op {
	case "<":
		return reg < num
	case "<=":
		return reg <= num
	case ">":
		return reg > num
	case ">=":
		return reg >= num
	case "==":
		return reg == num
	case "!=":
		return reg != num
	default:
		return false
	}
}

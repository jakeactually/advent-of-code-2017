package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	valids := 0

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		set := make(map[string]int)
		valid := true

		for _, word := range words {
			set[word] += 1
		}

		for _, value := range set {
			if value > 1 {
				valid = false
			}
		}

		if valid {
			valids += 1
		}
	}

	fmt.Println(valids)
}

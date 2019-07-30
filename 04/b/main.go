package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
			set[sortString(word)] += 1
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

func sortString(str string) string {
	chars := []rune(str)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const size int = 16

func main() {
	text, _ := ioutil.ReadFile("input.txt")
	trimmed := strings.Trim(string(text), "\n")
	words := strings.Split(trimmed, "\t")

	banks := [size]int{}
	for i, word := range words {
		banks[i], _ = strconv.Atoi(word)
	}

	set := make(map[[size]int]bool)
	done := false
	pointer := 0
	c := 0

	for !done {
		max := 0
		for i, bank := range banks {
			if bank > max {
				pointer, max = i, bank
			}
		}

		banks[pointer] = 0
		for max > 0 {
			pointer++
			max--
			banks[pointer%size] += 1
		}

		if !set[banks] {
			set[banks] = true
		} else {
			done = true
		}

		c++
	}

	fmt.Println(c)
}

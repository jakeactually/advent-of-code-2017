package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Disc struct {
	name     string
	children []string
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var discs []Disc
	r, _ := regexp.Compile("[a-z]+")
	for _, line := range lines {
		words := r.FindAllString(line, -1)
		discs = append(discs, Disc{words[0], words[1:]})
	}

	set := make(map[string]bool)
	for _, disc := range discs {
		for _, child := range disc.children {
			set[child] = true
		}
	}

	for _, disc := range discs {
		if !set[disc.name] {
			fmt.Println(disc.name)
		}
	}
}

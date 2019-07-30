package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Disc struct {
	name      string
	weight    int
	hasParent bool
	children  []*Disc
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// extract data
	r, _ := regexp.Compile("\\w+")
	var list [][]string
	for _, line := range lines {
		data := r.FindAllString(line, -1)
		list = append(list, data)
	}

	// make refs dictionary
	dic := make(map[string]*Disc)
	for _, data := range list {
		name := data[0]
		weight, _ := strconv.Atoi(data[1])
		dic[name] = &Disc{name, weight, false, nil}
	}

	// link children
	for _, data := range list {
		name := data[0]
		childrensNames := data[2:]

		disc := dic[name]
		for _, cn := range childrensNames {
			child := dic[cn]
			child.hasParent = true
			disc.children = append(disc.children, child)
		}
	}

	// get bottom
	var bottom *Disc
	for _, disc := range dic {
		if !disc.hasParent {
			bottom = disc
		}
	}

	// solution
	dic["marnqj"].weight -= 8

	printDisc(bottom, 0)
	fmt.Println()
	fmt.Println(dic["marnqj"].weight)
}

func printDisc(disc *Disc, level int) {
	weight := disc.weight

	last := 0
	if disc.children != nil {
		last = getWeight(disc.children[0])
	}

	for _, child := range disc.children {
		printDisc(child, level+1)

		w := getWeight(child)
		if w != last {
			fmt.Print("********************diff!")
		}
		last = w
		weight += w

		fmt.Println()
	}

	fmt.Printf(
		"%s%s %d %d",
		strings.Repeat("|", level),
		disc.name,
		disc.weight,
		weight,
	)
}

func getWeight(disc *Disc) int {
	weight := disc.weight
	for _, child := range disc.children {
		weight += getWeight(child)
	}
	return weight
}

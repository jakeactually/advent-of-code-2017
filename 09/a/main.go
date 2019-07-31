package main

import (
	"fmt"
	"io/ioutil"
)

type Group struct {
	children []Group
}

func main() {
	text, _ := ioutil.ReadFile("input.txt")
	g, _ := parse([]rune(string(text)))
	fmt.Println(getValue(g, 1))
}

func parse(str []rune) (Group, []rune) {
	rest := str[1:]

	var children []Group
	var child Group
	garbageState := false

	for {
	loop:
		for i, char := range rest {
			switch char {
			case '!':
				rest = rest[i+2:]
				break loop
			case '<':
				garbageState = true
			case '>':
				garbageState = false
			case '{':
				if !garbageState {
					child, rest = parse(rest[i:])
					children = append(children, child)
					break loop
				}
			case '}':
				if !garbageState {
					return Group{children}, rest[i+1:]
				}
			}
		}
	}
}

func getValue(group Group, level int) int {
	value := level
	for _, child := range group.children {
		value += getValue(child, level+1)
	}
	return value
}

func printGroup(group Group) {
	fmt.Print("{")
	for _, child := range group.children {
		printGroup(child)
	}
	fmt.Print("}")
}

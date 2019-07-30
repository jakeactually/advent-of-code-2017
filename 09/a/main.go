package main

import (
	"fmt"
)

type Group struct {
	children []*Group
}

func main() {
	//str := "{{<a!>},{<a!>},{<a!>},{<ab>}}"
	str := "{2¿{sdf}i{j}ii}"
	a, _ := parse([]rune(str))
	printGroup(a)
}

func parse(str []rune) (*Group, []rune) {
	var children []*Group
	var child *Group
	rest := str

	// eat bracket
	if len(str) == 0 || str[0] != '{' {
		return nil, str
	}
	rest = rest[1:]

	// eat any
	index := 0
	for i, char := range rest {
		if char == '{' {
			index = i
			break
		}
		if char == '}' {
			return &Group{children}, rest[i:]
		}
	}
	rest = rest[index:]

	// children
	for {
		child, rest = parse(rest)
		if child == nil {
			break
		}
		children = append(children, child)
	}

	// eat bracket
	if len(rest) == 0 || rest[0] != '}' {
		return nil, rest
	}
	rest = rest[1:]

	return &Group{children}, rest
}

func printGroup(group *Group) {
	fmt.Print("{")
	for _, child := range group.children {
		printGroup(child)
	}
	fmt.Print("}")
}

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
	rest := []rune(string(text))
	skipState := false
	garbageState := false
	garbage := 0

	for _, char := range rest {
		if skipState {
			skipState = false
		} else if char == '!' {
			skipState = true
		} else if char == '>' {
			garbageState = false
		} else if garbageState {
			garbage++
		} else if char == '<' {
			garbageState = true
		}
	}

	fmt.Println(garbage)
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	hall := []rune("abcdefghijklmnop")
	bytes, _ := ioutil.ReadFile("input.txt")
	entries := strings.Split(string(bytes), ",")

	for _, entry := range entries {
		data := entry[1:]

		switch entry[0] {
		case 's':
			offset, _ := strconv.Atoi(data)
			hall = spin(hall, offset)
		case 'x':
			pair := strings.Split(data, "/")
			a, _ := strconv.Atoi(pair[0])
			b, _ := strconv.Atoi(pair[1])
			exchange(hall, a, b)
		case 'p':
			a := indexOf(hall, rune(data[0]))
			b := indexOf(hall, rune(data[2]))
			exchange(hall, a, b)
		}
	}

	fmt.Println(string(hall))
}

func spin(runes []rune, offset int) []rune {
	temp := append([]rune(nil), runes...)

	for i := 0; i < 16; i++ {
		temp[(i+offset)%16] = runes[i]
	}

	return temp
}

func exchange(runes []rune, a, b int) {
	temp := runes[a]
	runes[a] = runes[b]
	runes[b] = temp
}

func indexOf(runes []rune, r1 rune) int {
	for i, r2 := range runes {
		if r1 == r2 {
			return i
		}
	}

	return -1
}

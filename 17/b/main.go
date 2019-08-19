package main

import "fmt"

func main() {
	posOne := 0
	pos := 0

	for i := 1; i < 50000001; i++ {
		pos = (pos+359)%i + 1

		if pos == 1 {
			posOne = i
		}
	}

	fmt.Println(posOne)
}

package main

import "fmt"

func main() {
	c := 0

	for i := 0; i < 1001; i++ {
		b := 109300 + 17*i

		for j := 2; j < b; j++ {
			if b%j == 0 {
				c++
				break
			}
		}
	}

	fmt.Println(c)
}

package main

import "fmt"

func main() {
	a := 883
	b := 879
	c := 0

	for i := 0; i < 40000000; i++ {
		a = a * 16807 % 2147483647
		b = b * 48271 % 2147483647

		if a&0xffff == b&0xffff {
			c++
		}
	}

	fmt.Println(c)
}

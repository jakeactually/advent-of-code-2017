package main

import "fmt"

func main() {
	a := Gen{16807, 4, 883}
	b := Gen{48271, 8, 879}
	c := 0

	for i := 0; i < 5000000; i++ {
		if a.next()&0xffff == b.next()&0xffff {
			c++
		}
	}

	fmt.Println(c)
}

type Gen struct {
	seed  int
	test  int
	value int
}

func (gen *Gen) next() int {
	for {
		gen.value = gen.value * gen.seed % 2147483647

		if gen.value%gen.test == 0 {
			break
		}
	}
	return gen.value
}

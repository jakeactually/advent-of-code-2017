package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	text := string(file)
	jumps := strings.Split(text, ",")

	var x, y, z, max int

	for _, jump := range jumps {
		switch jump {
		case "n":
			x--
			y++
		case "ne":
			x--
			z++
		case "se":
			y--
			z++
		case "s":
			x++
			y--
		case "sw":
			x++
			z--
		case "nw":
			y++
			z--
		}

		man := math.Abs(float64(x)) +
			math.Abs(float64(y)) +
			math.Abs(float64(z))

		if int(man) > max {
			max = int(man)
		}
	}

	fmt.Println(max / 2)
}

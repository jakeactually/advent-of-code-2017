package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	regex, _ := regexp.Compile("\\d+")
	var layers []Layer

	for scanner.Scan() {
		strs := regex.FindAllString(scanner.Text(), -1)
		depth, _ := strconv.Atoi(strs[0])
		rng, _ := strconv.Atoi(strs[1])
		layers = append(layers, Layer{depth, rng, 0, true})
	}

	delay := 0
	for {
		var copy []Layer
		copy = append(copy, layers...)

		if trespass(delay, copy) {
			break
		}

		for j := range layers {
			layers[j].step()
		}

		delay++
	}

	fmt.Println(delay)
}

func trespass(delay int, layers []Layer) bool {
	for x := 0; x < 98+1; x++ {
		for j := range layers {
			layer := &layers[j]
			if x == layer.depth && layer.pos == 0 {
				return false
			}
			layer.step()
		}
	}

	return true
}

type Layer struct {
	depth int
	rng   int
	pos   int
	down  bool
}

func (layer *Layer) step() {
	if layer.down {
		layer.pos++
		if layer.pos == layer.rng-1 {
			layer.down = false
		}
	} else {
		layer.pos--
		if layer.pos == 0 {
			layer.down = true
		}
	}
}

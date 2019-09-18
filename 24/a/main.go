package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var cs []component
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "/")
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])
		cs = append(cs, component{a, b, false})
	}

	var all []path
	for i, c := range cs {
		if c.a == 0 || c.b == 0 {
			cs2 := append([]component{}, cs...)
			cs2[i].dirty = true

			var c2 component
			if c.a == 0 {
				c2 = cs2[i]
			} else {
				c2 = cs2[i].flip()
			}

			p := path{cs2, []component{c2}, c2}
			all = append(all, p)
		}
	}
	prev := append([]path{}, all...)

	for {
		var next []path

		for _, p := range prev {
			next = append(next, p.out()...)
		}

		if len(next) == 0 {
			break
		}

		all = append(all, next...)
		prev = next
	}

	max := 0
	for _, p := range all {
		strength := p.strength()

		if strength > max {
			max = strength
		}
	}

	fmt.Println(max)
}

type path struct {
	available []component
	list      []component
	last      component
}

func (p path) strength() int {
	sum := 0

	for _, c := range p.list {
		sum += c.a + c.b
	}

	return sum
}

func (p path) out() []path {
	var paths []path

	for i, c := range p.available {
		if !c.dirty {
			if p.joins(c) {
				paths = append(paths, p.add(i))
			}
			if c.a != c.b && p.joins(c.flip()) {
				paths = append(paths, p.addFlip(i))
			}
		}
	}

	return paths
}

func (p path) joins(c component) bool {
	return p.last.joins(c)
}

func (p path) add(index int) path {
	pr := p.copy()
	pr.available[index].dirty = true
	pr.list = append(pr.list, pr.available[index])
	pr.last = pr.available[index]
	return pr
}

func (p path) addFlip(index int) path {
	pr := p.copy()
	pr.available[index].dirty = true
	pr.list = append(pr.list, pr.available[index].flip())
	pr.last = pr.available[index].flip()
	return pr
}

func (p path) copy() path {
	available := append([]component{}, p.available...)
	list := append([]component{}, p.list...)

	return path{available, list, p.last}
}

type component struct {
	a, b  int
	dirty bool
}

func (c component) joins(c2 component) bool {
	return c.b == c2.a
}

func (c component) flip() component {
	return component{c.b, c.a, c.dirty}
}

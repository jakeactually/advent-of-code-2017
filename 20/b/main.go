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

	r, _ := regexp.Compile("[-\\d]+")
	var particles []particle
	for scanner.Scan() {
		line := scanner.Text()
		numbers := r.FindAllString(line, -1)

		var parsed []int
		for _, n := range numbers {
			p, _ := strconv.Atoi(n)
			parsed = append(parsed, p)
		}

		p := particle{
			vec3{parsed[0], parsed[1], parsed[2]},
			vec3{parsed[3], parsed[4], parsed[5]},
			vec3{parsed[6], parsed[7], parsed[8]},
		}
		particles = append(particles, p)
	}

	min := 1000
	for {
		count := make(map[vec3]int)

		for i := range particles {
			particles[i].move()
			count[particles[i].pos]++
		}

		var newParticles []particle
		for _, p := range particles {
			if count[p.pos] < 2 {
				newParticles = append(newParticles, p)
			}
		}
		particles = newParticles

		if len(particles) < min {
			min = len(particles)
			fmt.Println(min)
		}
	}
}

type vec3 struct {
	x, y, z int
}

func (v1 vec3) add(v2 vec3) vec3 {
	return vec3{
		v1.x + v2.x,
		v1.y + v2.y,
		v1.z + v2.z,
	}
}

type particle struct {
	pos, vel, ace vec3
}

func (p *particle) move() {
	p.vel = p.vel.add(p.ace)
	p.pos = p.pos.add(p.vel)
}

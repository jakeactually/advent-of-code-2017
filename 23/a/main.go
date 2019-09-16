package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	var contextArray [30]int
	context := contextArray[0:30]
	instructions := parseInstructions(file, context)
	pointer := 0
	muls := 0

	for {
		if pointer == len(instructions) {
			break
		}

		ins := instructions[pointer]

		switch ins.name {
		case "set":
			context[ins.left.val()] = ins.right.val()
		case "sub":
			context[ins.left.val()] -= ins.right.val()
		case "mul":
			context[ins.left.val()] *= ins.right.val()
			muls++
		case "jnz":
			if context[ins.left.val()] != 0 {
				pointer += ins.right.val() - 1
			}
		}

		pointer++
	}

	fmt.Println(muls)
}

func parseInstructions(r io.Reader, context []int) []instruction {
	scanner := bufio.NewScanner(r)
	var instructions []instruction

	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, " ")

		var second value
		{
			number, err := strconv.Atoi(elements[1])
			if err == nil {
				second = raw{number}
			} else {
				second = raw{int(elements[1][0]) - 97}
			}
		}

		var third value
		if len(elements) < 3 {
			third = none{}
		} else {
			number, err := strconv.Atoi(elements[2])
			if err == nil {
				third = raw{number}
			} else {
				third = register{int(elements[2][0] - 97), context}
			}
		}

		i := instruction{elements[0], second, third}
		instructions = append(instructions, i)
	}

	return instructions
}

type instruction struct {
	name  string
	left  value
	right value
}

type value interface {
	val() int
}

type register struct {
	key     int
	context []int
}

func (r register) val() int {
	return r.context[r.key]
}

type raw struct {
	value int
}

func (r raw) val() int {
	return r.value
}

type none struct {
}

func (r none) val() int {
	return 0
}

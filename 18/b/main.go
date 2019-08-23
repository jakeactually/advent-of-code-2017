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
	instructions := parseInstructions(file)

	var a, b program
	a.friend = &b
	a.registers[int('p')-97] = 0
	b.friend = &a
	b.registers[int('p')-97] = 1

	for {
		if !a.step(instructions) && !b.step(instructions) {
			break
		}
	}

	fmt.Println(b.sent)
}

func (p *program) step(instructions []instruction) bool {
	if p.pointer > len(instructions) {
		return false
	}

	ins := instructions[p.pointer]
	regs := &p.registers
	left, right := ins.left, ins.right

	switch ins.name {
	case "snd":
		p.friend.messages = append(p.friend.messages, regs[left.val(regs)])
		p.sent++
	case "set":
		regs[left.val(regs)] = right.val(regs)
	case "add":
		regs[left.val(regs)] += right.val(regs)
	case "mul":
		regs[left.val(regs)] *= right.val(regs)
	case "mod":
		regs[left.val(regs)] %= right.val(regs)
	case "rcv":
		if len(p.messages) > 0 {
			regs[left.val(regs)] = p.messages[0]
			p.messages = p.messages[1:]
		} else {
			return false
		}
	case "jgz":
		if regs[left.val(regs)] > 0 {
			p.pointer += right.val(regs) - 1
		}
	}

	p.pointer++
	return true
}

func parseInstructions(r io.Reader) []instruction {
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
				third = register{int(elements[2][0] - 97)}
			}
		}

		i := instruction{elements[0], second, third}
		instructions = append(instructions, i)
	}

	return instructions
}

type program struct {
	registers [30]int
	pointer   int
	messages  []int
	sent      int
	friend    *program
}

type instruction struct {
	name  string
	left  value
	right value
}

type value interface {
	val(registers *[30]int) int
}

type register struct {
	key int
}

func (r register) val(registers *[30]int) int {
	return registers[r.key]
}

type raw struct {
	value int
}

func (r raw) val(registers *[30]int) int {
	return r.value
}

type none struct {
}

func (r none) val(registers *[30]int) int {
	return 0
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	dic := make(map[stateID]state)
	var i, j int
	var from, to0, to1 stateID
	var value0, value1 bool
	var dir0, dir1 direction

	for scanner.Scan() {
		if i < 3 {
			i++
			continue
		}

		switch j % 10 {
		case 0:
			text, replace := scanner.Text(), "In state "
			from = stateID(strings.Replace(text, replace, "", -1)[0])
		case 2:
			text, replace := scanner.Text(), "    - Write the value "
			value0 = strings.Replace(text, replace, "", -1) == "1."
		case 3:
			text, replace := scanner.Text(), "    - Move one slot to the "
			dir0 = toDir(strings.Replace(text, replace, "", -1))
		case 4:
			text, replace := scanner.Text(), "    - Continue with state "
			to0 = stateID(strings.Replace(text, replace, "", -1)[0])
		case 6:
			text, replace := scanner.Text(), "    - Write the value "
			value1 = strings.Replace(text, replace, "", -1) == "1."
		case 7:
			text, replace := scanner.Text(), "    - Move one slot to the "
			dir1 = toDir(strings.Replace(text, replace, "", -1))
		case 8:
			text, replace := scanner.Text(), "    - Continue with state "
			to1 = stateID(strings.Replace(text, replace, "", -1)[0])

			dic[from] = state{
				action{value0, dir0, to0},
				action{value1, dir1, to1},
			}
		}

		j++
	}

	const length int = 123022090
	var tape [length]bool
	pointer := length / 2
	stid := stateID('A')

	for i := 0; i < 12302209; i++ {
		st := dic[stid]

		var a action
		if tape[pointer] {
			a = st.action1
		} else {
			a = st.action0
		}

		tape[pointer] = a.value
		pointer += a.dir.offset()
		stid = a.nextState
	}

	count := 0
	for _, bit := range tape {
		if bit {
			count++
		}
	}

	fmt.Println(count)
}

func toDir(str string) direction {
	if str == "right." {
		return right
	}

	return left
}

type state struct {
	action0 action
	action1 action
}

type action struct {
	value     bool
	dir       direction
	nextState stateID
}

type stateID rune

type direction int

const (
	left  direction = 0
	right direction = 1
)

func (d direction) offset() int {
	if d == left {
		return -1
	}

	return 1
}

func (d direction) String() string {
	if d == left {
		return "left"
	}

	return "right"
}

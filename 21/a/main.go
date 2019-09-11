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

	initial := []string{".#.", "..#", "###"}
	art := toMatrix(initial)

	var patterns []pattern
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, " => ")

		key := toMatrix(strings.Split(pair[0], "/"))
		value := toMatrix(strings.Split(pair[1], "/"))

		pat := pattern{key, value}
		patterns = append(patterns, pat)
	}

	dic := make(map[string]matrix)
	for _, p := range patterns {
		key := p.key

		for i := 0; i < 4; i++ {
			dic[key.toString()] = p.value
			key = key.rotate()
		}

		key = key.flip()

		for i := 0; i < 4; i++ {
			dic[key.toString()] = p.value
			key = key.rotate()
		}
	}

	for i := 0; i < 5; i++ {
		s := split(art)
		//fmt.Println("split", s)
		m := mapMats(s, dic)
		//fmt.Println("map", m)
		art = concat(m)
		//fmt.Println("concat", art)
	}

	count := 0
	for _, row := range art {
		for _, col := range row {
			if col == '#' {
				count++
			}
		}
	}

	fmt.Println(count)
}

func mapMats(mats [][]matrix, dic map[string]matrix) [][]matrix {
	var newMatrix [][]matrix

	for y, row := range mats {
		newMatrix = append(newMatrix, []matrix{})
		for _, col := range row {
			newMatrix[y] = append(newMatrix[y], dic[col.toString()])
		}
	}

	return newMatrix
}

func concat(mats [][]matrix) matrix {
	var mat matrix

	for y, row := range mats {
		for x, col := range row {
			height := len(col)

			for sy, subrow := range col {
				if x == 0 {
					mat = append(mat, []rune{})
				}

				mat[y*height+sy] = append(mat[y*height+sy], subrow...)
			}
		}
	}

	return mat
}

func split(mat matrix) [][]matrix {
	var newMatrix [][]matrix
	var size int

	if len(mat)%2 == 0 {
		size = 2
	} else {
		size = 3
	}

	for y := 0; y < len(mat)/size; y++ {
		newMatrix = append(newMatrix, []matrix{})
		for x := 0; x < len(mat)/size; x++ {
			sample := mat.sample(x, y, size)
			newMatrix[y] = append(newMatrix[y], sample)
		}
	}

	return newMatrix
}

type pattern struct {
	key   matrix
	value matrix
}

type matrix [][]rune

func toMatrix(strs []string) matrix {
	var mat matrix

	for _, str := range strs {
		mat = append(mat, []rune(str))
	}

	return mat
}

func (mat matrix) matMap(mapper func(int, int, int, int) (int, int)) matrix {
	var newMatrix matrix
	rows := len(mat) - 1
	cols := len(mat[0]) - 1

	for y, row := range mat {
		newMatrix = append(newMatrix, []rune{})
		for x := range row {
			nx, ny := mapper(x, y, rows, cols)
			newMatrix[y] = append(newMatrix[y], mat[nx][ny])
		}
	}

	return newMatrix
}

func (mat matrix) rotate() matrix {
	return mat.matMap(func(x, y, rows, cols int) (int, int) {
		return rows - x, y
	})
}

func (mat matrix) flip() matrix {
	return mat.matMap(func(x, y, rows, cols int) (int, int) {
		return y, cols - x
	})
}

func (mat matrix) sample(x, y, size int) matrix {
	if size == 2 {
		return matrix{
			{mat[y*2][x*2], mat[y*2][x*2+1]},
			{mat[y*2+1][x*2], mat[y*2+1][x*2+1]},
		}
	}

	return matrix{
		{mat[y*3][x*3], mat[y*3][x*3+1], mat[y*3][x*3+2]},
		{mat[y*3+1][x*3], mat[y*3+1][x*3+1], mat[y*3+1][x*3+2]},
		{mat[y*3+2][x*3], mat[y*3+2][x*3+1], mat[y*3+2][x*3+2]},
	}
}

func (mat matrix) toString() string {
	str := ""

	for _, row := range mat {
		for _, col := range row {
			str += string(col)
		}
		str += "/"
	}

	return str
}

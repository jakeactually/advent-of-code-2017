package main

import "fmt"

func main() {
	var arr [2018]int
	pos := 0

	for i := 1; i < 2018; i++ {
		pos = (pos+359)%i + 1

		for j := i; j > pos; j-- {
			arr[j] = arr[j-1]
		}

		arr[pos] = i
	}

	fmt.Println(arr[pos+1])
}

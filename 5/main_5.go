package main

import "fmt"

func checkSlices(a []int, b []int) (bool, []int) {
	m := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range a {
		m[v] = true
	}

	for _, v := range b {
		if m[v] {
			result = append(result, v)
		}
	}

	if len(result) > 0 {
		return true, result
	} else {
		return false, []int{}
	}
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(checkSlices(a, b))
}

package main

import "fmt"

func checkSlice(slice1 []string, slice2 []string) []string {
	m := make(map[string]bool)
	result := make([]string, 0)

	for _, v := range slice2 {
		m[v] = true
	}
	for _, v := range slice1 {
		if _, ok := m[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := checkSlice(slice1, slice2)
	fmt.Println(result)
}

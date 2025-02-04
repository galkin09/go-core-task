package main

import (
	"fmt"
	"math/rand"
)

func sliceExample(slice []int) []int {
	var newSlice []int
	for _, v := range slice {
		if v%2 == 0 {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func addElements(slice []int, num int) []int {
	return append(slice, num)
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	var originalSlice []int

	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, rand.Intn(10)+1)
	}

	fmt.Println("Оригинал", originalSlice)
	fmt.Println("Четные числа", sliceExample(originalSlice))
	fmt.Println("Добавляем число например 10", addElements(originalSlice, 10))
	fmt.Println("Копируем слайс", copySlice(originalSlice))
	fmt.Println("Удаляем по индексу 5", removeElement(originalSlice, 5))
}

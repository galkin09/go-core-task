package main

import (
	"testing"
)

func Test_randomNum(t *testing.T) {
	seed := int64(42)

	expectedNumbers := []int{5, 87, 68, 50, 23}

	ch := randomNum(seed, 5)

	i := 0
	for num := range ch {
		if i >= len(expectedNumbers) {
			t.Errorf("Получено больше чисел, чем ожидалось")
			break
		}
		if num != expectedNumbers[i] {
			t.Errorf("Ожидалось %d, получено %d", expectedNumbers[i], num)
		}
		i++
	}

	if i != 5 {
		t.Errorf("Ожидалось %d чисел, получено %d", 5, i)
	}
}

package main

import "testing"

// Тест для функции sliceExample
func TestSliceExample(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
		{[]int{10, 20, 30}, []int{10, 20, 30}},
		{[]int{1, 3, 5}, []int{}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := sliceExample(test.input)
		if !compareSlices(result, test.expected) {
			t.Errorf("Ожидалось: %v, получено: %v", test.expected, result)
		}
	}
}

// Тест для функции addElements
func TestAddElements(t *testing.T) {
	tests := []struct {
		inputSlice []int
		inputNum   int
		expected   []int
	}{
		{[]int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
		{[]int{}, 10, []int{10}},
		{[]int{5}, 6, []int{5, 6}},
	}

	for _, test := range tests {
		result := addElements(test.inputSlice, test.inputNum)
		if !compareSlices(result, test.expected) {
			t.Errorf("Ожидалось: %v, получено: %v", test.expected, result)
		}
	}
}

// Тест для функции copySlice
func TestCopySlice(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{1, 2, 3}},
		{[]int{}},
		{[]int{10, 20, 30}},
	}

	for _, test := range tests {
		copied := copySlice(test.input)
		if !compareSlices(copied, test.input) {
			t.Errorf("Ожидалось: %v, получено: %v", test.input, copied)
		}
		if len(test.input) > 0 && &copied[0] == &test.input[0] {
			t.Errorf("Копия и оригинал ссылаются на один и тот же массив")
		}
	}
}

// Тест для функции removeElement
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		inputSlice []int
		index      int
		expected   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{[]int{1, 2, 3}, 0, []int{2, 3}},
		{[]int{1}, 0, []int{}},
		{[]int{}, 0, []int{}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
	}

	for _, test := range tests {
		result := removeElement(test.inputSlice, test.index)
		if !compareSlices(result, test.expected) {
			t.Errorf("Ожидалось: %v, получено: %v", test.expected, result)
		}
	}
}

// Вспомогательная функция для сравнения двух слайсов
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

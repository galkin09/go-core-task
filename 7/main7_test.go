package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()
	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()
	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	merged := MergeChannels(ch1, ch2, ch3)

	var result []int
	for num := range merged {
		result = append(result, num)
	}

	expected := []int{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Errorf("Expected %d numbers, got %d", len(expected), len(result))
	}

	for i, num := range expected {
		if result[i] != num {
			t.Errorf("Expected %d at position %d, got %d", num, i, result[i])
		}
	}
}

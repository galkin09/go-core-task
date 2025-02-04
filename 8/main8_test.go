package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup(3)

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			t.Logf("Goroutine %d finished", id)
		}(i)
	}

	wg.Wait()

	if len(wg.semaphore) != 0 {
		t.Errorf("Expected semaphore to be empty, but got %d items", len(wg.semaphore))
	}
}

func TestCustomWaitGroup_Zero(t *testing.T) {
	wg := NewCustomWaitGroup(3)

	wg.Add(0)

	wg.Wait()

	if len(wg.semaphore) != 0 {
		t.Errorf("Expected semaphore to be empty, but got %d items", len(wg.semaphore))
	}
}

func TestCustomWaitGroup_Negative(t *testing.T) {
	wg := NewCustomWaitGroup(3)

	wg.Add(-1)

	wg.Wait()

	if len(wg.semaphore) != 0 {
		t.Errorf("Expected semaphore to be empty, but got %d items", len(wg.semaphore))
	}
}

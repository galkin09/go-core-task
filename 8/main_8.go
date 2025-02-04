package main

import "fmt"

type CustomWaitGroup struct {
	semaphore chan struct{}
}

func NewCustomWaitGroup(count int) *CustomWaitGroup {
	return &CustomWaitGroup{
		semaphore: make(chan struct{}, count),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	for i := 0; i < delta; i++ {
		wg.semaphore <- struct{}{}
	}
}

func (wg *CustomWaitGroup) Done() {
	<-wg.semaphore
}

func (wg *CustomWaitGroup) Wait() {
	for i := 0; i < cap(wg.semaphore); i++ {
		select {
		case <-wg.semaphore:
		default:
		}
	}
}

func main() {
	wg := NewCustomWaitGroup(5)

	go func() {
		defer wg.Done()
		fmt.Println("Task 1 started")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Task 2 started")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Task 3 started")
	}()

	fmt.Println("Waiting for all tasks to complete...")
	wg.Wait()
	fmt.Println("All tasks completed!")
}

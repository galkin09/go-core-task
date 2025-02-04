package main

import (
	"fmt"
	"time"
)

func Pipeline(inputChan <-chan uint8, outputChan chan<- float64) {
	for num := range inputChan {
		result := float64(num) * float64(num) * float64(num)
		outputChan <- result
	}
	close(outputChan)
}

func main() {
	inputChan := make(chan uint8)
	outputChan := make(chan float64)

	go Pipeline(inputChan, outputChan)

	go func() {
		defer close(inputChan)
		for i := uint8(1); i <= 5; i++ {
			inputChan <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for result := range outputChan {
		fmt.Printf("Result: %.2f\n", result)
	}
}

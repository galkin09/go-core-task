package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNum(seed int64, count int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		rng := rand.New(rand.NewSource(seed))

		for i := 0; i < count; i++ {
			out <- rng.Intn(100)
		}
	}()

	return out
}

func main() {
	seed := time.Now().UnixNano()
	ch := randomNum(seed, 5)

	for v := range ch {
		fmt.Println(v)
	}
}

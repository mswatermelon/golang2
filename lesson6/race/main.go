package main

import (
	"math/rand"
)

const goroutineCount = 75694849

func main() {
	modifyingData := make([]int, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go func() {
			index := rand.Intn(goroutineCount - 1)

			modifyingData[index] = goroutineCount
		}()
	}
}

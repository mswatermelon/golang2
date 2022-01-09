package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

var count = 555

func worker() error {
	time.Sleep(time.Second)
	if rand.Intn(2) == 1 {
		return errors.New("Error just for fun")
	}
	return nil
}

func main() {
	var eg = errgroup.Group{}

	for i := 0; i < count; i++ {
		eg.Go(worker)
	}

	fmt.Println("Main: Waiting for workers to finish")
	if err := eg.Wait(); err != nil {
		fmt.Println("Error happened:", err)
	}

	fmt.Println("Main: Completed")
}

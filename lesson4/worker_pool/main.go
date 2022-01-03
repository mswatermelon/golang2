package main

import (
    "fmt"
)

func worker(jobs <-chan int, results chan<- int) {
    for _ = range jobs {
        results <- 1
    }
}

func main() {

    jobs := make(chan int, 10)
    results := make(chan int, 1000)

    for w := 1; w <= 100; w++ {
        go worker(jobs, results)
    }

    for j := 1; j <= 1000; j++ {
        jobs <- j
    }
    close(jobs)

	finalResult := 0
    for a := 1; a <= 1000; a++ {
        finalResult += <-results
    }

	fmt.Println(finalResult)
}

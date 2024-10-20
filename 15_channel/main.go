package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for i := range 10 {
		jobs <- i + 1
	}

	close(jobs)

	for range 2 {
		go double(jobs, results)
	}

	for range 10 {
		result := <-results
		fmt.Println(result)
	}
}

func double(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(1 * time.Second)
		results <- j * 2
	}
}

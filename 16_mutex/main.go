package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int = 1000
)

func main() {
	donCh := make(chan bool, 3)

	go UpdateBalance(donCh, 100)
	go UpdateBalance(donCh, 200)
	go UpdateBalance(donCh, 150)

	<-donCh
	<-donCh
	<-donCh

	fmt.Println("Final balance: ", balance)
}

func UpdateBalance(doneCh chan<- bool, amount int) {
	mu.Lock()

	fmt.Println("Updating balance")
	time.Sleep(1 * time.Second)

	balance -= amount
	doneCh <- true

	mu.Unlock()
	fmt.Println("Balance updated")
}

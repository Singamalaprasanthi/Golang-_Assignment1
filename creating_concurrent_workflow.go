// 1. creating concurrent workflows in go with goroutines and channels

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	a := 10
	go worker(c, 1)
	go worker(c, 2)
	go worker(c, 3)
	go generateInts(c, a)

	time.Sleep(1 * time.Second)
}

func worker(c chan int, id int) {
	for i := range c {
		fmt.Printf("Worker %d received value %d\n", id, i)
		time.Sleep(1 * time.Second)
		fmt.Println("Executing Goroutines")
	}
}

func generateInts(c chan int, a int) {
	for i := 0; i < a; i++ {
		c <- i
		fmt.Printf("Sent value %d\n", i)
		time.Sleep(1 * time.Millisecond)
	}
	close(c)
}

$go run main.go

Sent value 0
Worker 1 received value 0
Sent value 1
Worker 2 received value 1
Sent value 2
Worker 3 received value 2
Sent value 3
Sent value 4
Sent value 5
Executing Goroutines




package main

import (
	"fmt"
	"time"
)

func kelipatan(ch chan int) {
	for i := 1; ; i++ {
		ch <- i * 3
		time.Sleep(3 * time.Second)
	}
}

func main() {
	ch := make(chan int, 1)
	go kelipatan(ch)
	for {
		select {
		case multiple := <-ch:
			fmt.Printf("%d ", multiple)
		}
	}
}

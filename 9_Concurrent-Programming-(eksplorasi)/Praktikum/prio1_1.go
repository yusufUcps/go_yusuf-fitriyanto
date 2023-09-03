package main

import (
	"fmt"
	"time"
)

func kelipatan(x int) {
	for i := 1; ; i++ {
		fmt.Printf("%d ", i*x)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	fmt.Print("Masukkan angka X: ")
	var x int
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	go kelipatan(x)
	select {}
}


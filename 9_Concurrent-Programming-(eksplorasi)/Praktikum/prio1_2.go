package main

import (
	"fmt"
	"time"
)
func kelipatan(stopChan chan struct{}, doneChan chan struct{}) {
	for i := 1; ; i++ {
		select {
		case <-stopChan:
			fmt.Println("Kelipatan selesai.")
			doneChan <- struct{}{}
			return
		default:
			fmt.Printf("%d ", i*3)
			time.Sleep(3 * time.Second)
		}
	}
}

func main() {
	var x int
	fmt.Print("Masukan waktu berjalan dalam detik : ")
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	stopChan := make(chan struct{})
	doneChan := make(chan struct{}) 
	go kelipatan(stopChan, doneChan)
	time.Sleep(time.Duration(x) * time.Second)
	close(stopChan)
	<-doneChan 
}

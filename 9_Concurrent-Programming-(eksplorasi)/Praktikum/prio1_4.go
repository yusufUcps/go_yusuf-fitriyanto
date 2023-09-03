package main
import (
    "fmt"
    "sync"
    "time"
)
func main() {
    
    var m sync.Mutex
    go func() {
        for i := 1; i <= 5; i++ {

            m.Lock()
            fmt.Println("kelipatan 3 = ",i * 3)
            m.Unlock()
            time.Sleep(3 * time.Second)
        }
    }()
	 
    go func() {
        for i := 1; i <= 10; i++ {
  
            m.Lock()
            fmt.Println("kelipatan 5 = ",i * 5)
            m.Unlock()
            time.Sleep(1 * time.Second)
        }
    }()
 
    time.Sleep(22 * time.Second)
}
package main

import (
	"fmt"
)

func pow(x, n int) int {

	switch {
	case n == 0:
		return 1
	case n%2 == 0:
		setengahpow := pow(x, n/2)
		return setengahpow * setengahpow
	default:
		setengahpow := pow(x, (n-1)/2)
		return x * setengahpow * setengahpow
	}

}


func main() {
   fmt.Println(pow(2, 3))  // 8
   fmt.Println(pow(5, 3))  // 125
   fmt.Println(pow(10, 2)) // 100
   fmt.Println(pow(2, 5))  // 32
   fmt.Println(pow(7, 3))  // 343
}

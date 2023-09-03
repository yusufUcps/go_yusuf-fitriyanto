package main

import (
	"fmt"
	"math"
)
func primeNumber(number int) bool {
	
	switch {
	case number <= 1:
		return false
	case number <= 3:
		return true
	case number%2 == 0 || number%3 == 0:
		return false
	default:
		for i := 5; i <= int(math.Sqrt(float64(number))); i += 6 {
			if number%i == 0 || number%(i+2) == 0 {
				return false
			}
		}
		return true
	}
	
}

func primeX(number int) int {
	var primes []int
	i := 0 
    for {
		if number < len(primes){
			break
		}
		if (primeNumber(i)){
        primes = append(primes,i)
		}
        i++ 
    }

    if number <= 0 || number > len(primes) {
        return -1
    }
    return primes[number-1]
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
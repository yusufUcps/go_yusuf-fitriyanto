package main

import "fmt"

func main() {
	var angka int = 1
	
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if angka%3 == 0 && angka%5 == 0 {
				fmt.Print("FizzBuzz ")
			
			}else if angka%3 == 0 {
				fmt.Print("Fizz ")
			}else if angka%5 == 0{
				fmt.Print("Buzz ")
			}else{
				fmt.Printf("%d ",angka)
			}
			angka++
		}	
		fmt.Println("")
	} 
}

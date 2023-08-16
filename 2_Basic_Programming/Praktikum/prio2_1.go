package main

import "fmt"

func main() {
	var angka int
	fmt.Println("\nProgram segitiga *")
	fmt.Print("Input : ")
	fmt.Scan(&angka)

	fmt.Println("Output : ")
	for i := 1; i <= angka; i++ {
		for j := 1; j <= angka-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k%2 == 0{
				fmt.Print(" ")
			}else{
				fmt.Print("*")
			}
		}

		fmt.Println("")
		
	} 
}

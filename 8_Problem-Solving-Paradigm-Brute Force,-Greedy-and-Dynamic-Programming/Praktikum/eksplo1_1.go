package main

import "fmt"

func ubahRoman(number int) string {
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	
	roman := ""
	
	i := 0
	for number > 0 {
		for number >= values[i] {
			roman += romans[i]
			number -= values[i]
		}
		i++
	}
	return roman
}

func main() {
	angka := []int{4 ,9, 23, 2021 }
	for i := 0; i < len(angka) ; i++{
		fmt.Println("Input : ",angka[i])
		angkaRomawi := ubahRoman(angka[i])
		fmt.Println("Angka Romawi:", angkaRomawi,"\n")
	}
}
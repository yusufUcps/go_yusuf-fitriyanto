package main

import "fmt"

func main() {
	var angka int
	var h string

	fmt.Println("\nProgram Ganjil Genap")
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		h = "genap"
	} else {
		h = "ganjil"
	}

	fmt.Println("\n------------------------------------------")
	fmt.Printf("Angka %d merupakan bilangan %s\n", angka, h)
}

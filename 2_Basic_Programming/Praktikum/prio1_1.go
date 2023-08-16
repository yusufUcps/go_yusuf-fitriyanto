package main

import "fmt"

func main() {
	var panjangA, panjangB, tinggi, luas float64

	fmt.Println("Masukkan panjang sisi yang sejajar")
	fmt.Print("Panjang sisi 1: ")
	fmt.Scan(&panjangA)
	fmt.Print("Panjang sisi 2: ")
	fmt.Scan(&panjangB)

	fmt.Print("\nMasukkan tinggi trapsesium: ")
	fmt.Scan(&tinggi)

	luas = (panjangA + panjangB) * tinggi / 2

	fmt.Println("\n------------------------------------------")
	fmt.Println("Luas trapesium = ",luas)

}
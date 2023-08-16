package main

import "fmt"

func main() {
	var angka int
	var h string

	fmt.Println("\nProgram garade nilai")
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&angka)

	if angka >= 80 && angka <= 100 {
		h = "A"
	} else if angka >= 65 && angka <= 79 {
		h = "B"
	} else if angka >= 50 && angka <= 64 {
		h = "c"
	} else if angka >= 35 && angka <= 49 {
		h = "D"
	} else if angka >= 0 && angka <= 34 {
		h = "E"
	} else {
		h = "Nilai Invalid"
	}

	fmt.Println("\n------------------------------------------")
	fmt.Println("Nilai : ", angka)
	fmt.Println("Grade : ", h)
}

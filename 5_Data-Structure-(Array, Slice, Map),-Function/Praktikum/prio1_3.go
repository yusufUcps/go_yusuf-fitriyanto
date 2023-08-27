package main

import (
	"fmt"
)
func munculSekali(angka string) []int {
	jumlah := make(map[int]int)

	for _, char := range angka {
		ank := int(char - '0')
		jumlah[ank]++
	}

	hasil := solo(jumlah)
	return hasil
}

func solo(jumlah map[int]int) []int {
	angkaSolo := make([]int, 0)
	for no, jum := range jumlah {
		if jum == 1 {
			angkaSolo = append(angkaSolo, no)
		}
	}

	return angkaSolo
}


func main() {
	// Test cases
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}

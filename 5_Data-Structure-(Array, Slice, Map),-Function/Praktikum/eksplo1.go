package main

import (
	"fmt"
	"math"
)
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	diagonal_kiri := 0
	diagonal_kanan := 0
	n := len(matrix)

	for i := 0; i < n; i++ {
		diagonal_kiri += matrix[i][i]
		diagonal_kanan += matrix[i][n-i-1]
	}
	selisih_mutlak := int(math.Abs(float64(diagonal_kiri - diagonal_kanan)))

	fmt.Println("Matriks:")
	for _, isi := range matrix {
		fmt.Println(isi)
	}

	fmt.Printf("Jumlah Diagonal Kiri ke Kanan: %d\n", diagonal_kiri)
	fmt.Printf("Jumlah Diagonal Kanan ke Kiri: %d\n", diagonal_kanan)
	fmt.Printf("Selisih Absolut: %d\n", selisih_mutlak)
}

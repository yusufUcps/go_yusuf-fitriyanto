package main

import (
	"fmt"
)
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func DecimalToBinary(n int) [][]int {
	if n == 0 {
		return [][]int{{0}}
	}

	steps := [][]int{}
	for i := 0; i <= n; i++ {
		biner := []int{}
		temp := i
		for temp > 0 {
			bit := temp % 2
			biner = append(biner, bit)
			temp = temp / 2
		}
		reverse(biner)
		if len(biner) == 0 {
			biner = append(biner, 0)
		}
		steps = append(steps, biner)
	}

	return steps
}

func main() {
	var decimalNumber int
	fmt.Print("Masukkan angka desimal: ")
	fmt.Scanln(&decimalNumber)

	binaryResult := DecimalToBinary(decimalNumber)
	fmt.Printf("Output: %v\n", binaryResult)
}

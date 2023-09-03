package main

import "fmt"

func MaxSequence(arr []int) int{
    totalSekarang := 0
    totalAkhir := 0
    for i := 0; i < len(arr); i++ {
        totalAkhir = totalAkhir + arr[i]
        if  totalSekarang < totalAkhir {
             totalSekarang = totalAkhir
        }
        if totalAkhir < 0 {
            totalAkhir = 0
        }
    }
    return  totalSekarang
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))  // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
  fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}
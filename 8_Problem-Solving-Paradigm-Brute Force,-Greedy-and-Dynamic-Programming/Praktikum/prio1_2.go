package main

import (
	"fmt"
)

func pohonPascal(numRows int) [][]int {
    segitigaPascal := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        segitigaPascal[i] = make([]int, i+1)
        segitigaPascal[i][0] = 1
        segitigaPascal[i][i] = 1
        for j := 1; j < i; j++ {
            segitigaPascal[i][j] = segitigaPascal[i-1][j-1] + segitigaPascal[i-1][j]
        }
    }
    return segitigaPascal
}


func main() {
    var numRows int
    fmt.Print("Masukkan jumlah baris segitiga Pascal: ")
    _, err := fmt.Scan(&numRows)
	
    if err != nil {
        fmt.Println("Input tidak valid. Harap masukkan angka.")
        return
    }

    segitigaPascal := pohonPascal(numRows)
    fmt.Println(segitigaPascal)
}

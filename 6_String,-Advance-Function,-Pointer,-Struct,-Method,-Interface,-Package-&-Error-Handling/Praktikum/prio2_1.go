package main

import "fmt"

func caesar(offset int, input string) string {
	hasil := ""
	for _, char := range input {
		perubahan := 'a' + (char-'a'+rune(offset))%26
		hasil += string(perubahan)
	}
	return hasil
}

func main() {
	fmt.Println(caesar(3, "abc")) // Output: def
	fmt.Println(caesar(2, "alta")) // Output: cnvc
	fmt.Println(caesar(10, "alterraacademy")) // Output: kvdobbkkmknowcxi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // Output: bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // Output: mnopqrstuvwxyzabcdefghijkl
}

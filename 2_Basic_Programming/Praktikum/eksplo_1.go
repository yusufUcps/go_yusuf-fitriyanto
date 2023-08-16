package main

import "fmt"
import "strings"
import "bufio"
import "os"

func main() {
	var huruf,kecil string
	var k int
	fmt.Println("\nProgram palidrome")
	fmt.Println("\nApakah palidrome?")
	fmt.Print("Masukan kata : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	huruf = scanner.Text()
	fmt.Println("captured : ",huruf)
	k=1
	kecil = strings.ToLower(strings.TrimSpace(huruf))
	for i := 0; i < len(kecil)/2; i++ {
		if kecil[i] != kecil[len(kecil)-1-i] {
			k = 0
			break
		}
	}

	if k == 1 {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan palindrome")
	}
	
}

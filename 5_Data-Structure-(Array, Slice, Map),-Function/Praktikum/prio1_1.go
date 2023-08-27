package main

import (
	"fmt"
	"strings"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	gabungan := append(arrayA, arrayB...)
	katabaru := make(map[string]bool)
	hasil := []string{}
	
	for _, kata := range gabungan {
		if !katabaru[kata] {
			katabaru[kata] = true
			hasil = append(hasil,kata)
		}
	}
	return hasil
}

func main() {
	// Test cases
	gabung := ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"})
	fmt.Printf("[%s]\n", strings.Join(gabung, ", "))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	gabung = ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"})
	fmt.Printf("[%s]\n", strings.Join(gabung, ", "))
	// ["sergei", "jin", "steve", "bryan"]
	gabung = ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"})
	fmt.Printf("[%s]\n", strings.Join(gabung, ", "))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	gabung = ArrayMerge([]string{}, []string{"devil jin", "sergei"})
	fmt.Printf("[%s]\n", strings.Join(gabung, ", "))
	// ["devil jin", "sergei"]
	gabung = ArrayMerge([]string{"hwoarang"}, []string{})
	fmt.Printf("[%s]\n", strings.Join(gabung, ", "))
	// ["hwoarang"]
	gabung = ArrayMerge([]string{}, []string{})
	fmt.Printf("[%s]\n", strings.Join(gabung, ", "))
	// []
}

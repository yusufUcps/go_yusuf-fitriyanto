package main
import (
	"fmt"
)
type pair struct {
	name  string
	count int
}
func MostAppearItem(items []string) []pair {
	hitung := make(map[string]int)
	var hasil []pair
	for _, item := range items {
		hitung[item]++
	}
	for nama,jumlah := range hitung {
		hasil = append(hasil, pair{nama, jumlah})
	}
	return hasil
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
}

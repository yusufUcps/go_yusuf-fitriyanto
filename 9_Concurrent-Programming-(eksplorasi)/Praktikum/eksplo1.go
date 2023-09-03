package main
import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Produk struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func main() {

	productChan := make(chan []Produk)
	go func() {
		resp, err := http.Get("https://fakestoreapi.com/products")
		if err != nil {
			fmt.Println("Error : ", err)
			return
		}
		defer resp.Body.Close()

		var daftarproduk []Produk
		err = json.NewDecoder(resp.Body).Decode(&daftarproduk)
		if err != nil {
			fmt.Print("Error : ", err)
			return
		}
		productChan <- daftarproduk
	}()

	daftarproduk := <-productChan
	fmt.Println("DaftarProduk data")
	for _, produk:= range daftarproduk {
		fmt.Printf("==\nTitle: %s\nPrice: %.2f\ncategory:  %s\n==\n", produk.Title, produk.Price, produk.Category)
	}
}
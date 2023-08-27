package main
import "fmt"
type Car struct {
    carType string
    fuelin  float64
}
func (c *Car) tempuh() float64 {
    permill := 1.5 
    jarak_tempuh := c.fuelin * permill
    return jarak_tempuh
}
func main() {
	var carType string
	var fuelin float64
    fmt.Print("Masukkan jenis mobil: ")
    fmt.Scanln(&carType)

    fmt.Print("Masukkan jumlah bahan bakar (dalam L): ")
    fmt.Scanln(&fuelin)
    car := Car{
        carType: carType,
        fuelin:  fuelin,
    }
    jarak_tempuh := car.tempuh()
    fmt.Printf("car type: %s, est. distance: %.2f\n", car.carType, jarak_tempuh)
}

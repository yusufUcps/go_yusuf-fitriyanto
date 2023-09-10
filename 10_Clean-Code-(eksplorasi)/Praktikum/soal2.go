package main

import "fmt"

//stuktur kendaran
type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

//stuktur mobil
type Mobil struct {
	Kendaraan
}

//fungsi berjalan lamban
func (m *Mobil) BerjalanLamban() {
	m.TambahKecepatan(10)
}

//fungsi berjalan cepat
func (m *Mobil) BerjalanCepat() {
	for i:= 1 ; i<4 ; i++{
		m.TambahKecepatan(10)
	}
}

// fungsi menambhakan kecepatan
func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.BerjalanCepat()

	mobilLamban := Mobil{}
	mobilLamban.BerjalanLamban()

	fmt.Println("Kecepatan Mobil Cepat:", mobilCepat.KecepatanPerJam)
	fmt.Println("Kecepatan Mobil Lamban:", mobilLamban.KecepatanPerJam)
}

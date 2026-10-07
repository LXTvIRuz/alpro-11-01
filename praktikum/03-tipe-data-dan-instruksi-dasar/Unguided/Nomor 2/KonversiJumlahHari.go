package main

import "fmt"

func main() {
	var hari int

	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&hari)

	tahun := hari / 360
	sisa := hari % 360
	bulan := sisa / 30
	sisa = sisa % 30
	minggu := sisa / 7
	sisa = sisa % 7

	fmt.Println("Tahun:", tahun)
	fmt.Println("Bulan:", bulan)
	fmt.Println("Minggu:", minggu)
	fmt.Println("Sisa hari:", sisa)
}

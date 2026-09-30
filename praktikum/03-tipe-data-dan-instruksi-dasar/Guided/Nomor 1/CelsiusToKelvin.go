package main

import "fmt"

func main() {
	var celcius float64

	// Membaca input suhu dalam Celcius dari pengguna
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scanln(&celcius)

	// Mengonversi Celcius ke Kelvin
	fmt.Println(celcius + 273)
}

package main

import "fmt"

func main() {
	var suhu float64

	// Membaca input dari pengguna
	fmt.Scan(&suhu)

	// Konversi suhu
	celcius := (suhu - 32) * 5 / 9
	kelvin := celcius + 273.15
	fahrenheit := celcius*9/5 + 32

	// Menampilkan hasil konversi
	fmt.Printf("Suhu dalam Celsius: %.2f\n", celcius)
	fmt.Printf("Suhu dalam Kelvin: %.2f\n", kelvin)
	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", fahrenheit)

}

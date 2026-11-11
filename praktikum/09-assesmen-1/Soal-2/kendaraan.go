package main

import "fmt"

func main() {
	var kendaraan string
	var Motor, Mobil, Truk int

	for {
		fmt.Scan(&kendaraan)

		if kendaraan == "M" {
			Motor++
		} else if kendaraan == "K" {
			Mobil++
		} else if kendaraan == "T" {
			Truk++
		} else {
			break
		}
	}
	fmt.Printf("Tipe M: %d\n", Motor)
	fmt.Printf("Tipe K: %d\n", Mobil)
	fmt.Printf("Tipe T: %d\n", Truk)
}

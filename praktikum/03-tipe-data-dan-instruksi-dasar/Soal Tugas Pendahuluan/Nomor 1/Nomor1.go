package main

import "fmt"

func main() {
	var y, x int

	// Membaca input dari pengguna
	fmt.Scan(&y)
	fmt.Scan(&x)

	// kalkulator sederhana

	fmt.Println("Hasil Pembagian:", y%x)
}

package main

import "fmt"

func main() {
	var r, luas, keliling float64

	// Membaca input dari pengguna
	fmt.Scan(&r)

	// Menghitung luas dan keliling lingkaran
	luas = 3.14 * r * r
	keliling = 2 * 3.14 * r

	// Menampilkan output
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)

}

package main

import "fmt"

func main() {
	var Nama string
	var SkorMatematika, SkorBahasaInggris int

	// Membaca input dari pengguna
	fmt.Scan(&Nama)
	fmt.Scan(&SkorMatematika)
	fmt.Scan(&SkorBahasaInggris)

	// Menghitung total & rata-rata (pembagian bilangan;

	total := SkorMatematika + SkorBahasaInggris
	ratarata := total / 2

	// Menampilkan output
	fmt.Println(Nama)
	fmt.Println(total)
	fmt.Println(ratarata)

}

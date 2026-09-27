package main

import "fmt"

func main() {
	var a, b int

	// Membaca input dari pengguna
	fmt.Scan(&a)
	fmt.Scan(&b)

	// kalkulator sederhana
	fmt.Println("Hasil Penjumlahan:", a+b)
	fmt.Println("Hasil Pengurangan:", a-b)
	fmt.Println("Hasil Perkalian:", a*b)
	fmt.Println("Hasil Pembagian:", a/b)
	fmt.Println("Hasil Modulus:", a%b)

}

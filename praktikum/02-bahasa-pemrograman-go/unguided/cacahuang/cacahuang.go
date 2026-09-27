package main

import "fmt"

func main() {
	var uang int

	// Membaca input nilai uang dari pengguna
	fmt.Scan(&uang)

	// Menentukan jumlah pecahan uang
	sepuluhribu := uang / 10000
	sisa := uang % 10000

	limaribu := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000
	sisa = sisa % 1000

	// Menampilkan hasil
	fmt.Println("Jumlah pecahan 10.000:", sepuluhribu)
	fmt.Println("Jumlah pecahan 5.000:", limaribu)
	fmt.Println("Jumlah pecahan 1.000:", seribu)
	fmt.Println("Sisa uang:", sisa)
}

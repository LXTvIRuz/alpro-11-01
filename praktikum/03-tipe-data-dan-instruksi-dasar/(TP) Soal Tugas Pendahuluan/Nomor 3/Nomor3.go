package main

import "fmt"

func main() {
	var mil float64

	// Membaca input dari pengguna
	fmt.Scan(&mil)

	// Mengonversi mil ke kilometer
	km := mil * 1.6
	fmt.Printf("%.1f\n", km)
}

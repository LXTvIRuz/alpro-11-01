package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Input harus berupa bilangan bulat positif")
		return
	}

	jumlahFaktor := 0

	fmt.Print("Faktor dari ", n, ": ")

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}

	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Ya, ", n, " adalah bilangan prima")
	} else {
		fmt.Println("Tidak, ", n, " bukan bilangan prima")
	}
}

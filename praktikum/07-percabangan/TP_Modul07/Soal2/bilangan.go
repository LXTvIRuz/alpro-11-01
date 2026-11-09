package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	if bilangan == 0 {
		fmt.Println("Nol")
	} else {
		if bilangan%2 == 0 {
			fmt.Print("Genap, ")
		} else {
			fmt.Print("Ganjil, ")
		}

		if bilangan > 0 {
			fmt.Println("Positif")
		} else {
			fmt.Println("Negatif")
		}
	}
}

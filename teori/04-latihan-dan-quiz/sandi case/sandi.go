package main

import "fmt"

func main() {
	var n, angka, total int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&angka)

		digitPertama := angka / 1000
		digitTerakhir := angka % 10

		total += digitPertama + digitTerakhir
	}

	fmt.Println(total)
}

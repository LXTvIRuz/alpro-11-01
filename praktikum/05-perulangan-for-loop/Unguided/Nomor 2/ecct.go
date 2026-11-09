package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	jumlah := 0
	for i := 0; i < m; i++ {
		total := 0
		for j := 0; j < 6; j++ {
			var nilai int
			fmt.Scan(&nilai)
			total += nilai
		}

		rata := float64(total) / 6
		if rata >= 3.33 {
			fmt.Println("Mendapat Sertifikat")
			jumlah++
		} else {
			fmt.Println("Tidak Mendapat Sertifikat")
		}
	}
	fmt.Println(jumlah)
}

package main

import "fmt"

func main() {
	total := 0

	for hari := 1; hari <= 5; hari++ {
		var x int
		for {
			fmt.Printf("Hari %d : ", hari)
			fmt.Scan(&x)
			if x >= 0 && x <= 200 {
				break
			}
		}
		total += x
	}
	fmt.Printf("Jumlah pengunjung : %d orang\n", total)
}

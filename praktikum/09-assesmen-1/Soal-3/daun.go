package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var lebar, terkecil int
	for i := 0; i < n; i++ {
		fmt.Scan(&lebar)

		if i == 0 {
			terkecil = lebar
		} else if lebar < terkecil {
			terkecil = lebar
		}
	}
	fmt.Println(terkecil)
}

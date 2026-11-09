package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	terendah1 := math.MaxInt
	terendah2 := math.MaxInt

	var namaTerendah1, namaTerendah2 string

	i := 0

	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < terendah1 {
			terendah2 = terendah1
			namaTerendah2 = namaTerendah1

			terendah1 = nilai
			namaTerendah1 = nama
		} else if nilai < terendah2 {
			terendah2 = nilai
			namaTerendah2 = nama
		}

		i++
	}

	fmt.Println(namaTerendah1)
	fmt.Println(namaTerendah2)
}

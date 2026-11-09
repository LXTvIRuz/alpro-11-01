package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	biaya := 10000000.0

	for i := 2; i <= n; i++ {
		biaya = biaya * 1.06
	}

	fmt.Printf("%.0f\n", biaya)
}

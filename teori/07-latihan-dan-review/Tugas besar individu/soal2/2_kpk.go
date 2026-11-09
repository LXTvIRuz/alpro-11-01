package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	max := a
	if b > max {
		max = b
	}

	kpk := max
	for kpk%a != 0 || kpk%b != 0 {
		kpk += max
	}
	fmt.Println(kpk)
}

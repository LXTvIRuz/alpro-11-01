package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	min := a
	if b < min {
		min = b
	}

	fpb := 1
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			fpb = i
		}
	}
	fmt.Println(fpb)
}

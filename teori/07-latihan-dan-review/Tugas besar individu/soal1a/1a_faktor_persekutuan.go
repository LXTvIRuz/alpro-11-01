package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	min := a
	if b < min {
		min = b
	}

	first := true
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(i)
			first = false
		}
	}
	fmt.Println()
}

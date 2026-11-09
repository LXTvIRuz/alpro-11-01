package main

import "fmt"

func main() {
	var s [10]string
	a, b := 0, 0

	for i := 0; i < 10; i++ {
		fmt.Scan(&s[i])
		if s[i] == "A" {
			a++
		} else if s[i] == "B" {
			b++
		}
	}

	if a != 5 || b != 5 {
		fmt.Println("tidak valid")
		return
	}

	ca, cb := 0, 0
	for i := 0; i < 10; i++ {
		if s[i] == "A" {
			ca++
		} else {
			cb++
		}
		if ca == 4 {
			fmt.Println("A")
			return
		}
		if cb == 4 {
			fmt.Println("B")
			return
		}
	}
}

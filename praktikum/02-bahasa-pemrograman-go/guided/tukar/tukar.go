package main

import "fmt"

func main() {
	var a, b int

	// Membaca input dari;
	fmt.Scan(&a)
	fmt.Scan(&b)

	// menukar a dan b
	a, b = b, a

	// Menampilkan output
	fmt.Println(a)
	fmt.Println(b)

}

package main

import "fmt"

func main() {
	var nilai float64
	fmt.Print("Massukan nilai:")
	fmt.Scan(&nilai)

	if nilai < 50 {
		fmt.Println("Grade E")
	} else if nilai < 70 {
		fmt.Println("Grade C")
	} else if nilai < 85 {
		fmt.Println("Grade B")
	} else if nilai <= 100 {
		fmt.Println("Grade A")
	} else {
		fmt.Println("Nilai tidak valid")
	}
}

package main

import "fmt"

func main() {
	var C, R float64

	fmt.Print("Masukkan suhu Celsius: ")
	fmt.Scan(&C)

	R = (4.0 / 5.0) * C

	fmt.Println("Suhu Reamur:", R)
}

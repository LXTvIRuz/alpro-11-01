package main

import "fmt"

func main() {
	var berat, tinggiCm float64
	fmt.Scan(&berat, &tinggiCm)

	tinggi := tinggiCm / 100
	bmi := berat / (tinggi * tinggi)

	fmt.Printf("%.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Kurus")
	} else if bmi < 25 {
		fmt.Println("Normal")
	} else if bmi < 30 {
		fmt.Println("Gemuk")
	} else {
		fmt.Println("Obesitas")
	}
}

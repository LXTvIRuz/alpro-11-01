package main

import "fmt"

func main() {
	lamp_status := "Merah"

	if lamp_status == "Merah" {
		fmt.Println("Berhenti")
	} else if lamp_status == "Kuning" {
		fmt.Println("Hati-hati")
	} else if lamp_status == "Hijau" {
		fmt.Println("Jalan")
	} else {
		fmt.Println("Perintah tidak ditemukan")
	}
}

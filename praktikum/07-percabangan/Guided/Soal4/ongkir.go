package main

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimal (30KG)")
		return
	}

	var tarifperkg int

	switch zona {
	case 1:
		tarifperkg = 10000
	case 2:
		tarifperkg = 15000
	case 3:
		tarifperkg = 25000
	default:
		fmt.Println("Zona tidak Valid")
		return
	}

	beratbulat := int(math.Ceil(berat))
	total := beratbulat * tarifperkg

	if layanan == "Express" {
		total += total / 2
	}

	fmt.Println(beratbulat)
	fmt.Println(total)
}

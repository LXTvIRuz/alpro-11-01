package main

import "fmt"

func kabisat(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0
}

func hariDalamBulan(bulan, tahun int) int {
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if kabisat(tahun) {
			return 29
		}
		return 28
	}
	return 0
}

func main() {
	var tgl, bln, thn int
	fmt.Scan(&tgl, &bln, &thn)

	if bln < 1 || bln > 12 || tgl < 1 || tgl > hariDalamBulan(bln, thn) {
		fmt.Println("Tidak valid")
		return
	}

	tgl++
	if tgl > hariDalamBulan(bln, thn) {
		tgl = 1
		bln++
		if bln > 12 {
			bln = 1
			thn++
		}
	}

	fmt.Println("Valid")
	fmt.Printf("%d/%d/%d\n", tgl, bln, thn)
}

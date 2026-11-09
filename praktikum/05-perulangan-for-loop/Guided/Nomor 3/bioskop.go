package main

import "fmt"

func main() {
	var t int
	fmt.Print("Input t (jumlah transaksi): ")
	fmt.Scan(&t)

	totalpendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisfilm, hari string
		var jumlahtiket int

		var hargasatuan int

		fmt.Printf("input transaksi ke-%d (jenisfilm, hari, jumlahtiket): ", i)
		fmt.Scan(&jenisfilm, &hari, &jumlahtiket)

		if hari == "senin" || hari == "selasa" || hari == "rabu" || hari == "kamis" {
			if jenisfilm == "reguler" {
				hargasatuan = 40000
			} else if jenisfilm == "3D" {
				hargasatuan = 60000

			} else if jenisfilm == "IMAX" {
				hargasatuan = 80000
			}
		} else if hari == "jumat" || hari == "sabtu" || hari == "minggu" {
			if jenisfilm == "reguler" {
				hargasatuan = 50000
			} else if jenisfilm == "3D" {
				hargasatuan = 70000
			} else if jenisfilm == "IMAX" {
				hargasatuan = 90000
			}
		}

		totalharga := hargasatuan * jumlahtiket

		if jumlahtiket >= 5 {
			diskon := (totalharga * 10) / 100
			totalharga = totalharga - diskon
		}

		fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalharga)

		totalpendapatan = totalpendapatan + totalharga
	}
	fmt.Printf("output total pendapatan: %d\n", totalpendapatan)
}

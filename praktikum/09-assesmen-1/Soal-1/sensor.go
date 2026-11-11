package main

import "fmt"

func main() {
	var kelembapan, sebelumnya, tertinggi, terendah, total, jumlah int

	for {
		fmt.Scan(&kelembapan)

		if jumlah > 0 && kelembapan == sebelumnya {
			break
		}

		if jumlah == 0 {
			tertinggi = kelembapan
			terendah = kelembapan
		} else {
			if kelembapan > tertinggi {
				tertinggi = kelembapan
			}

			if kelembapan < terendah {
				terendah = kelembapan
			}
		}
		total += kelembapan
		jumlah++
		sebelumnya = kelembapan
	}
	rataRata := float64(total) / float64(jumlah)
	fmt.Printf("%d, %d, %.2f\n", tertinggi, terendah, rataRata)
}

package main

import (
	"fmt"
)

func main() {
	const passwordBenar = "1234"
	saldo := 10000000
	kesempatan := 3
	isLogin := false

	for kesempatan > 0 && !isLogin {
		var inputPassword string
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&inputPassword)

		if inputPassword == passwordBenar {
			isLogin = true
		} else {
			kesempatan--
			if kesempatan > 0 {
				fmt.Printf("Password salah. Sisa kesempatan: %d\n", kesempatan)
			}
		}
	}

	if !isLogin {
		fmt.Println("ATM terblokir dan program berakhir.")
		return
	}

	pilihan := 0
	for pilihan != 3 {
		fmt.Println("\n--- Menu ATM ---")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu (1-3): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Printf("Saldo saat ini: Rp %d\n", saldo)
		case 2:
			var jumlahTarik int
			fmt.Print("Masukkan jumlah uang yang akan ditarik: ")
			fmt.Scan(&jumlahTarik)

			if jumlahTarik <= saldo {
				saldo -= jumlahTarik
				fmt.Printf("Penarikan berhasil. Sisa saldo Anda: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		case 3:
			fmt.Println("Keluar dari program. Terima kasih.")
		default:
			fmt.Println("Pilihan menu tidak valid.")
		}
	}
}

# <h1 align="center">Tugas Pendahuluan Modul [03] - [Algoritma Pemrograman - Ganjil 2025/2026 “VARIABEL DAN OPERATOR”]</h1>
<p align="center">[Deta Arief Syahputra] - [109092600019]</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main() {
	var y, x int

	// Membaca input dari pengguna
	fmt.Scan(&y)
	fmt.Scan(&x)

	// kalkulator sederhana

	fmt.Println("Hasil Pembagian:", y%x)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/TP%20Soal%20Tugas%20Pendahuluan/Nomor%201/Nomor1_output.png)



#### Deskripsi
Soal Nomor 1 memberi saya perintah untuk membuat program dalam bahasa Go untuk membagi kue sama rata yang dimasak sang ibu untuk sekeluarga. kode Golang di atas adalah solusi saya untuk masalah tersebut, dan sudah ada bukti outputnya.

### 2. Program Bahasa Go Bertipe Bool

```go
package main

import "fmt"

func main() {
	var benar bool = true
	var salah bool = false

	fmt.Println(benar)
	fmt.Println(salah)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/TP%20Soal%20Tugas%20Pendahuluan/Nomor%202/Nomor2_output.png)


#### Deskripsi
Soal Nomor 2 memberi saya perintah untuk membuat program dalam bahasa Go untuk membaca dan mencetak nilai bertipe bool. kode Golang di atas adalah solusi saya untuk masalah tersebut, dan sudah ada bukti outputnya.

### 3. Konversi Mil to Km

```go
package main

import "fmt"

func main() {
	var mil float64

	// Membaca input dari pengguna
	fmt.Scan(&mil)

	// Mengonversi mil ke kilometer
	km := mil * 1.6
	fmt.Printf("%.1f\n", km)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/TP%20Soal%20Tugas%20Pendahuluan/Nomor%203/Nomor3_output.png)


#### Deskripsi
Soal Nomor 3 memberi saya perintah untuk membuat program dalam bahasa Go untuk mengkonversi Jarak system metrik mil ke km. kode Golang di atas adalah solusi saya untuk masalah tersebut, dan sudah ada bukti outputnya.

## Kesimpulan
Berdasarkan tugas yang telah dilakukan, saya dapat memahami penggunaan variabel, operator, dan tipe data bool dalam bahasa Go.
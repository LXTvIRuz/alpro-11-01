# <h1 align="center">Laporan Praktikum Modul 03 - Variabel dan Operator</h1>
<p align="center">Deta Arief Syahputra - 109092600019</p>

## Dasar Teori

### A. Variabel dan Tipe Data pada Bahasa Go
Variabel merupakan wadah dalam memori komputer yang digunakan untuk menyimpan data yang nilainya dapat diubah selama eksekusi program. Pada bahasa pemrograman Go (Golang), variabel bersifat statically typed, yang artinya setiap variabel memiliki tipe data yang jelas dan tetap. Deklarasi variabel pada Go dapat dilakukan secara eksplisit menggunakan kata kunci var maupun secara implisit dengan operator pendek :=.

Beberapa tipe data dasar yang sering digunakan antara lain:
1. **Integer (int)**: Digunakan untuk menyimpan bilangan bulat tanpa koma desimal (contoh: -5, 0, 100).
2. **Floating Point (float64)**: Digunakan untuk menyimpan bilangan real atau desimal (contoh: 3.14, 273.15).
3. **String (string)**: Digunakan untuk menyimpan sekumpulan karakter atau teks.

### B. Operator Aritmatika dan Modulo
Operator dalam bahasa pemrograman Go digunakan untuk melakukan operasi matematika dasar maupun logika pada variabel atau nilai.
1. **Operator Aritmatika Dasar**:
   - Penjumlahan (+): Menjumlahkan dua nilai.
   - Pengurangan (-): Mengurangi nilai sebelah kiri dengan nilai sebelah kanan.
   - Perkalian (*): Mengalikan dua nilai.
   - Pembagian (/): Membagi nilai sebelah kiri dengan nilai sebelah kanan. Jika kedua operan adalah integer, pembagian akan menghasilkan pembulatan ke bawah (pembagian bulat).
2. **Operator Modulo (%)**: Digunakan untuk mendapatkan sisa hasil bagi dari dua bilangan bulat.

---

## Guided

### 1. Konversi Suhu Celsius ke Kelvin

```go
package main

import "fmt"

func main() {
	var celcius float64

	// Membaca input suhu dalam Celcius dari pengguna
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scanln(&celcius)

	// Mengonversi Celcius ke Kelvin
	fmt.Println(celcius + 273)
}

```

#### Deskripsi
Program ini bertujuan untuk mengonversi suhu dari derajat Celsius menjadi derajat Kelvin. Input yang diterima oleh program berupa bilangan real (float64) yang menyatakan suhu dalam Celsius. Hasil konversi dihitung menggunakan rumus K = C + 273 dan ditampilkannya sebagai keluaran dalam satuan Kelvin.

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/Guided/Nomor%201/CelsiusToKelvin_output.png)

---

### 2. Pertukaran Nilai Tiga Variabel (x, y, z)

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}

```

#### Deskripsi
Program ini digunakan untuk mempertukarkan nilai dari tiga buah bilangan bulat x, y, dan z. Ketentuan pertukaran nilai ditentukan sebagai berikut:
- Nilai y baru diisi dengan nilai x awal.
- Nilai x baru diisi dengan nilai z awal.
- Nilai z baru diisi dengan nilai y awal.

Input berupa 3 buah bilangan bulat x, y, dan z, kemudian keluaran berupa nilai x, y, dan z setelah mengalami proses pertukaran.

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/Guided/Nomor%202/PertukaranNilaibulat_output.png)

---

### 3. Penghitung Lembar Uang Kembalian (AlgoMart)

```go
package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&x)

	var sepuluhribu int = x / 10000
	var sisa int = x % 10000

	var limaribu int = sisa / 5000
	sisa = sisa % 5000

	var seribu int = sisa / 1000

	fmt.Println(sepuluhribu, limaribu, seribu)
}

```

#### Deskripsi
Program ini dibuat untuk membantu kasir AlgoMart menghitung pecahan uang kembalian berupa lembaran sepuluh ribuan (10.000), lima ribuan (5.000), dan seribuan (1.000). Masukan program berupa satu bilangan bulat x yang menyatakan jumlah total uang kembalian. Output yang dihasilkan berupa tiga bilangan bulat dipisahkan spasi yang masing-masing merepresentasikan jumlah lembar uang Rp10.000, Rp5.000, dan Rp1.000.

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/Guided/Nomor%203/KasirAlgomart_output.png)

---

## Unguided

### 1. Konversi Suhu Celsius ke Reamur

```go
package main

import "fmt"

func main() {
	var C, R float64

	fmt.Print("Masukkan suhu Celsius: ")
	fmt.Scan(&C)

	R = (4.0 / 5.0) * C

	fmt.Println("Suhu Reamur:", R)
}

```

#### Deskripsi
Program ini berfungsi untuk mengonversi suhu dari derajat Celsius menjadi derajat Reamur. Masukan berupa bilangan real (float64) yang menyatakan suhu Celsius. Perhitungan dilakukan menggunakan rumus $R = \frac{4}{5} \times C$. Output yang dikeluarkan merupakan suhu dalam derajat Reamur berbentuk bilangan real.

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/Unguided/Nomor%201/output_nomor1.png)

---

### 2. Konversi Hari ke Tahun, Bulan, Minggu, dan Sisa Hari

```go
package main

import "fmt"

func main() {
	var hari int

	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&hari)

	tahun := hari / 360
	sisa := hari % 360
	bulan := sisa / 30
	sisa = sisa % 30
	minggu := sisa / 7
	sisa = sisa % 7

	fmt.Println("Tahun:", tahun)
	fmt.Println("Bulan:", bulan)
	fmt.Println("Minggu:", minggu)
	fmt.Println("Sisa hari:", sisa)
}

```

##### Output
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/Unguided/Nomor%202/output_nomor2.png)

#### Deskripsi
Program ini mengonversi total jumlah hari ke dalam satuan tahun, bulan, minggu, dan sisa hari dengan pendekatan aturan konversi:
- 1 minggu = 7 hari
- 1 bulan = 30 hari
- 1 tahun = 12 bulan (360 hari)

Masukan berupa bilangan bulat yang menyatakan total jumlah hari. Hasil pengolahan menampilkan 4 bilangan bulat secara berurutan yang merepresentasikan jumlah tahun, bulan, minggu, dan sisa hari.

---

## Kesimpulan
Berdasarkan praktikum Modul 03 mengenai Variabel dan Operator pada bahasa Go, dapat disimpulkan bahwa:
1. Pemilihan tipe data yang tepat (seperti int untuk bilangan bulat dan float64 untuk desimal) sangat penting agar perhitungan matematis menghasilkan nilai yang akurat.
2. Operator aritmatika (+, -, *, /) serta operator modulo (%) dan pembagian bulat memungkinkan pengolahan data numerik yang kompleks, seperti konversi waktu dan pecahan uang kembalian.
3. Pertukaran nilai variabel dapat dilakukan secara efisien dengan memahami alur penyimpanan nilai sementara (*temporary variable*) atau fitur multiple assignment pada bahasa Go.

---

## Referensi
1. The Go Authors. (2026). Go Documentation. Diakses melalui https://go.dev/doc/
2. Donovan, A. A. A., & Kernighan, B. W. (2015). *The Go Programming Language*. Boston: Addison-Wesley Professional. Diakses melalui https://www.gopl.io/
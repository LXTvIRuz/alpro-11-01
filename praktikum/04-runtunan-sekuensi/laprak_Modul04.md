# <h1 align="center">Laporan Praktikum Modul 04 - Runtunan/Sekuensi</h1>
<p align="center">[Nama Praktikan] - [NIM]</p>

## Dasar Teori

### A. Percabangan dengan if-else
Percabangan adalah mekanisme yang membuat program memilih alur eksekusi berdasarkan suatu kondisi. Menurut The Go Authors (n.d.), pernyataan `if` pada Go menentukan eksekusi bersyarat untuk dua cabang sesuai nilai boolean dari kondisinya: cabang `if` dijalankan jika kondisi bernilai `true`, dan cabang `else` (jika ada) dijalankan jika `false`. Beberapa kondisi dapat disusun berantai dengan `else if`, dan hanya cabang pertama yang kondisinya terpenuhi yang dijalankan. Kondisi pada Go tidak perlu dibungkus tanda kurung, tetapi kurung kurawal `{ }` wajib ada, dan kata kunci `else` harus berada di baris yang sama dengan kurung kurawal penutup cabang sebelumnya.

### B. Operator Relasional dan Operator Logika

#### 1. Operator Relasional
Operator relasional membandingkan dua nilai dan menghasilkan nilai boolean. Go menyediakan enam operator: `==` (sama dengan), `!=` (tidak sama dengan), `<`, `<=`, `>`, dan `>=`. Hasil perbandingan inilah yang dipakai sebagai kondisi pada `if` maupun `switch`.

#### 2. Operator Logika dan Prioritas Operator
Operator logika menggabungkan atau membalik nilai boolean: `&&` (AND, true jika kedua sisi true), `||` (OR, true jika salah satu sisi true), dan `!` (NOT, membalik nilai). The Go Authors (n.d.) menjelaskan bahwa operand sebelah kanan `&&` dan `||` hanya dievaluasi jika diperlukan, yang dikenal sebagai *short-circuit evaluation*. Misalnya pada `a || b`, jika `a` sudah true maka `b` tidak dievaluasi.

Urutan prioritas dari yang tertinggi: `!`, operator aritmatika perkalian (`*`, `/`, `%`), operator aritmatika penjumlahan (`+`, `-`), operator relasional, `&&`, lalu `||`. Akibatnya, ekspresi `a > 0 && b > 0 || c > 0` dibaca sebagai `(a > 0 && b > 0) || c > 0`. Tanda kurung dapat dipakai untuk memperjelas atau mengubah urutan evaluasi.

### C. Pernyataan switch

#### 1. switch dengan Ekspresi
Pernyataan `switch` menyediakan percabangan banyak arah. Nilai ekspresi pada `switch` dibandingkan dengan setiap `case` dari atas ke bawah, dan `case` pertama yang cocok dijalankan (The Go Authors, n.d.). Berbeda dengan bahasa C atau Java, Go tidak melanjutkan ke `case` berikutnya secara otomatis sehingga tidak diperlukan `break`. Satu `case` dapat memuat beberapa nilai yang dipisahkan koma, dan `default` dijalankan jika tidak ada `case` yang cocok.

#### 2. switch Tanpa Ekspresi
Jika ekspresi pada `switch` dihilangkan, Go menganggapnya sama dengan `switch true`, sehingga setiap `case` berisi kondisi boolean dan `case` pertama yang bernilai true dijalankan. Bentuk ini cocok untuk menyeleksi rentang nilai, misalnya pengelompokan nilai huruf atau klasifikasi usia dan penghasilan, dan sering dipakai sebagai pengganti rantai `if-else if` yang panjang (The Go Authors, n.d.).

### D. Membaca Masukan dari Pengguna
Fungsi `fmt.Scan` membaca nilai dari masukan standar dan memisahkan nilai-nilainya berdasarkan spasi, sehingga masukan seperti nama yang mengandung spasi akan terpotong (The Go Authors, n.d.). Untuk membaca satu baris penuh, dapat dipakai `bufio.Reader` dengan metode `ReadString('\n')`, kemudian karakter baris baru di ujungnya dibuang dengan `strings.TrimSpace`. Nilai berikutnya dapat dibaca dari `bufio.Reader` yang sama menggunakan `fmt.Fscan` agar tidak ada masukan yang hilang akibat buffering.

## Guided

### 1. grade.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()
	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else if nilai >= 0 && nilai < 60 {
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
```
#### Deskripsi
Program menerima nama dan nilai seorang siswa, lalu menentukan nilai huruf dengan rantai `if - else if - else`. Nama dibaca satu baris penuh menggunakan `bufio.Reader` agar nama yang memuat spasi, seperti "Hong Gil-dong", tidak terpotong, sedangkan nilai dibaca dengan `fmt.Fscan` ke variabel `float64` supaya nilai desimal tetap diperlakukan benar.

Kondisi diperiksa dari rentang tertinggi: `>= 90` untuk A, `>= 80` untuk B, `>= 70` untuk C, `>= 60` untuk D, dan sisanya F pada blok `else`. Karena rantai `else if` hanya menjalankan cabang pertama yang terpenuhi, batas atas tiap rentang tidak perlu ditulis lagi. Untuk masukan `Hong Gil-dong` dan `95`, program mencetak `Hong Gil-dong mendapatkan nilai A`.

##### Output
![Screenshot Output Guided](/praktikum/04-runtunan-sekuensi/guided/Nomor%201/output_nomor1.png)

### 2. penilaian.go

```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
```
#### Deskripsi
Program menampilkan menu, membaca pilihan pengguna, lalu mengambil tindakan sesuai pilihan tersebut. Jika pilihan `1`, program meminta nama dan nilai siswa kemudian menampilkan nilai huruf dengan logika yang sama seperti `grade.go`. Jika pilihan `0`, program mencetak `Keluar dari program.` dan berhenti. Untuk pilihan lain, program mencetak `Pilihan tidak valid.` dan berhenti.

Pilihan dibaca sebagai teks (string) lalu dibandingkan dengan `"1"` dan `"0"`. Dengan cara ini masukan yang bukan angka, misalnya huruf, tetap masuk ke cabang `else` dan dianggap tidak valid, bukan salah terbaca sebagai `0`. Hasil pengujian:

| Masukan | Keluaran |
|---|---|
| `1`, `Hong Gil-dong`, `95` | `Hong Gil-dong mendapatkan nilai A` |
| `0` | `Keluar dari program.` |
| `5` | `Pilihan tidak valid.` |

##### Output
![Screenshot Output Guided](/praktikum/04-runtunan-sekuensi/guided/Nomor%202/output_nomor2.png)

### 3. klasifikasi.go

```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}
```
#### Deskripsi
Program mengklasifikasikan seseorang berdasarkan usia dan gaji tahunan (dalam juta) menggunakan `switch` tanpa ekspresi, sehingga setiap `case` berisi kondisi boolean. Usia dan gaji dibaca sebagai bilangan bulat dengan `fmt.Scan`.

Urutan `case` dimanfaatkan agar kondisi tidak perlu ditulis berulang. Karena `case` dicek dari atas ke bawah, `case usia <= 25` yang muncul setelah `usia < 18` otomatis berarti usia 18 sampai 25. Begitu pula `usia <= 40` setelah `usia <= 25` berarti 26 sampai 40, dan `default` menangani usia di atas 40 dengan gaji di bawah 150. Hasil pengujian:

| Masukan | Keluaran |
|---|---|
| `22`, `60` | `Muda sukses` |
| `45`, `130` | `Perlu evaluasi finansial` |

##### Output
![Screenshot Output Guided](/praktikum/04-runtunan-sekuensi/guided/Nomor%203/output_nomor3.png)

## Unguided

### 1. grade

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	var nilai float64
	fmt.Fscan(reader, &nilai)

	var huruf string
	switch {
	case nilai >= 90:
		huruf = "A"
	case nilai >= 80:
		huruf = "B"
	case nilai >= 70:
		huruf = "C"
	case nilai >= 60:
		huruf = "D"
	default:
		huruf = "F"
	}

	fmt.Println(nama, "mendapatkan nilai", huruf)
}
```

##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/Nomor%201/output_nomor1.png)

#### Deskripsi
Program ini adalah modifikasi `grade.go` pada soal guided dengan mengganti rantai `if - else if - else` menjadi `switch` tanpa ekspresi. Setiap `case` berisi kondisi rentang nilai yang diperiksa dari yang tertinggi, dan `default` menangani nilai di bawah 60 sebagai F. Cara pembacaan nama dan nilai tidak berubah. Untuk masukan `Lee Haechan` dan `80`, program mencetak `Lee Haechan mendapatkan nilai B`, sama dengan contoh pada soal. Hasilnya tidak berbeda dengan versi `if-else`, tetapi penulisannya lebih rapi karena setiap rentang berada pada `case` yang sejajar.

### 2. pajak

```go
package main

import "fmt"

func main() {
	var penghasilan float64
	fmt.Scan(&penghasilan)

	var pajak float64

	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
	} else if penghasilan <= 100 {
		pajak = 0.05*50 + 0.10*(penghasilan-50)
	} else if penghasilan <= 200 {
		pajak = 0.05*50 + 0.10*50 + 0.15*(penghasilan-100)
	} else {
		pajak = 0.05*50 + 0.10*50 + 0.15*100 + 0.20*(penghasilan-200)
	}

	fmt.Println(pajak)
}
```

##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/Nomor%202/output_nomor2.png)

#### Deskripsi
Program menghitung pajak penghasilan (dalam juta) dengan sistem bracket bertingkat menggunakan `if - else if - else` dan tipe data `float64` untuk penghasilan maupun pajak. Setiap bracket hanya mengenakan tarifnya pada bagian penghasilan yang berada di dalam rentang bracket tersebut:
- Penghasilan sampai 50 juta dikenai 5% dari seluruh penghasilan.
- Di atas 50 sampai 100 juta: 5% dari 50 juta pertama, ditambah 10% dari sisa di atas 50 juta.
- Di atas 100 sampai 200 juta: 5% dari 50 juta pertama, 10% dari 50 juta berikutnya, ditambah 15% dari sisa di atas 100 juta.
- Di atas 200 juta: ditambah 15% dari 100 juta berikutnya, dan 20% dari sisa di atas 200 juta.

Karena rantai `else if` dicek dari atas, batas bawah tiap bracket tidak perlu ditulis lagi. Hasil pengujian:

| Masukan | Perhitungan | Keluaran |
|---|---|---|
| `75` | 2.5 + 0.1 × 25 | `5` |
| `150` | 2.5 + 5 + 0.15 × 50 | `15` |
| `250` | 2.5 + 5 + 15 + 0.2 × 50 | `32.5` |

## Kesimpulan
Pada praktikum ini diterapkan percabangan dalam Go menggunakan `if - else if - else` dan `switch`. Rantai `if-else` cocok untuk kondisi bertingkat seperti penentuan nilai huruf dan bracket pajak, karena hanya cabang pertama yang terpenuhi yang dijalankan sehingga batas bawah tiap rentang tidak perlu ditulis ulang. Pernyataan `switch`, terutama bentuk tanpa ekspresi, menghasilkan kode yang lebih ringkas dan mudah dibaca untuk menyeleksi rentang nilai maupun gabungan beberapa kondisi, seperti pada klasifikasi usia dan gaji. Selain itu, praktikum ini menunjukkan pentingnya urutan pengecekan kondisi, prioritas operator relasional dan logika, serta cara membaca masukan berspasi dengan `bufio.Reader`. Seluruh program memberikan keluaran yang sesuai dengan contoh pada soal.

## Referensi
1. The Go Authors. (n.d.). *The Go Programming Language Specification*. https://go.dev/ref/spec
2. The Go Authors. (n.d.). *A Tour of Go*. https://go.dev/tour/
3. The Go Authors. (n.d). Docs https://go.dev/doc/
4. Google.com
# <h1 align="center">Tugas Pendahuluan Modul 05 - Perulangan For Loop</h1>
<p align="center">Deta Arief Syahputra - 109092600019</p>

### 1. Fungsi Perulangan `for` dan Tiga Komponennya

#### Deskripsi
Fungsi utama perulangan `for` adalah **menjalankan blok kode berulang kali** selama kondisi tertentu terpenuhi, sehingga instruksi yang sama tidak perlu ditulis berkali-kali. Di Go, `for` adalah satu-satunya struktur perulangan.

Bentuk umumnya:

```go
for inisialisasi; kondisi; post {
	...
}
```

Tiga komponennya:
- **Inisialisasi**: dijalankan satu kali di awal untuk menyiapkan variabel pengontrol, misalnya `i := 1`.
- **Kondisi**: dicek sebelum setiap iterasi. Selama bernilai `true`, badan perulangan dijalankan. Jika `false`, perulangan berhenti, misalnya `i <= 5`.
- **Post statement**: dijalankan di akhir setiap iterasi, biasanya untuk mengubah variabel pengontrol, misalnya `i++`.

### 2. Perbedaan `for` Tiga Komponen dan `for` Kondisi Saja

Go tidak memiliki perulangan `while` maupun `do-while`. Fungsinya digantikan oleh bentuk `for` yang berbeda.

#### `for` dengan tiga komponen
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
```

#### `for` hanya dengan kondisi
```go
package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
```

#### Deskripsi
Pada **`for` tiga komponen**, inisialisasi, kondisi, dan perubahan nilai ditulis sekaligus di header. Bentuk ini cocok ketika jumlah pengulangan sudah jelas.

Pada **`for` kondisi saja**, header hanya berisi kondisi (berfungsi seperti `while` di bahasa lain). Inisialisasi ditulis sebelum `for`, dan perubahan nilai ditulis di dalam badan perulangan. Bentuk ini cocok ketika jumlah pengulangan belum pasti. Variabel pengontrol harus diubah di dalam badan agar tidak terjadi perulangan tak hingga.

Kedua contoh di atas menghasilkan keluaran yang sama, yaitu angka 1 sampai 5.

### 3. Bilangan Ganjil dan Jumlahnya 

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlah := 0
	for i := 1; i <= n; i += 2 {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(i)
		jumlah += i
	}
	fmt.Println()
	fmt.Println(jumlah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder soal3/output1.png, output2.png, dan output3.png -->
![Screenshot Output Soal 3 - n = 10](/praktikum/05-perulangan-for-loop/TP_Modul05/Nomor%203/Output_Nomor3.png)

#### Deskripsi
Program membaca sebuah bilangan bulat positif `n` dengan `fmt.Scan`. Perulangan `for` dimulai dari `i := 1` dan bertambah 2 setiap iterasi (`i += 2`), sehingga hanya bilangan ganjil (1, 3, 5, ...) yang diproses selama `i <= n`.

Di setiap iterasi, nilai `i` dicetak lalu ditambahkan ke variabel `jumlah`. Spasi dicetak sebelum angka kedua dan seterusnya (`if i > 1`) agar tidak ada spasi berlebih di akhir baris. Setelah perulangan selesai, program pindah baris dan mencetak jumlah.

Hasil pengujian:

| Masukan | Keluaran |
|---|---|
| `10` | `1 3 5 7 9` dan `25` |
| `7` | `1 3 5 7` dan `16` |
| `1` | `1` dan `1` |

### 4. Tabel Perkalian 

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder soal4/output1.png dan output2.png -->
![Screenshot Output Soal 4 - a = 7, b = 5](/praktikum/05-perulangan-for-loop/TP_Modul05/Nomor%204/Output_Nomor4.png)

#### Deskripsi
Program membaca dua bilangan bulat positif, yaitu `a` (bilangan yang dikalikan) dan `b` (batas perkalian). Perulangan `for` berjalan dari `i := 1` sampai `i <= b`, dan setiap iterasi mencetak satu baris dengan format `a x i = hasil` memakai `fmt.Printf`. Hasil perkalian dihitung langsung lewat `a*i`, sehingga total baris yang dicetak sebanyak `b`.

Hasil pengujian:

| Masukan | Keluaran |
|---|---|
| `7` `5` | `7 x 1 = 7` sampai `7 x 5 = 35` (5 baris) |
| `3` `4` | `3 x 1 = 3` sampai `3 x 4 = 12` (4 baris) |

## Kesimpulan
Pada modul ini dipelajari cara kerja perulangan `for` dalam Go. Perulangan `for` terdiri dari tiga komponen, yaitu inisialisasi, kondisi, dan post statement, dan dapat juga ditulis hanya dengan kondisi untuk menggantikan `while` yang tidak tersedia di Go. Soal 3 menunjukkan bahwa langkah perulangan dapat diatur (`i += 2`) untuk memproses hanya bilangan ganjil sambil mengakumulasi jumlahnya. Soal 4 menunjukkan bahwa variabel pengontrol perulangan dapat langsung dipakai dalam perhitungan, sehingga tabel perkalian dapat dibuat dengan kode yang ringkas dibanding menulis tiap baris satu per satu.

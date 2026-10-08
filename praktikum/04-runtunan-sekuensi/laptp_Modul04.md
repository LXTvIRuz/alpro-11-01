# <h1 align="center">Tugas Pendahuluan Modul 04 - Runtunan/Sekuensi</h1>
<p align="center">[Nama Praktikan] - [NIM]</p>

### 1. Evaluasi Ekspresi Kontrol dalam Go

```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	if intOther+2*intNum != 30 || !(sngNum > 0) {
		fmt.Println("Beep")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder soal1/output.png -->
![Screenshot Output Soal 1](/praktikum/04-runtunan-sekuensi/TP%20Modul%204/Nomor%201/output_nomor1.png)

#### Deskripsi
Kondisi yang dipilih sebagai contoh adalah nomor 9: `intOther + 2 * intNum != 30 || !(sngNum > 0)`.

Proses evaluasinya:
1. `2 * intNum` dihitung lebih dulu karena `*` berprioritas lebih tinggi dari `+`, hasilnya `2 * 5 = 10`.
2. `intOther + 10 = 10 + 10 = 20`.
3. `20 != 30` bernilai **true**.
4. Karena sisi kiri operator `||` sudah true, Go langsung menyimpulkan seluruh ekspresi true tanpa mengevaluasi sisi kanan (*short-circuit*). Seandainya dievaluasi, `!(sngNum > 0)` juga true karena `-3 > 0` false lalu dibalik oleh `!`.
5. Kondisi `if` terpenuhi sehingga program mencetak `Beep`.

Jawaban true/false untuk seluruh kondisi pada soal:

| No | true/false | Kondisi |
|----|-----------|---------|
| 1 | false | `intNum > 5` |
| 2 | true | `intNum >= 5 && intOther < 11` |
| 3 | true | `sngNum != -1 \|\| intOther < 0` |
| 4 | true | `!(intNum > 3) \|\| intNum <= 5` |
| 5 | false | `!(intOther >= intNum)` |
| 6 | true | `0 - sngNum > 0` |
| 7 | true | `4 / 2 == intOther / intNum` |
| 8 | true | `intOther % 2 == 0` |
| 9 | true | `intOther + 2 * intNum != 30 \|\| !(sngNum > 0)` |
| 10 | true | `intOther > 0 && intNum > 0 \|\| sngNum > 0` |
| 11 | true | `sngNum > 0 \|\| (intNum >= 0 && -1 * intOther == -10)` |
| 12 | true | `intNum == 5` |
| 13 | true | `intNum > 0 \|\| (sngNum <= 0 && intOther == 13)` |
| 14 | true | `!(!(!(!(intNum > 0))))` |

### 2. Tracing: Evaluasi Pernyataan Kondisi

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder soal2/output.png -->
![Screenshot Output Soal 2](/praktikum/04-runtunan-sekuensi/TP%20Modul%204/Nomor%202/output_nomor2.png)

#### Deskripsi
Nilai awal: `x = 10`, `y = 5`, `z = 15`, `result = 0`. Keempat blok `if` dijalankan berurutan dan nilai `result` dibawa ke blok berikutnya.

**Alur eksekusi**
- **Kondisi 1:** `x > 5` → `10 > 5` benar, program masuk ke `if` di dalamnya. `y < 10` → `5 < 10` benar, sehingga `result = x + y = 10 + 5 = 15`.
- **Kondisi 2:** `z > 10 && x == 10` → `true && true` = benar, sehingga `result += z` menjadi `15 + 15 = 30`. Blok `else` dilewati.
- **Kondisi 3:** `x == 10 || y > 10` → `true || false` = benar (cukup salah satu sisi), sehingga `result += 5` menjadi `30 + 5 = 35`. Cabang `else if` dan `else` tidak dicek.
- **Kondisi 4:** `x < 15 && y < 10` → `true && true` = true, lalu dibalik oleh `!` menjadi **false**. Program masuk ke `else`, sehingga `result -= 10` menjadi `35 - 10 = 25`.

| Setelah | Kondisi | Aksi | result |
|---|---|---|---|
| Awal | - | - | 0 |
| Kondisi 1 | true | `result = x + y` | 15 |
| Kondisi 2 | true | `result += z` | 30 |
| Kondisi 3 | true | `result += 5` | 35 |
| Kondisi 4 | false | `result -= 10` (else) | 25 |

**Jawaban pertanyaan**
1. Nilai akhir `result` adalah **25**. Nilai variabel lain tidak berubah: `x = 10`, `y = 5`, `z = 15`.
2. Output program:
   ```
   Nilai akhir result: 25
   ```

### 3. Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan Bulan

```go
package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun, &bulan)

	kabisat := (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)
	case "Feb":
		if kabisat {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		fmt.Println("Input bulan tidak valid")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder soal3/output1.png dan output2.png -->
![Screenshot Output Soal 3](/praktikum/04-runtunan-sekuensi/TP%20Modul%204/Nomor%203/output_nomor3.png)

#### Deskripsi
Program membaca dua input, yaitu `tahun` (int) dan `bulan` (string) dengan `fmt.Scan`. Tahun kabisat ditentukan dengan rumus `(tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0`, yaitu habis dibagi 4 dan tidak habis dibagi 100, atau habis dibagi 400.

Jumlah hari ditentukan dengan `switch` pada nama bulan:
- 31 hari: Jan, Mar, Mei, Jul, Agu, Okt, Des.
- 30 hari: Apr, Jun, Sep, Nov.
- Feb: 29 hari jika tahun kabisat, selain itu 28 hari.
- Selain nama di atas, program menampilkan pesan kesalahan lewat `default`.

Pencocokan `switch` pada string bersifat *case-sensitive*, sehingga `jan` tidak cocok dengan `Jan` dan masuk ke `default`. Hasil pengujian:

| Masukan | Keluaran |
|---|---|
| `2001` `Jan` | `31` |
| `2016` `jan` | `Input bulan tidak valid` |

### 4. Switch Case

```go
package main

import "fmt"

func main() {
	var hari int

	fmt.Print("Masukkan angka hari (1-7): ")
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Angka tidak valid")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder soal4/output.png -->
![Screenshot Output Soal 4](/praktikum/04-runtunan-sekuensi/TP%20Modul%204/Nomor%204/output_nomor4.png)
#### Deskripsi
Program mengubah angka 1 sampai 7 menjadi nama hari menggunakan `switch case`. Nilai `hari` dibandingkan dengan tiap `case` dari atas ke bawah. Begitu ada yang cocok, kodenya dijalankan lalu `switch` selesai tanpa perlu `break`, karena Go tidak melanjutkan ke `case` berikutnya secara otomatis. Jika tidak ada yang cocok, `default` dijalankan dan program menampilkan `Angka tidak valid`.

Contoh hasil: input `3` menghasilkan `Rabu`, input `9` menghasilkan `Angka tidak valid`.

## Kesimpulan
Pada modul ini dipelajari cara kerja pernyataan kondisi dalam Go. Operator relasional (`>`, `<=`, `==`, `!=`) dan operator logika (`&&`, `||`, `!`) dipakai untuk membentuk ekspresi yang menghasilkan `true` atau `false`, dengan memperhatikan urutan prioritas operator dan sifat *short-circuit* pada `&&` dan `||`. Tracing pada Soal 2 menunjukkan bahwa nilai variabel dapat berubah di setiap blok `if`, dan pada rantai `if / else if / else` hanya satu cabang yang dijalankan. Soal 3 dan 4 menerapkan `switch case` untuk memilih aksi berdasarkan nilai tertentu, ditambah logika tahun kabisat pada Soal 3, sehingga program menjadi lebih ringkas dan mudah dibaca dibanding rangkaian `if` yang panjang.

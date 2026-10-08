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

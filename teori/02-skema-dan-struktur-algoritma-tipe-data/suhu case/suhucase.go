package main

import (
	"fmt"
)

func main() {
	var umur int8
	var suhu float32

	suhu = 36.3
	umur = 10

	fmt.Println("Umur: ", umur)
	fmt.Println("Suhu: ", suhu)
	fmt.Println("Alamat memori dari variabel umur: ", &umur)
	fmt.Println("Alamat memori dari variabel suhu: ", &suhu)
}

package main

import "fmt"

func main() {
	var name string
	name = "Deta Arief Syahputra"
	fmt.Println("Hello, " + name + "!")

	var firstName = "Deta"
	fmt.Println("Nama Depan : ", firstName)

	var middleName = "Arief"
	fmt.Println("Nama Tengah : ", middleName)

	var lastName = "Syahputra"
	fmt.Println("Nama Belakang : ", lastName)

	var (
		fullName = "Deta Arief Syahputra"
		nickName = "Deta"
	)

	fmt.Println("Nama Lengkap : ", fullName)
	fmt.Println("Nama Panggilan : ", nickName)
}

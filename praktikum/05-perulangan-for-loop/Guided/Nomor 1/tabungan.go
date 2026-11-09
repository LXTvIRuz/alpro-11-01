package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	total := 0
	tabunganhariini := a

	for i := 1; i <= n; i++ {
		total = total + tabunganhariini
		tabunganhariini = tabunganhariini + b
	}
	fmt.Println(total)
}

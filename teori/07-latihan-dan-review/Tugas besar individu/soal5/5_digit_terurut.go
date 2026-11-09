package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	asc, desc := true, true
	for i := 1; i < len(n); i++ {
		if n[i] < n[i-1] {
			asc = false
		}
		if n[i] > n[i-1] {
			desc = false
		}
	}

	if asc {
		fmt.Println("ascending")
	} else if desc {
		fmt.Println("descending")
	} else {
		fmt.Println("tidak terurut")
	}
}

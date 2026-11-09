package main

import "fmt"

func main() {
	var x, prev int
	hari, total, naik := 0, 0, 0

	for {
		fmt.Scan(&x)
		if x < 0 || x > 200 {
			break
		}
		hari++
		total += x
		if hari > 1 && x > prev {
			naik++
		}
		prev = x
	}

	if hari == 0 {
		fmt.Println(0, 0)
	} else {
		fmt.Printf("%d %.2f\n", naik, float64(total)/float64(hari))
	}
}

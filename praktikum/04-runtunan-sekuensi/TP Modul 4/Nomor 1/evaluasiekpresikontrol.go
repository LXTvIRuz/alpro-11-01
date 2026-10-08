package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	if intOther+2*intNum != 30 || !(sngNum > 0) {
		fmt.Println("Ding! Ding! Ding!")
	}
}

package main

import "fmt"

func main() {
	var a, b, hasil float64
	fmt.Scan(&a, &b)
	if b != 0 {
		hasil = a / b
		fmt.Print("hasilnya adalah ", hasil)

	}
}

package main

import (
	"fmt"
)

func main() {
	var x, n int

	fmt.Print("Silahkan masukan nilai x: ")
	fmt.Scanln(&x)

	fmt.Print("Lalu masukan nilai n: ")
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println("Hasilnya adalah", 1)
	}
	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}
	fmt.Println("Hasilnya adalah", result)
}

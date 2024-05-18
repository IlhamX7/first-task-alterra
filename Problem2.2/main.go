package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Silahkan masukan bilangan: ")
	fmt.Scanln(&bilangan)

	for i := bilangan; i >= 1; i-- {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}

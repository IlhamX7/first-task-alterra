package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Silahkan masukan bilangan: ")
	fmt.Scanln(&bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}

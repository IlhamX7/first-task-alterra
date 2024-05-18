package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan int

	fmt.Print("Silahkan masukan bilangan: ")
	fmt.Scanln(&bilangan)

	if bilangan <= 1 {
		fmt.Println("false")
	} else if bilangan == 2 {
		fmt.Println("true")
	} else if bilangan%2 == 0 {
		fmt.Println("false")
	} else {
		sqrt := int(math.Sqrt(float64(bilangan)))
		for i := 3; i <= sqrt; i += 2 {
			if bilangan%i == 0 {
				fmt.Println("false")
				break
			}
		}
		fmt.Println("true")
	}
}

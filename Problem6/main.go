package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var bilangan int

	fmt.Print("Silahkan masukan bilangan: ")
	fmt.Scanln(&bilangan)

	if bilangan <= 1 {
		fmt.Println("false")
	} else if bilangan == 2 && areAllDigitsPrime(bilangan) {
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
		if areAllDigitsPrime(bilangan) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}

func areAllDigitsPrime(n int) bool {
	primeDigits := map[rune]bool{'2': true, '3': true, '5': true, '7': true}
	for _, digit := range strconv.Itoa(n) {
		if !primeDigits[digit] {
			return false
		}
	}
	return true
}

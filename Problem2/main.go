package main

import (
	"fmt"
	"math"
)

func getMinMax(numbers ...*int) (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for _, num := range numbers {
		if *num > max {
			max = *num
		}
		if *num < min {
			min = *num
		}
	}
	return max, min
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Masukkan 6 angka:")
	fmt.Scanln(&a1)
	fmt.Scanln(&a2)
	fmt.Scanln(&a3)
	fmt.Scanln(&a4)
	fmt.Scanln(&a5)
	fmt.Scanln(&a6)
	max, min = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}

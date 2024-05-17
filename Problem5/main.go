package main

import (
	"fmt"
	"math"
)

func main() {
	var r, T float64

	r = 4
	T = 20

	luasPermukaan := 2 * math.Pi * r * (r + T)
	fmt.Printf("Luas permukaan tabung adalah: %.2f\n", luasPermukaan)
}

package main

import (
	"fmt"
)

func main() {
	var hargaBaju, diskon float64

	hargaBaju = 370000
	diskon = 15

	hargaDiskon := (diskon / 100) * hargaBaju
	totalHarga := hargaBaju - hargaDiskon

	fmt.Printf("Total harga setelah diskon adalah: %.2f\n", totalHarga)
}

package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	found := false

	for x := -100; x <= 100; x++ {
		for y := -100; y <= 100; y++ {
			if x != y {
				for z := -100; z <= 100; z++ {
					if x != z && y != z {
						if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
							fmt.Println(x, y, z)
							found = true
							break
						}
					}
				}
				if found {
					break
				}
			}
		}
		if found {
			break
		}
	}

	if !found {
		fmt.Println("No solution.")
	}
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}

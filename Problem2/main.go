package main

import (
	"fmt"
)

func moneyChange(money int) []int {
	denominations := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	result := []int{}

	for _, denomination := range denominations {
		for money >= denomination {
			money -= denomination
			result = append(result, denomination)
		}
	}

	return result
}

func main() {
	fmt.Println(moneyChange(123))
	fmt.Println(moneyChange(432))
}

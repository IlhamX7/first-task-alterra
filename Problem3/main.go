package main

import "fmt"

func ArrayUnique(arrayA, arrayB []int) []int {
	elementMap := make(map[int]bool)
	for _, num := range arrayB {
		elementMap[num] = true
	}

	var result []int
	for _, num := range arrayA {
		if !elementMap[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	arrayA := []int{3, 8}
	arrayB := []int{2, 8}
	fmt.Println(ArrayUnique(arrayA, arrayB))
}

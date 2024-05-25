package main

import "fmt"

func RemoveDuplicates(array []int) int {
	if len(array) == 0 {
		return 0
	}
	uniqueIndex := 1

	for i := 1; i < len(array); i++ {
		if array[i] != array[i-1] {
			array[uniqueIndex] = array[i]
			uniqueIndex++
		}
	}
	return uniqueIndex
}

func main() {
	array := []int{2, 2, 2, 11}
	fmt.Println(RemoveDuplicates(array))
}

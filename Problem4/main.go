package main

import "fmt"

func BinarySearch(array []int, x int) {
	beg := 0
	end := len(array) - 1

	for beg <= end {
		mid := (beg + end) / 2

		if array[mid] == x {
			fmt.Println(mid)
			return
		} else if array[mid] < x {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Println(-1)
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)
}

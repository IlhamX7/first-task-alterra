package main

import "fmt"

func findMaxSumSubArray(k int, arr []int) int {
	n := len(arr)
	if n < k {
		return 0
	}

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum

	for i := k; i < n; i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	k := 2
	arr := []int{2, 3, 4, 1, 5}
	fmt.Println(findMaxSumSubArray(k, arr))
}

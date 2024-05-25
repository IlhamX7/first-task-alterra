package main

import (
	"fmt"
)

func Compare(a, b string) string {
	m, n := len(a), len(b)
	longest, endIndex := 0, 0

	lcsuff := make([][]int, m+1)
	for i := range lcsuff {
		lcsuff[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				lcsuff[i][j] = lcsuff[i-1][j-1] + 1
				if lcsuff[i][j] > longest {
					longest = lcsuff[i][j]
					endIndex = i
				}
			}
		}
	}

	return a[endIndex-longest : endIndex]
}

func main() {
	a := "KUPU-KUPU"
	b := "KUPU"
	fmt.Println(Compare(a, b))
}

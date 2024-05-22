package main

import (
	"fmt"
	"strings"
)

func DrawXYZ(N int) string {
	var result strings.Builder
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			position := (i-1)*N + j
			if position%3 == 0 {
				result.WriteString("X ")
			} else if position%2 == 0 {
				result.WriteString("Z ")
			} else {
				result.WriteString("Y ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func main() {
	input := 5
	fmt.Println(DrawXYZ(input))
}

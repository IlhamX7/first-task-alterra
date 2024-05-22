package main

import (
	"fmt"
	"strings"
)

func playWithAsterik(bintang int) string {
	var result strings.Builder
	for i := 1; i <= bintang; i++ {
		spaces := strings.Repeat(" ", bintang-i)
		stars := strings.Repeat("*", i)
		result.WriteString(spaces + stars + "\n")
	}
	return result.String()
}

func main() {
	bintang := 5
	fmt.Println(playWithAsterik(bintang))
}

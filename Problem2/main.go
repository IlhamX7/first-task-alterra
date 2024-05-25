package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	shifted := []rune(input)

	for i, char := range input {
		if char >= 'a' && char <= 'z' {
			shifted[i] = 'a' + (char-'a'+rune(offset))%26
		} else {
			shifted[i] = char
		}
	}

	return string(shifted)
}

func main() {
	offset := 2
	input := "alta"
	fmt.Println(caesar(offset, input))
}

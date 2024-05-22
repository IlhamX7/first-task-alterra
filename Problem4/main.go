package main

import (
	"fmt"
	"strings"
)

func ubahHuruf(setence string) string {
	originalAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shiftedAlphabet := "KLMNOPQRSTUVWXYZABCDEFGHIJ"

	shiftMap := make(map[rune]rune)
	for i, char := range originalAlphabet {
		shiftMap[char] = rune(shiftedAlphabet[i])
	}

	upperInput := strings.ToUpper(setence)
	var result strings.Builder

	for _, char := range upperInput {
		if char == ' ' {
			result.WriteRune(' ')
		} else {
			shiftedChar, exists := shiftMap[char]
			if exists {
				result.WriteRune(shiftedChar)
			} else {
				result.WriteRune(char)
			}
		}
	}

	return result.String()
}

func main() {
	input := "SEPULSA OKE"
	fmt.Println(ubahHuruf(input))
}

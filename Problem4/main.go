package main

import "fmt"

type student struct {
	name     string
	nameCode string
	score    int
}

type Cipher interface {
	encode(text string) string
	decode(text string) string
}

type SubstitutionCipher struct {
	shift int
}

func NewSubstitutionCipher(shift int) *SubstitutionCipher {
	return &SubstitutionCipher{shift: shift}
}

func (s *SubstitutionCipher) Encode(text string) string {
	encoded := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			encoded += string((char-'a'+rune(s.shift))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			encoded += string((char-'A'+rune(s.shift))%26 + 'A')
		} else {
			encoded += string(char)
		}
	}
	return encoded
}

func (s *SubstitutionCipher) Decode(text string) string {
	decoded := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			decoded += string((char-'a'-rune(s.shift)+26)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			decoded += string((char-'A'-rune(s.shift)+26)%26 + 'A')
		} else {
			decoded += string(char)
		}
	}
	return decoded
}

func main() {
	var choice int
	var name string

	cipher := NewSubstitutionCipher(3)

	fmt.Println("[1] Encrypt")
	fmt.Println("[2] Decrypt")
	fmt.Print("Choose your menu? ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Input Student's Name: ")
		fmt.Scan(&name)
		encodedName := cipher.Encode(name)
		fmt.Printf("Encode of Student's Name %s is %s\n", name, encodedName)
	case 2:
		fmt.Print("Input Encoded Name: ")
		fmt.Scan(&name)
		decodedName := cipher.Decode(name)
		fmt.Printf("Decode of Encoded Name %s is %s\n", name, decodedName)
	default:
		fmt.Println("Invalid choice")
	}
}

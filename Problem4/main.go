package main

import (
	"fmt"
	"strings"
)

func main() {
	var kata string
	var kataDibalik string = ""

	fmt.Print("Silahkan masukan kata: ")
	fmt.Scanln(&kata)
	length := len(kata)

	for i := length - 1; i >= 0; i-- {
		kataDibalik = kataDibalik + string(kata[i])
	}

	if strings.ToLower(kata) == strings.ToLower(kataDibalik) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

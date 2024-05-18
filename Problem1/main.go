package main

import "fmt"

func main() {
	var studentScore int

	fmt.Print("Silahkan masukan nilai: ")
	fmt.Scanln(&studentScore)

	if studentScore <= 34 {
		fmt.Println("Nilai D")
	} else if studentScore <= 49 && studentScore >= 35 {
		fmt.Println("Nilai C")
	} else if studentScore <= 64 && studentScore >= 50 {
		fmt.Println("Nilai B")
	} else if studentScore <= 79 && studentScore >= 65 {
		fmt.Println("Nilai B+")
	} else if studentScore <= 100 && studentScore >= 80 {
		fmt.Println("Nilai A")
	} else {
		fmt.Println("Maaf score yang anda masukan tidak ada")
	}
}

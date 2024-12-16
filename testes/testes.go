package main

import (
	"fmt"
	"time"
)

// func main() {
// 	tm := time.Now()

// 	fmt.Printf("%s", tm)
// }

func main() {
	currentYear := time.Now().Year()

	var age int
	fmt.Println("Digite sua idade: ")
	fmt.Scanln(&age)

	yearOfBirth := currentYear - age

	fmt.Printf("Seu ano de nascimento é: %d\n", yearOfBirth)
}

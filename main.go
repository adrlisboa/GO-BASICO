package main

import "fmt"

// func main() {
// 	var name string
// 	fmt.Println("Digite seu nome")
// 	fmt.Scanln(&name)

// 	fmt.Printf("Bem-vindo ao mundo Golang, %s!\n", name)
// }

func main() {
	var num1, num2 int
	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)

	sum := num1 + num2
	fmt.Printf("A soma de %d e %d é igual a %d \n", num1, num2, sum)
}

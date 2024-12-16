package main

import "fmt"

func main() {
	var age int
	for {
		fmt.Println("Por favor digite sua idade: (+18)")
		_, err := fmt.Scanln(&age)
		if err != nil {
			fmt.Println("Entrada inválida, por favor, repita")
			continue
		}

		if age < 18 {
			fmt.Println("Entrada inválida, por favor, repita")
			continue
		}
		break
	}

	fmt.Println("Muito bem! Seja Bem-Vindo!")
}

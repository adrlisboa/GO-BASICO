package main

import "fmt"

func main() {
	numero := 42
	ponteiro := &numero
	fmt.Println("Antes da alteração:", numero)
	// Mostra 42
	*ponteiro = 99
	// Altera o valor armazenado no endereço
	fmt.Println("Depois da alteração:", numero)
	// Mostra 99
}

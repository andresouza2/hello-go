package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle em Golang")

	numero := 10

	if numero > 15 { 
		fmt.Println("maior que 15") 
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if outroNumero < -10 {
		fmt.Println("Número é menor que menos dez")
	} else {
		fmt.Println("Número está entre -10 e 0")
	}
}

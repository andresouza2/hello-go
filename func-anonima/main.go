package main

import "fmt"

func main() {
	fmt.Println("Funções Anônimas em Golang")

	resultado := func (texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá mundo!")

	fmt.Println(resultado)
}
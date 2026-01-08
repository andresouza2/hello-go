package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	rua string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var user usuario
	user.nome = "John Doe"
	user.idade = 36
	fmt.Println(user)

	enderecoExemplo := endereco{"Rua sem saida", 23}

	user2 := usuario{"Jane Doe", 28, enderecoExemplo}
	fmt.Println(user2)

	user3 := usuario{idade: 23}
	fmt.Println(user3)
}

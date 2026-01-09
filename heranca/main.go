package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade uint8
	altura float32
}

type estudante struct{
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Hello!")
	p1 := pessoa{"John Doe", "Silva", 30, 1.65}

	e1 := estudante{p1, "Sistemas de Informação", "UniFacul"}
	fmt.Println("nome:",e1.nome)
	fmt.Println("sobrenome:",e1.sobrenome)
	fmt.Println("idade:",e1.idade)
	fmt.Println("altura:",e1.altura)
	fmt.Println("curso:",e1.curso)
	fmt.Println("faculdade:",e1.faculdade)

}
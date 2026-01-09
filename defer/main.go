package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funçao 1")
}
func funcao2() {
	fmt.Println("Executando a funçao 2")
}
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2
	if media >= 6 {		
		return true
	}

	return false
}

func main() {
	fmt.Println("Defer Example in Go")

	// DEFER = ADIAR
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoEstaAprovado(2, 4))
}
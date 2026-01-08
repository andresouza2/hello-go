package main

import "fmt"

func main() {
	var nome string = "golang"
	nomeDois := "go dois"

	fmt.Println(nome)
	fmt.Println(nomeDois)

	var (
		nomeTres string = "go três"
		nomeQuatro string = "go quatro"
	)

	fmt.Println(nomeTres, nomeQuatro)

	idade1, idade2 := 36, 50

	fmt.Println(idade1, idade2)

	// constantes
	const contanteUm string = "constante um"
	fmt.Println(contanteUm)

	nomeTres, nomeQuatro = nomeQuatro, nomeTres
	fmt.Println(nomeTres, nomeQuatro)
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Looping in Go! Go! Go!...")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println("Finalizado!", i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"Ana", "Bia", "Carlos"}

	// for _, nome := range nomes {
	// 	fmt.Println("Nome:", nome)
	// 	time.Sleep(time.Second)
	// }

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// }

	usuario := map[string]string {
		"nome": "John",
		"sobrenome": "Doe",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, ":", valor)
	}

	// obs: não se pode usar iteração com structs
}
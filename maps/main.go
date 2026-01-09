package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	usuario := map[string]string {
		"nome": "John",
		"sobrenome": "Doe",
		//"idade": 30, - não pode misturar tipos
	}

	fmt.Println(usuario)
	// fmt.Println(usuario["nome"])
	// fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "Jane",
			"ultimo": "Doe",
		},
		"cursos": {
			"c2": "React",
			"c1": "Golang",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string {
		"nome": "Leão",
	}
	fmt.Println(usuario2)
}
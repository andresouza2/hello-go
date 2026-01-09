package main

import "fmt"

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func diaDaSemana2(dia int) string {
	var diaDaSemana string
	switch {
		case dia == 1: 
			diaDaSemana = "Domingo"
		case dia == 2:
			diaDaSemana = "Segunda-feira"
		case dia == 3:
			diaDaSemana = "Terça-feira"
		case dia == 4:
			diaDaSemana = "Quarta-feira"
		case dia ==5:
			diaDaSemana = "Quinta-feira"
		case dia == 6:
			diaDaSemana = "Sexta-feira"
		case dia == 7:
			diaDaSemana = "Sábado"
		default:
			diaDaSemana = "Dia inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch em Golang")

	fmt.Println(diaDaSemana(10))
	fmt.Println("--------------------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
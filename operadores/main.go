package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 3 - 2
	multiplicacao := 2 * 2
	divisao := 10 / 2
	resto := 10 % 3

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	var num1 int16 = 10
	var num2 int16 = 20
	somaDois := num1 + num2
	fmt.Println(somaDois)
	// Fim dos aritméticos

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// Fim das atribuição

	// Operadores relacionais
	fmt.Println("===== Operadores relacionais =====")
	fmt.Println("1 > 2:", 1 > 2)
	fmt.Println("1 >= 2:", 1 >= 2)
	fmt.Println("1 == 2:", 1 == 2)
	fmt.Println("1 < 2:", 1 < 2)
	fmt.Println("1 <= 2:", 1 <= 2)
	fmt.Println("1 != 2:", 1 != 2)
	// Fim dos relacionais

	// Operadores lógicos
	fmt.Println("===== Operadores lógicos =====")
	verdadeiro, falso := true, false
	fmt.Println("verdadeiro && falso:", verdadeiro && falso)
	fmt.Println("verdadeiro || falso:", verdadeiro || falso)
	fmt.Println("!verdadeiro:", !verdadeiro)
	fmt.Println("!falso:", !falso)
	// Fim dos lógicos

	// Operadores unários
	numero := 10
	numero++
	numero += 5
	fmt.Println(numero)

	numero-- // numero = numero - 1
	numero -= 3 // numero = numero - 3
	numero *= 2 // numero = numero * 2
	numero /= 4 // numero = numero / 4
	numero %= 3 // numero = numero % 3
	fmt.Println(numero)
	// Fim dos unários

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
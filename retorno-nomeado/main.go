package main

import "fmt"

func calcMath(a int, b int) (soma int, subtracao int) {
	soma = a + b
	subtracao = a - b
	return
}

func main() {
	soma, subtracao := calcMath(10, 5)
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
}
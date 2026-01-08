package main

import "fmt"

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := soma(32, 32)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("print", txt)
		return txt
	}

	fmt.Println(f("Olá, mundo!"))

	resultSoma, _ := calculosMatematicos(50, 25)
	fmt.Println(resultSoma)
}
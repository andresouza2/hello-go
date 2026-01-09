package main

import "fmt"

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := Sum(1, 2, 3, 4, 5)
	fmt.Println(totalSoma)

	escrever("Olá mundo", 23, 54, 20, 652)
}
package main

import (
	"errors"
	"fmt"
)

func main() {
	//aula daqui
	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	var numero2 uint64 = 10000000000000000000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// INT8 = BYTE
	var numero4 byte = 255
	fmt.Println(numero4)

	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 999999999999999.45
	fmt.Println(numeroReal2)

	// inferência
	nomeroReal3 := 123456789.45
	fmt.Println(nomeroReal3)

	// FIM DOS NÚMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// FIM STRINGS E CHARS

	// boolean
	var bool1 bool = true
	fmt.Println(bool1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}

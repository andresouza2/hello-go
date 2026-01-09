package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices in Golang")
	const lines = "-------------------------"

	// Arrays
	fmt.Println("Arrays:")
	var array1 [3]string
	array1[0] = "Apple"
	fmt.Println(array1)

	array2 := [3]string{"Banana", "Grapes", "Orange"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	
	// Slices
	fmt.Println("Slices:")
	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println(slice1)
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	fmt.Println(lines)
	slice1 = append(slice1, 18)
	fmt.Println(slice1)

	fmt.Println(lines)
	slice2 := array2[1:3]
	fmt.Println(slice2)

	fmt.Println(lines)
	array2[1] = "Mango"
	fmt.Println(slice2)

}
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("arrays e slices")

	var array1 [5]string // todos os dados precisam ser do mesmo tipo, no array
	array1[0] = "posição 1"
	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)
	// array2[5] = "Posição 6"
	// fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} //mesma coisa que [5]int{...}
	fmt.Println(array3)

	//SLICE

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	//slice não é um array

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
	// tipos diferentes, não da para fazer operações com eles

	var slice1 []int //posso usar a mesma variaver se eu quiser
	slice1 = append(slice, 18)
	fmt.Println(slice1) // [10 11 12 13 14 15 16 17 18] acrescenta e retorna um novo

	slice2 := array2[1:3]
	fmt.Println(slice2) //[posição 2 posição 3]

	array2[1] = "Posição Alterada"
	fmt.Println(slice2) //[Posição Alterada posição 3]
}

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

	//ARRAYS INTERNOS
	fmt.Println("-------------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)      // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(slice3)) //lenght  //10
	fmt.Println(cap(slice3)) //capacidadee //11

	slice3 = append(slice3, 11) // aqui atingiu o limite to slice que aponta p array
	slice3 = append(slice3, 12) // aqui eu add mais um, estourando !!!
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // 12 *****-
	fmt.Println(cap(slice3)) // 24 *****-  go além de deixar incluir, ainda dobrou a capacidade de alocação, por conta disso que não declaro o tamanho no slice, pois é dinamico sacou, só nos arraysss

	//O TAMANHO É OPCIONALIS
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 6) // adiconando é aquilo , tam 6 e capacidade de 12
	fmt.Println(len(slice4))   //5
	fmt.Println(cap(slice4))   //5
	//LOGO, A CAPACIDADE VAI SER O PRÓPRIO TAMANHO DELE

	//SLICE É UMA LISTA SEM TAM FIXO
	//ARRAY É UMA LISTA DE TAM FIXO













	
	slice7 := make([]int64, 10, 11)
	slice7 = append(slice7, 11)
	slice7 = append(slice7, 12)
	fmt.Println("slice7", slice7)
	fmt.Println(len(slice7))
	fmt.Println(cap(slice7))

}

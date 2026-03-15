package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j <= 10; j += 2 { // posso controlar com j += 2 quantos que vai acrescentarr
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"joão", "Davi", "Lucas"}

	// for indice, nome := range nomes {
	// 	fmt.Println(indice, nome)
	// 	//SAÍDA
	// 	// 0 joão
	// 	// 1 Davi
	// 	// 2 Lucas
	// }

	// for _, nome := range nomes {  // underline para ignorar o indice, retornando apenas os nomes
	// 	fmt.Println(nome)
	// joão Davi Lucas
	// }

	// for indice, letra := range "PALAVRA" { // CPF
	// 	fmt.Println(indice, string(letra))
	// 	0 P
	//  1 A
	//  2 L
	// 	3 A
	// 	4 V
	// 	5 R
	// 	6 A
	// }

	// usuario := map[string]string{
	// 	"nome":      "Leonardo",
	// 	"sobrenome": "Silva",
	// }

	// fmt.Println(usuario)

	// for chave, valor := range usuario {
	// 	fmt.Println(chave, valor)
	// }

	//EM STRUCTS NÃO DA PRA FAZER :/

	// type usuarioStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }

	// usuario2 := usuarioStruct{"zé","robson"}

	// for chave, valor := range usuario2 { //cannot range over usuario2 (variable of struct type usuarioStruct)compilerInvalidRangeExpr
	// 	fmt.Println(chave,valor)
	// }

	// curiosidade q eu não deveria saber ...

	for {
		fmt.Println("executando infinitamente")
		time.Sleep(time.Second)
	}

}

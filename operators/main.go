package main

import "fmt"

func main() {
	//ARITIMÉTICOS, nada de novo
	// +
	// -
	// /
	// *
	// %

	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10.0 / 4.0
	multiplicacao := 4 * 10
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// não posso fazer NADA!!!!!! se as variaveis são de tipos diferentes, exemplo abaixo
	// var n1 int16 = 10
	// var n2 int32 = 20
	// somar := n1 + n2
	// fmt.Println(somar)

	//agora sim, converti n2 para ser do mesmo tipo do n1
	var n1 int16 = 10
	var n2 int32 = 20
	somar := n1 + int16(n2)
	fmt.Println(somar)

	//FIM DOS ARITIMÉTICOS

	// *-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*

	//ATRIBUIÇÃO

	var variavel string = "String"
	variavel2 := "\nString2"
	fmt.Println(variavel, variavel2)

	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	// *-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*

	//OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//FIM DOS OPERADORES RELACIONAIS

	// *-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*

	// OPERADORES LÓGICOS
	fmt.Println("-------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// FIM OPERADORES LÓGICOS

	// *-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*

	//OPERADORES UNÁRIOS

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 20
	numero *= 5
	numero++
	numero /= 6
	numero %= 6
	fmt.Println(numero)

	// FIM OPERADORES UNÁRIOS

	//ternários não existem em go ! :) bora para o bom e velho if

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)
}

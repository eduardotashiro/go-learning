package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estrutura de controle")

	numero := -5

	if numero >= 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que zero")
	} else if numero < -3 {
		fmt.Println("numero é menor que zero")
	} else {
		fmt.Println("numero é menor que zero")
	}

	//fmt.Println(outroNumero) //undefined: outroNumerocompiler não podemos acessar fora do escopo do if

}

package main

import (
	"fmt"
)

func inverterSinal(numero int) int { // passando um parametro por copia
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) { // passando uma referencia para a func
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido) // -20
	fmt.Println(numero)          // 20
	fmt.Println(numero)          // 20

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero) // -40
	fmt.Println(novoNumero) // -40
}

package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2) //10 10
	variavel1++
	fmt.Println(variavel1, variavel2) //11 10

	//PONTEIRO É UM REFERÊNCIA DE MEMÓRIA

	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro) // 0 <nil>

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)  // 100   0x10cf7bbf80d8
	fmt.Println(variavel3, *ponteiro) // desreferenciação   //100 100

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)  // 150 0x2c00d89780d8
	fmt.Println(variavel3, *ponteiro) // 150 150

	var (
		numero    int
		texto     string
		valores   [5]int
		auxiliarN *int
		auxiliarS *string
		auxiliarV *[5]int
	)

	numero = 1
	texto = "texto"
	valores = [5]int{1, 2, 3, 4, 5}

	auxiliarN = &numero
	auxiliarS = &texto
	auxiliarV = &valores

	fmt.Println(numero, texto, valores, *auxiliarN, *auxiliarS, auxiliarV) // 15 min para entender que a desreferenciação é automática para array

}

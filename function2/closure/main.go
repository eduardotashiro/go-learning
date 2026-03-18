package main

import "fmt"

func closure() func() {
	txt := "dentro da funcao closure"
	funcao := func() {
		fmt.Println(txt)
	}
	return funcao
}

func main() {
	texto := "dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}

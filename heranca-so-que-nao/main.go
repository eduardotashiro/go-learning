package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("herança")

	p1 := pessoa{"eduardo", "odraude", 25, 190}
	fmt.Println(p1)

	e1 := estudante{ p1, "psicologia", "UFSC"}
	fmt.Println(e1.nome)

}

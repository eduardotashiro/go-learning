package main

import "fmt"

type usuario struct { // lembra a interface do ts
	nome     string
	idade    uint8
	endereco endereco // struct dentro deoutra struct
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")
	// """"""""instanciando'"""""""""
	var u usuario
	u.nome = "tashiro"
	u.idade = 25
	fmt.Println(u) //{tashiro 25 { 0}} //aninhamento ?

	endereco := endereco{"rua jau", 8}

	u2 := usuario{"eduardo", 25, endereco}
	fmt.Println(u2) // {eduardo 25 {rua jau 8}}

	u3 := usuario{idade: 25}
	fmt.Println(u3) // { 25 { 0}}

}

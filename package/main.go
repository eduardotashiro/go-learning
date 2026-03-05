package main

import (
	"modulo/utils"
	"fmt"
	"github.com/badoux/checkmail" // CMD  go get github.com/badoux/checkmail
)

func main() {
	fmt.Println("escrevendo do arquivo main")
	utils.Escrever()

	//usando pacotes externos
	erro := checkmail.ValidateFormat("bugzilla@gmail.com")
	fmt.Println(erro)
}
